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

// The return type of the onPremisesExtensionAttributes property of the user object and extensionAttributes property of
// the device object.
// Returns fifteen custom extension attribute properties.
//
// On the user entity and for an onPremisesSyncEnabled user, the source of authority for this set of properties is the
// on-premises Active Directory which is synchronized to Azure AD, and is read-only. For a cloud-only user (where
// onPremisesSyncEnabled is false), these properties can be set during creation or update. If a cloud-only user was
// previously synced from on-premises Active Directory, these properties cannot be managed via the Microsoft Graph API.
// Instead, they can be managed through the Exchange Admin Center or the Exchange Online V2 module in PowerShell.
//
// The extensionAttributes property of the device entity is managed only in Azure AD during device creation or update.
// Note: These extension attributes are also known as Exchange custom attributes 1-15.
type OnPremisesExtensionAttributes struct {
	ExtensionAttribute1  string `json:"extensionAttribute1,omitempty"`
	ExtensionAttribute2  string `json:"extensionAttribute2,omitempty"`
	ExtensionAttribute3  string `json:"extensionAttribute3,omitempty"`
	ExtensionAttribute4  string `json:"extensionAttribute4,omitempty"`
	ExtensionAttribute5  string `json:"extensionAttribute5,omitempty"`
	ExtensionAttribute6  string `json:"extensionAttribute6,omitempty"`
	ExtensionAttribute7  string `json:"extensionAttribute7,omitempty"`
	ExtensionAttribute8  string `json:"extensionAttribute8,omitempty"`
	ExtensionAttribute9  string `json:"extensionAttribute9,omitempty"`
	ExtensionAttribute10 string `json:"extensionAttribute10,omitempty"`
	ExtensionAttribute11 string `json:"extensionAttribute11,omitempty"`
	ExtensionAttribute12 string `json:"extensionAttribute12,omitempty"`
	ExtensionAttribute13 string `json:"extensionAttribute13,omitempty"`
	ExtensionAttribute14 string `json:"extensionAttribute14,omitempty"`
	ExtensionAttribute15 string `json:"extensionAttribute15,omitempty"`
}
