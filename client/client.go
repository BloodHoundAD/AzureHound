// Copyright (C) 2022 Specter Ops, Inc.
//
// This file is part of AzureHound.
//
// AzureHound is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// AzureHound is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package client

//go:generate go run go.uber.org/mock/mockgen -destination=./mocks/client.go -package=mocks . AzureClient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bloodhoundad/azurehound/v2/client/config"
	"github.com/bloodhoundad/azurehound/v2/client/query"
	"github.com/bloodhoundad/azurehound/v2/client/rest"
	"github.com/bloodhoundad/azurehound/v2/models/azure"
	"github.com/bloodhoundad/azurehound/v2/panicrecovery"
	"github.com/bloodhoundad/azurehound/v2/pipeline"
)

func NewClient(config config.Config) (AzureClient, error) {
	if msgraph, err := rest.NewRestClient(config.GraphUrl(), config); err != nil {
		return nil, err
	} else if resourceManager, err := rest.NewRestClient(config.ResourceManagerUrl(), config); err != nil {
		return nil, err
	} else {
		if config.JWT != "" {
			if aud, err := rest.ParseAud(config.JWT); err != nil {
				return nil, err
			} else if aud == config.GraphUrl() {
				return initClientViaGraph(msgraph, resourceManager)
			} else if aud == config.ResourceManagerUrl() {
				if body, err := rest.ParseBody(config.JWT); err != nil {
					return nil, err
				} else {
					return initClientViaRM(msgraph, resourceManager, body["tid"])
				}
			} else {
				return nil, fmt.Errorf("error: invalid token audience")
			}
		} else {
			return initClientViaGraph(msgraph, resourceManager)
		}
	}
}

func initClientViaRM(msgraph, resourceManager rest.RestClient, tid interface{}) (AzureClient, error) {
	client := &azureClient{
		msgraph:         msgraph,
		resourceManager: resourceManager,
	}
	if result, err := client.GetAzureADTenants(context.Background(), true); err != nil {
		return nil, err
	} else {
		for _, tenant := range result.Value {
			if tenant.TenantId == tid.(string) {
				client.tenant = tenant
				break
			}
		}
		return client, nil
	}
}

func initClientViaGraph(msgraph, resourceManager rest.RestClient) (AzureClient, error) {
	client := &azureClient{
		msgraph:         msgraph,
		resourceManager: resourceManager,
	}
	if org, err := client.GetAzureADOrganization(context.Background(), nil); err != nil {
		return nil, err
	} else {
		client.tenant = org.ToTenant()
		return client, nil
	}
}

type azureResult[T any] struct {
	Error error
	Ok    T
}

func getAzureObjectList[T any](client rest.RestClient, ctx context.Context, path string, params query.Params, out chan azureResult[T]) {
	defer panicrecovery.PanicRecovery()
	defer close(out)

	var (
		errResult azureResult[T]
		nextLink  string
	)

	for {
		var (
			list struct {
				CountGraph    int    `json:"@odata.count,omitempty"`    // The total count of all graph results
				NextLinkGraph string `json:"@odata.nextLink,omitempty"` // The URL to use for getting the next set of graph values.
				ContextGraph  string `json:"@odata.context,omitempty"`
				NextLinkRM    string `json:"nextLink,omitempty"` // The URL to use for getting the next set of rm values.
				Value         []T    `json:"value"`              // A list of azure values
			}
			res *http.Response
			err error
		)

		if nextLink != "" {
			if nextUrl, err := url.Parse(nextLink); err != nil {
				errResult.Error = err
				_ = pipeline.Send(ctx.Done(), out, errResult)
				return
			} else {
				if req, err := rest.NewRequest(ctx, "GET", nextUrl, nil, params.AsMap(), nil); err != nil {
					errResult.Error = err
					_ = pipeline.Send(ctx.Done(), out, errResult)
					return
				} else if res, err = client.Send(req); err != nil {
					errResult.Error = err
					_ = pipeline.Send(ctx.Done(), out, errResult)
					return
				}
			}
		} else {
			if res, err = client.Get(ctx, path, params, nil); err != nil {
				errResult.Error = err
				_ = pipeline.Send(ctx.Done(), out, errResult)
				return
			}
		}

		if err := rest.Decode(res.Body, &list); err != nil {
			errResult.Error = err
			_ = pipeline.Send(ctx.Done(), out, errResult)
			return
		} else {
			for _, u := range list.Value {
				if ok := pipeline.Send(ctx.Done(), out, azureResult[T]{Ok: u}); !ok {
					return
				}
			}
		}

		if list.NextLinkRM == "" && list.NextLinkGraph == "" {
			break
		} else if list.NextLinkGraph != "" {
			nextLink = list.NextLinkGraph
		} else if list.NextLinkRM != "" {
			nextLink = list.NextLinkRM
		}
	}
}

func getAzureObject[T any](client rest.RestClient, ctx context.Context, path string, params query.Params) (T, error) {
	var response T
	if res, err := client.Get(ctx, path, params, nil); err != nil {
		return response, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return response, err
	} else {
		return response, nil
	}
}

type azureClient struct {
	msgraph         rest.RestClient
	resourceManager rest.RestClient
	tenant          azure.Tenant
}

type AzureGraphClient interface {
	GetAzureADOrganization(ctx context.Context, selectCols []string) (*azure.Organization, error)

	ListAzureADGroups(ctx context.Context, params query.GraphParams) <-chan azureResult[azure.Group]
	ListAzureADGroupMembers(ctx context.Context, objectId string, params query.GraphParams) <-chan azureResult[json.RawMessage]
	ListAzureADGroupOwners(ctx context.Context, objectId string, params query.GraphParams) <-chan azureResult[json.RawMessage]
	ListAzureADAppOwners(ctx context.Context, objectId string, params query.GraphParams) <-chan azureResult[json.RawMessage]
	ListAzureADApps(ctx context.Context, params query.GraphParams) <-chan azureResult[azure.Application]
	ListAzureADUsers(ctx context.Context, params query.GraphParams) <-chan azureResult[azure.User]
	ListAzureADRoleAssignments(ctx context.Context, params query.GraphParams) <-chan azureResult[azure.UnifiedRoleAssignment]
	ListAzureADRoles(ctx context.Context, params query.GraphParams) <-chan azureResult[azure.Role]
	ListAzureADServicePrincipalOwners(ctx context.Context, objectId string, params query.GraphParams) <-chan azureResult[json.RawMessage]
	ListAzureADServicePrincipals(ctx context.Context, params query.GraphParams) <-chan azureResult[azure.ServicePrincipal]
	ListAzureDeviceRegisteredOwners(ctx context.Context, objectId string, params query.GraphParams) <-chan azureResult[json.RawMessage]
	ListAzureDevices(ctx context.Context, params query.GraphParams) <-chan azureResult[azure.Device]
	ListAzureADAppRoleAssignments(ctx context.Context, servicePrincipalId string, params query.GraphParams) <-chan azureResult[azure.AppRoleAssignment]
}

type AzureResourceManagerClient interface {
	GetAzureADTenants(ctx context.Context, includeAllTenantCategories bool) (azure.TenantList, error)

	ListRoleAssignmentsForResource(ctx context.Context, resourceId string, filter, tenantId string) <-chan azureResult[azure.RoleAssignment]
	ListAzureADTenants(ctx context.Context, includeAllTenantCategories bool) <-chan azureResult[azure.Tenant]
	ListAzureContainerRegistries(ctx context.Context, subscriptionId string) <-chan azureResult[azure.ContainerRegistry]
	ListAzureWebApps(ctx context.Context, subscriptionId string) <-chan azureResult[azure.WebApp]
	ListAzureManagedClusters(ctx context.Context, subscriptionId string) <-chan azureResult[azure.ManagedCluster]
	ListAzureVMScaleSets(ctx context.Context, subscriptionId string) <-chan azureResult[azure.VMScaleSet]
	ListAzureKeyVaults(ctx context.Context, subscriptionId string, params query.RMParams) <-chan azureResult[azure.KeyVault]
	ListAzureManagementGroups(ctx context.Context, skipToken string) <-chan azureResult[azure.ManagementGroup]
	ListAzureManagementGroupDescendants(ctx context.Context, groupId string, top int32) <-chan azureResult[azure.DescendantInfo]
	ListAzureResourceGroups(ctx context.Context, subscriptionId string, params query.RMParams) <-chan azureResult[azure.ResourceGroup]
	ListAzureSubscriptions(ctx context.Context) <-chan azureResult[azure.Subscription]
	ListAzureVirtualMachines(ctx context.Context, subscriptionId string, params query.RMParams) <-chan azureResult[azure.VirtualMachine]
	ListAzureStorageAccounts(ctx context.Context, subscriptionId string) <-chan azureResult[azure.StorageAccount]
	ListAzureStorageContainers(ctx context.Context, subscriptionId string, resourceGroupName string, saName string, filter string, includeDeleted string, maxPageSize string) <-chan azureResult[azure.StorageContainer]
	ListAzureAutomationAccounts(ctx context.Context, subscriptionId string) <-chan azureResult[azure.AutomationAccount]
	ListAzureLogicApps(ctx context.Context, subscriptionId string, filter string, top int32) <-chan azureResult[azure.LogicApp]
	ListAzureFunctionApps(ctx context.Context, subscriptionId string) <-chan azureResult[azure.FunctionApp]
}

type AzureClient interface {
	AzureGraphClient
	AzureResourceManagerClient

	TenantInfo() azure.Tenant
	CloseIdleConnections()
}

func (s azureClient) TenantInfo() azure.Tenant {
	return s.tenant
}

func (s azureClient) CloseIdleConnections() {
	s.msgraph.CloseIdleConnections()
	s.resourceManager.CloseIdleConnections()
}
