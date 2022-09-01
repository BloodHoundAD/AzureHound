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

package enums

type Kind string

const (
	KindAZApp                            Kind = "AZApp"
	KindAZAppMember                      Kind = "AZAppMember"
	KindAZAppOwner                       Kind = "AZAppOwner"
	KindAZDevice                         Kind = "AZDevice"
	KindAZDeviceOwner                    Kind = "AZDeviceOwner"
	KindAZGroup                          Kind = "AZGroup"
	KindAZGroupMember                    Kind = "AZGroupMember"
	KindAZGroupOwner                     Kind = "AZGroupOwner"
	KindAZKeyVault                       Kind = "AZKeyVault"
	KindAZKeyVaultAccessPolicy           Kind = "AZKeyVaultAccessPolicy"
	KindAZKeyVaultContributor            Kind = "AZKeyVaultContributor"
	KindAZKeyVaultOwner                  Kind = "AZKeyVaultOwner"
	KindAZKeyVaultUserAccessAdmin        Kind = "AZKeyVaultUserAccessAdmin"
	KindAZManagementGroup                Kind = "AZManagementGroup"
	KindAZManagementGroupOwner           Kind = "AZManagementGroupOwner"
	KindAZManagementGroupDescendant      Kind = "AZManagementGroupDescendant"
	KindAZManagementGroupUserAccessAdmin Kind = "AZManagementGroupUserAccessAdmin"
	KindAZResourceGroup                  Kind = "AZResourceGroup"
	KindAZResourceGroupOwner             Kind = "AZResourceGroupOwner"
	KindAZResourceGroupUserAccessAdmin   Kind = "AZResourceGroupUserAccessAdmin"
	KindAZRole                           Kind = "AZRole"
	KindAZRoleAssignment                 Kind = "AZRoleAssignment"
	KindAZServicePrincipal               Kind = "AZServicePrincipal"
	KindAZServicePrincipalOwner          Kind = "AZServicePrincipalOwner"
	KindAZSubscription                   Kind = "AZSubscription"
	KindAZSubscriptionOwner              Kind = "AZSubscriptionOwner"
	KindAZSubscriptionUserAccessAdmin    Kind = "AZSubscriptionUserAccessAdmin"
	KindAZTenant                         Kind = "AZTenant"
	KindAZUser                           Kind = "AZUser"
	KindAZVM                             Kind = "AZVM"
	KindAZVMAdminLogin                   Kind = "AZVMAdminLogin"
	KindAZVMAvereContributor             Kind = "AZVMAvereContributor"
	KindAZVMContributor                  Kind = "AZVMContributor"
	KindAZVMOwner                        Kind = "AZVMOwner"
	KindAZVMUserAccessAdmin              Kind = "AZVMUserAccessAdmin"
	KindAZStorageAccount                 Kind = "AZStorageAccount"
	KindAZSAContributor                  Kind = "AZSAContributor"
	KindAZSAOwner                        Kind = "AZSAOwner"
	KindAZSADataReader                   Kind = "AZSADataReader"
	KindAZSAKeyOperator                  Kind = "AZSAKeyOperator"
	KindAZStorageContainer               Kind = "AZStorageContainer"
	KindAZAutomationAccount              Kind = "AZAutomationAccount"
	KindAZAutomationAccountOwner         Kind = "AZAutomationAccountOwner"
	KindAZAutomationAccountContributor   Kind = "AZAutomationAccountContributor"
)
