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

// Represents directory synchronization errors for the user, group and orgContact resources when synchronizing
// on-premises directories to Azure Active Directory.
type OnPremisesProvisioningError struct {
	// Category of the provisioning error. Note: Currently, there is only one possible value.
	// Possible value: PropertyConflict - indicates a property value is not unique. Other objects contain the same value
	// for the property.
	Category string `json:"category,omitempty"`

	// The date and time at which the error occurred.
	OccurredDateTime string `json:"occurredDateTime,omitempty"`

	// Name of the directory property causing the error. Current possible values: UserPrincipalName or ProxyAddress.
	PropertyCausingError string `json:"propertyCausingError,omitempty"`

	// Value of the property causing the error.
	Value string `json:"value,omitempty"`
}
