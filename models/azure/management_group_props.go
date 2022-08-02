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

// The properties of the management group.
type ManagementGroupProperties struct {
	// The list of children.
	Children []ManagementGroupChildInfo `json:"children,omitempty"`

	// The details of the management group.
	Details ManagementGroupDetails `json:"details,omitempty"`

	// The friendly name of the management group.
	DisplayName string `json:"displayName,omitempty"`

	// The Azure AD Tenant ID associated with the management group. E.g. 00000000-0000-0000-0000-000000000000
	TenantId string `json:"tenantId,omitempty"`
}
