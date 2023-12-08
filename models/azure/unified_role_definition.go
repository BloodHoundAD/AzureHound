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

type UnifiedRoleDefinition struct {
	Entity

	// The description for the unifiedRoleDefinition.
	// Read-only when isBuiltIn is true.
	Description string `json:"description,omitempty"`

	// The display name for the unifiedRoleDefinition.
	// Read-only when isBuiltIn is true.
	// Required.
	// Supports $filter (eq, in).
	DisplayName string `json:"displayName,omitempty"`

	// Flag indicating whether the role definition is part of the default set included in
	// Azure Active Directory (Azure AD) or a custom definition.
	// Read-only.
	// Supports $filter (eq, in).
	IsBuiltIn bool `json:"isBuiltIn,omitempty"`

	// Flag indicating whether the role is enabled for assignment.
	// If false the role is not available for assignment.
	// Read-only when isBuiltIn is true.
	IsEnabled bool `json:"isEnabled,omitempty"`

	// List of the scopes or permissions the role definition applies to.
	// Currently only / is supported.
	// Read-only when isBuiltIn is true.
	// DO NOT USE. This will be deprecated soon. Attach scope to role assignment.
	ResourceScopes []string `json:"resourceScopes,omitempty"`

	// List of permissions included in the role.
	// Read-only when isBuiltIn is true.
	// Required.
	RolePermisions []UnifiedRolePermission `json:"rolePermisions,omitempty"`

	// Custom template identifier that can be set when isBuiltIn is false but is read-only when isBuiltIn is true.
	// This identifier is typically used if one needs an identifier to be the same across different directories.
	TemplateId string `json:"templateId,omitempty"`

	// Indicates version of the role definition.
	// Read-only when isBuiltIn is true.
	Version string `json:"version,omitempty"`
}
