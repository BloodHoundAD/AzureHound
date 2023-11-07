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

	"github.com/bloodhoundad/azurehound/v2/client/config"
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

func (s azureClient) TenantInfo() azure.Tenant {
	return s.tenant
}

func (s azureClient) CloseIdleConnections() {
	s.msgraph.CloseIdleConnections()
	s.resourceManager.CloseIdleConnections()
}

type AzureClient interface {
	GetAzureADApp(ctx context.Context, objectId string, selectCols []string) (*azure.Application, error)
	GetAzureADApps(ctx context.Context, filter, search, orderBy, expand string, selectCols []string, top int32, count bool) (azure.ApplicationList, error)
	GetAzureADDirectoryObject(ctx context.Context, objectId string) (json.RawMessage, error)
	GetAzureADGroup(ctx context.Context, objectId string, selectCols []string) (*azure.Group, error)
	GetAzureADGroupOwners(ctx context.Context, objectId string, filter string, search string, orderBy string, selectCols []string, top int32, count bool) (azure.DirectoryObjectList, error)
	GetAzureADGroups(ctx context.Context, filter, search, orderBy, expand string, selectCols []string, top int32, count bool) (azure.GroupList, error)
	GetAzureADOrganization(ctx context.Context, selectCols []string) (*azure.Organization, error)
	GetAzureADRole(ctx context.Context, roleId string, selectCols []string) (*azure.Role, error)
	GetAzureADRoleAssignment(ctx context.Context, objectId string, selectCols []string) (*azure.UnifiedRoleAssignment, error)
	GetAzureADRoleAssignments(ctx context.Context, filter, search, orderBy, expand string, selectCols []string, top int32, count bool) (azure.UnifiedRoleAssignmentList, error)
	GetAzureADRoles(ctx context.Context, filter, expand string) (azure.RoleList, error)
	GetAzureADServicePrincipal(ctx context.Context, objectId string, selectCols []string) (*azure.ServicePrincipal, error)
	GetAzureADServicePrincipalOwners(ctx context.Context, objectId string, filter string, search string, orderBy string, selectCols []string, top int32, count bool) (azure.DirectoryObjectList, error)
	GetAzureADServicePrincipals(ctx context.Context, filter, search, orderBy, expand string, selectCols []string, top int32, count bool) (azure.ServicePrincipalList, error)
	GetAzureADTenants(ctx context.Context, includeAllTenantCategories bool) (azure.TenantList, error)
	GetAzureADUser(ctx context.Context, objectId string, selectCols []string) (*azure.User, error)
	GetAzureADUsers(ctx context.Context, filter string, search string, orderBy string, selectCols []string, top int32, count bool) (azure.UserList, error)
	GetAzureDevice(ctx context.Context, objectId string, selectCols []string) (*azure.Device, error)
	GetAzureDevices(ctx context.Context, filter, search, orderBy, expand string, selectCols []string, top int32, count bool) (azure.DeviceList, error)
	GetAzureKeyVault(ctx context.Context, subscriptionId, groupName, vaultName string) (*azure.KeyVault, error)
	GetAzureKeyVaults(ctx context.Context, subscriptionId string, top int32) (azure.KeyVaultList, error)
	GetAzureManagementGroup(ctx context.Context, groupId, filter, expand string, recurse bool) (*azure.ManagementGroup, error)
	GetAzureManagementGroups(ctx context.Context) (azure.ManagementGroupList, error)
	GetAzureResourceGroup(ctx context.Context, subscriptionId, groupName string) (*azure.ResourceGroup, error)
	GetAzureResourceGroups(ctx context.Context, subscriptionId string, filter string, top int32) (azure.ResourceGroupList, error)
	GetAzureSubscription(ctx context.Context, objectId string) (*azure.Subscription, error)
	GetAzureSubscriptions(ctx context.Context) (azure.SubscriptionList, error)
	GetAzureVirtualMachine(ctx context.Context, subscriptionId, groupName, vmName, expand string) (*azure.VirtualMachine, error)
	GetAzureVirtualMachines(ctx context.Context, subscriptionId string, statusOnly bool) (azure.VirtualMachineList, error)
	GetAzureStorageAccount(ctx context.Context, subscriptionId, groupName, saName, expand string) (*azure.StorageAccount, error)
	GetAzureStorageAccounts(ctx context.Context, subscriptionId string) (azure.StorageAccountList, error)
	GetResourceRoleAssignments(ctx context.Context, subscriptionId string, filter string, expand string) (azure.RoleAssignmentList, error)
	GetRoleAssignmentsForResource(ctx context.Context, resourceId string, filter string) (azure.RoleAssignmentList, error)
	ListAzureADAppMemberObjects(ctx context.Context, objectId string, securityEnabledOnly bool) <-chan azure.MemberObjectResult
	ListAzureADAppOwners(ctx context.Context, objectId string, filter, search, orderBy string, selectCols []string) <-chan azure.AppOwnerResult
	ListAzureADApps(ctx context.Context, filter, search, orderBy, expand string, selectCols []string) <-chan azure.ApplicationResult
	ListAzureADGroupMembers(ctx context.Context, objectId string, filter, search, orderBy string, selectCols []string) <-chan azure.MemberObjectResult
	ListAzureADGroupOwners(ctx context.Context, objectId string, filter, search, orderBy string, selectCols []string) <-chan azure.GroupOwnerResult
	ListAzureADGroups(ctx context.Context, filter, search, orderBy, expand string, selectCols []string) <-chan azure.GroupResult
	ListAzureADRoleAssignments(ctx context.Context, filter, search, orderBy, expand string, selectCols []string) <-chan azure.UnifiedRoleAssignmentResult
	ListAzureADRoles(ctx context.Context, filter, expand string) <-chan azure.RoleResult
	ListAzureADServicePrincipalOwners(ctx context.Context, objectId string, filter, search, orderBy string, selectCols []string) <-chan azure.ServicePrincipalOwnerResult
	ListAzureADServicePrincipals(ctx context.Context, filter, search, orderBy, expand string, selectCols []string) <-chan azure.ServicePrincipalResult
	ListAzureADTenants(ctx context.Context, includeAllTenantCategories bool) <-chan azure.TenantResult
	ListAzureADUsers(ctx context.Context, filter string, search string, orderBy string, selectCols []string) <-chan azure.UserResult
	ListAzureContainerRegistries(ctx context.Context, subscriptionId string) <-chan azure.ContainerRegistryResult
	ListAzureWebApps(ctx context.Context, subscriptionId string) <-chan azure.WebAppResult
	ListAzureManagedClusters(ctx context.Context, subscriptionId string, statusOnly bool) <-chan azure.ManagedClusterResult
	ListAzureVMScaleSets(ctx context.Context, subscriptionId string, statusOnly bool) <-chan azure.VMScaleSetResult
	ListAzureDeviceRegisteredOwners(ctx context.Context, objectId string, securityEnabledOnly bool) <-chan azure.DeviceRegisteredOwnerResult
	ListAzureDevices(ctx context.Context, filter, search, orderBy, expand string, selectCols []string) <-chan azure.DeviceResult
	ListAzureKeyVaults(ctx context.Context, subscriptionId string, top int32) <-chan azure.KeyVaultResult
	ListAzureManagementGroupDescendants(ctx context.Context, groupId string) <-chan azure.DescendantInfoResult
	ListAzureManagementGroups(ctx context.Context) <-chan azure.ManagementGroupResult
	ListAzureResourceGroups(ctx context.Context, subscriptionId, filter string) <-chan azure.ResourceGroupResult
	ListAzureSubscriptions(ctx context.Context) <-chan azure.SubscriptionResult
	ListAzureVirtualMachines(ctx context.Context, subscriptionId string, statusOnly bool) <-chan azure.VirtualMachineResult
	ListAzureStorageAccounts(ctx context.Context, subscriptionId string) <-chan azure.StorageAccountResult
	ListAzureStorageContainers(ctx context.Context, subscriptionId string, resourceGroupName string, saName string, filter string, includeDeleted string, maxPageSize string) <-chan azure.StorageContainerResult
	ListAzureAutomationAccounts(ctx context.Context, subscriptionId string) <-chan azure.AutomationAccountResult
	ListAzureLogicApps(ctx context.Context, subscriptionId string, filter string, top int32) <-chan azure.LogicAppResult
	ListAzureFunctionApps(ctx context.Context, subscriptionId string) <-chan azure.FunctionAppResult
	ListResourceRoleAssignments(ctx context.Context, subscriptionId string, filter string, expand string) <-chan azure.RoleAssignmentResult
	ListRoleAssignmentsForResource(ctx context.Context, resourceId string, filter string) <-chan azure.RoleAssignmentResult
	ListAzureADAppRoleAssignments(ctx context.Context, servicePrincipal, filter, search, orderBy, expand string, selectCols []string) <-chan azure.AppRoleAssignmentResult
	TenantInfo() azure.Tenant
	CloseIdleConnections()
}
