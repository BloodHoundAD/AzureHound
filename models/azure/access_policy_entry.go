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

// An identity that have access to the key vault. All identities in the array must use the same tenant ID as the key
// vault's tenant ID.
type AccessPolicyEntry struct {
	// Application ID of the client making request on behalf of a principal
	ApplicationId string `json:"applicationId,omitempty"`

	// The object ID of a user, service principal or security group in the Azure Active Directory tenant for the vault.
	// The object ID must be unique for the list of access policies.
	ObjectId string `json:"objectId,omitempty"`

	// Permissions the identity has for keys, secrets and certificates.
	Permissions KeyVaultPermissions `json:"permissions,omitempty"`

	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantId string `json:"tenantId,omitempty"`
}
