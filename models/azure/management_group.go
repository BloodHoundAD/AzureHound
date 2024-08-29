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

type ManagementGroup struct {
	Entity

	// The name of the management group. E.g. 00000000-0000-0000-0000-000000000000
	Name string `json:"name,omitempty"`

	// The properties of the management group.
	Properties ManagementGroupProperties `json:"properties,omitempty"`

	// The type of resource: "Microsoft.Management/managementGroups"
	Type string `json:"type,omitempty"`
}
