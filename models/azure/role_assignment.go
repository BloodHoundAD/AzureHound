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

type RoleAssignmentPropertiesWithScope struct {
	// The principal ID.
	PrincipalId string `json:"principalId,omitempty"`

	// The role definition ID.
	RoleDefinitionId string `json:"roleDefinitionId,omitempty"`

	// The role assignment scope.
	Scope string `json:"scope,omitempty"`
}

type RoleAssignment struct {
	// The role assignment ID.
	Id string `json:"id,omitempty"`

	// The role assignment name.
	Name string `json:"name,omitempty"`

	// The role assignment type.
	Type string `json:"type,omitempty"`

	// Role assignment properties
	Properties RoleAssignmentPropertiesWithScope `json:"properties,omitempty"`
}

func (s RoleAssignment) GetPrincipalId() string {
	return s.Properties.PrincipalId
}
