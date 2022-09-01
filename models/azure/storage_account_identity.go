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

package azure

import "github.com/bloodhoundad/azurehound/enums"

// Identity for the virtual machine.
type StorageAccountIdentity struct {
	// The principal id of the virtual machine identity. The property will only be provided for a system assigned
	// identity.
	PrincipalId string `json:"principalId"`

	// The tenant id associated with the virtual machine. This property will only be provided for a system assigned
	// identity.
	TenantId string `json:"tenantId"`

	// The type of identity used for the virtual machine.
	Type enums.Identity `json:"type"`

	// The list of user identities associated with the Virtual Machine. The user identity dictionary key references will be
	// ARM resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'
	UserAssignedIdentities map[string]UserAssignedIdentity `json:"userAssignedIdentities"`
}
