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

type UnifiedRoleEligibilityScheduleInstance struct {
	Entity

	// The status of the role eligibility request.
	// Read-only.
	// Supports $filter (eq, ne).
	Status string `json:"status"`

	// Identifier of the principal that has been granted the role eligibility.
	// Can be a user or a role-assignable group. You can grant only active
	// assignments service principals.
	// Supports $filter (eq, in).
	PrincipalId string `json:"principalId"`

	// Identifier of the unifiedRoleDefinition object that is being assigned to the principal.
	// Supports $filer (eq, in).
	RoleDefinitionId string `json:"roleDefinitionId"`

	// Identifier of the directory object representing the scope of the role eligibility.
	// Either this property or appScopeId is required.
	// The scope of a role eligibility determines the set of resources for which the principal has been granted access.
	// Directory scopes are shared scopes stored in the directory that are understood by multiple applications.
	//
	// Use / for tenant-wide scope.
	// Use appScopeId to limit the scope to an application only.
	//
	// Supports $filter (eq, ne, and on null values).
	DirectoryScopeId string `json:"directoryScopeId,omitempty"`

	// Identifier of the app-specific scope when the role eligibility scope is scoped to an app.
	// The scope of a role eligibility determines the set of resources for which the principal is eligible to access.
	// App scopes are scopes that are defined and understood by this application only.
	//
	// Use / for tenant-wide app scopes.
	// Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units.
	//
	// Supports $filter (eq, ne, and on null values).
	AppScopeId string `json:"appScopeId,omitempty"`
}

type UnifiedRoleEligibilityScheduleInstanceList struct {
	Count    int                                      `json:"@odata.count,omitempty"`    // The total count of all results
	NextLink string                                   `json:"@odata.nextLink,omitempty"` // The URL to use for getting the next set of values.
	Value    []UnifiedRoleEligibilityScheduleInstance `json:"value"`                     // A list of role assignments.
}

type UnifiedRoleEligibilityScheduleInstanceResult struct {
	Error error
	Ok    UnifiedRoleEligibilityScheduleInstance
}
