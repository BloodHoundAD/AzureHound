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

type ResourceGroup struct {
	Entity

	// The location of the resource group. It cannot be changed after the resource group has been created. It must be
	// one of the supported Azure locations.
	Location string `json:"location,omitempty"`

	// The ID of the resource that manages this resource group.
	ManagedBy string `json:"managedBy,omitempty"`

	// The name of the resource group.
	Name string `json:"name,omitempty"`

	// The resource group properties.
	Properties ResourceGroupProperties `json:"properties,omitempty"`

	// The tags attached to the resource group.
	Tags map[string]string `json:"tags,omitempty"`

	// The type of the resource group.
	Type string `json:"type,omitempty"`
}
