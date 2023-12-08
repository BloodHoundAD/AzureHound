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

import "github.com/gofrs/uuid"

// Represents an application role that can be requested by (and granted to) a client application, or that can be used to
// assign an application to users or groups in a specified role.
//
// To add, update, or remove app roles for an application, update the application for the app or service. App roles on
// the application entity will be available in all tenants where the application is used. To define app roles that are
// only applicable in your tenant (for example, app roles representing custom roles in your instance of a multi-tenant
// application), you can also update the service principal for the app, to add or update app roles to the appRoles
// collection.
//
// With appRoleAssignments, app roles can be assigned to users, groups, or other applications' service principals.
// For more detail see https://docs.microsoft.com/en-us/graph/api/resources/approle?view=graph-rest-1.0
type AppRole struct {
	AllowedMemberTypes []string  `json:"allowedMemberTypes,omitempty"`
	Description        string    `json:"description,omitempty"`
	DisplayName        string    `json:"displayName,omitempty"`
	Id                 uuid.UUID `json:"id,omitempty"`
	IsEnabled          bool      `json:"isEnabled,omitempty"`
	Origin             string    `json:"origin,omitempty"`
	Value              string    `json:"value,omitempty"`
}
