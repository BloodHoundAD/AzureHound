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
	"fmt"

	"github.com/bloodhoundad/azurehound/v2/client/config"
	"github.com/bloodhoundad/azurehound/v2/client/query"
	"github.com/bloodhoundad/azurehound/v2/client/rest"
	"github.com/bloodhoundad/azurehound/v2/models/azure"
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

type azureClient struct {
	msgraph         rest.RestClient
	resourceManager rest.RestClient
	tenant          azure.Tenant
}

type AzureGraphClient interface {
	GetAzureADApps(ctx context.Context, params query.GraphParams) (azure.ApplicationList, error)
	GetAzureADGroupOwners(ctx context.Context, objectId string, params query.GraphParams) (azure.DirectoryObjectList, error)
	GetAzureADGroups(ctx context.Context, params query.GraphParams) (azure.GroupList, error)
	GetAzureADOrganization(ctx context.Context, selectCols []string) (*azure.Organization, error)
	GetAzureADRoles(ctx context.Context, filter string) (azure.RoleList, error)
	GetAzureADRoleAssignments(ctx context.Context, params query.GraphParams) (azure.UnifiedRoleAssignmentList, error)
	GetAzureADServicePrincipalOwners(ctx context.Context, objectId string, params query.GraphParams) (azure.DirectoryObjectList, error)
	GetAzureADServicePrincipals(ctx context.Context, params query.GraphParams) (azure.ServicePrincipalList, error)
	GetAzureADUsers(ctx context.Context, params query.GraphParams) (azure.UserList, error)
	GetAzureDevices(ctx context.Context, params query.GraphParams) (azure.DeviceList, error)
	GetAzureDeviceRegisteredOwners(ctx context.Context, objectId string, params query.GraphParams) (azure.DirectoryObjectList, error)
	GetAzureADAppRoleAssignments(ctx context.Context, servicePrincipalId string, params query.GraphParams) (azure.AppRoleAssignmentList, error)

	// https://learn.microsoft.com/en-us/graph/api/group-list-members?view=graph-rest-beta
	GetAzureADGroupMembers(ctx context.Context, objectId string, params query.GraphParams) (azure.MemberObjectList, error)
	ListAzureADGroupMembers(ctx context.Context, objectId string, params query.GraphParams) <-chan azure.MemberObjectResult

	ListAzureADUsers(ctx context.Context, params query.GraphParams) <-chan azure.UserResult
	ListAzureADApps(ctx context.Context, params query.GraphParams) <-chan azure.ApplicationResult
	ListAzureADAppOwners(ctx context.Context, objectId string, params query.GraphParams) <-chan azure.AppOwnerResult
	ListAzureADGroups(ctx context.Context, params query.GraphParams) <-chan azure.GroupResult
	ListAzureADRoleAssignments(ctx context.Context, params query.GraphParams) <-chan azure.UnifiedRoleAssignmentResult
	ListAzureADGroupOwners(ctx context.Context, objectId string, params query.GraphParams) <-chan azure.GroupOwnerResult
	ListAzureADRoles(ctx context.Context, filter string) <-chan azure.RoleResult
	ListAzureADServicePrincipalOwners(ctx context.Context, objectId string, params query.GraphParams) <-chan azure.ServicePrincipalOwnerResult
	ListAzureADServicePrincipals(ctx context.Context, params query.GraphParams) <-chan azure.ServicePrincipalResult
	ListAzureDeviceRegisteredOwners(ctx context.Context, objectId string, params query.GraphParams) <-chan azure.DeviceRegisteredOwnerResult
	ListAzureDevices(ctx context.Context, params query.GraphParams) <-chan azure.DeviceResult
	ListAzureADAppRoleAssignments(ctx context.Context, servicePrincipal string, params query.GraphParams) <-chan azure.AppRoleAssignmentResult
}

type AzureResourceManagerClient interface {
	GetAzureADTenants(ctx context.Context, includeAllTenantCategories bool) (azure.TenantList, error)
	GetAzureKeyVaults(ctx context.Context, subscriptionId string, params query.RMParams) (azure.KeyVaultList, error)
	GetAzureManagementGroups(ctx context.Context, skipToken string) (azure.ManagementGroupList, error)
	GetAzureResourceGroups(ctx context.Context, subscriptionId string, params query.RMParams) (azure.ResourceGroupList, error)
	GetAzureSubscriptions(ctx context.Context) (azure.SubscriptionList, error)
	GetAzureVirtualMachines(ctx context.Context, subscriptionId string, params query.RMParams) (azure.VirtualMachineList, error)
	GetAzureStorageAccounts(ctx context.Context, subscriptionId string) (azure.StorageAccountList, error)
	GetRoleAssignmentsForResource(ctx context.Context, resourceId string, filter, tenantId string) (azure.RoleAssignmentList, error)
	GetAzureContainerRegistries(ctx context.Context, subscriptionId string) (azure.ContainerRegistryList, error)
	GetAzureWebApps(ctx context.Context, subscriptionId string) (azure.WebAppList, error)
	GetAzureVMScaleSets(ctx context.Context, subscriptionId string) (azure.VMScaleSetList, error)
	GetAzureStorageContainers(ctx context.Context, subscriptionId string, resourceGroupName string, saName string, filter string, includeDeleted string, maxPageSize string) (azure.StorageContainerList, error)
	GetAzureAutomationAccounts(ctx context.Context, subscriptionId string) (azure.AutomationAccountList, error)
	GetAzureLogicApps(ctx context.Context, subscriptionId string, filter string, top int32) (azure.LogicAppList, error)
	GetAzureFunctionApps(ctx context.Context, subscriptionId string) (azure.FunctionAppList, error)
	GetAzureManagementGroupDescendants(ctx context.Context, groupId string, top int32) (azure.DescendantInfoList, error)

	ListAzureADTenants(ctx context.Context, includeAllTenantCategories bool) <-chan azure.TenantResult
	ListAzureContainerRegistries(ctx context.Context, subscriptionId string) <-chan azure.ContainerRegistryResult
	ListAzureWebApps(ctx context.Context, subscriptionId string) <-chan azure.WebAppResult
	ListAzureManagedClusters(ctx context.Context, subscriptionId string) <-chan azure.ManagedClusterResult
	ListAzureVMScaleSets(ctx context.Context, subscriptionId string) <-chan azure.VMScaleSetResult
	ListAzureKeyVaults(ctx context.Context, subscriptionId string, params query.RMParams) <-chan azure.KeyVaultResult
	ListAzureManagementGroups(ctx context.Context, skipToken string) <-chan azure.ManagementGroupResult
	ListAzureResourceGroups(ctx context.Context, subscriptionId string, params query.RMParams) <-chan azure.ResourceGroupResult
	ListAzureSubscriptions(ctx context.Context) <-chan azure.SubscriptionResult
	ListAzureVirtualMachines(ctx context.Context, subscriptionId string, params query.RMParams) <-chan azure.VirtualMachineResult
	ListAzureStorageAccounts(ctx context.Context, subscriptionId string) <-chan azure.StorageAccountResult
	ListAzureStorageContainers(ctx context.Context, subscriptionId string, resourceGroupName string, saName string, filter string, includeDeleted string, maxPageSize string) <-chan azure.StorageContainerResult
	ListAzureAutomationAccounts(ctx context.Context, subscriptionId string) <-chan azure.AutomationAccountResult
	ListAzureLogicApps(ctx context.Context, subscriptionId string, filter string, top int32) <-chan azure.LogicAppResult
	ListAzureFunctionApps(ctx context.Context, subscriptionId string) <-chan azure.FunctionAppResult
	ListRoleAssignmentsForResource(ctx context.Context, resourceId string, filter, tenantId string) <-chan azure.RoleAssignmentResult
	ListAzureManagementGroupDescendants(ctx context.Context, groupId string, top int32) <-chan azure.DescendantInfoResult
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
