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

type AutomationAccountProperties struct {
	AutomationHybridServiceUrl string                      `json:"automationHybridServiceUrl,omitempty"`
	CreationTime               string                      `json:"creationTime,omitempty"`
	Description                string                      `json:"description,omitempty"`
	DisableLocalAuth           bool                        `json:"disableLocalAuth,omitempty"`
	Encryption                 EncryptionProperties        `json:"encryption,omitempty"`
	LastModifiedBy             string                      `json:"lastModifiedBy,omitempty"`
	LastModifiedTime           string                      `json:"lastModifiedTime,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection `json:"privateEndpointConnections,omitempty"`
	PublicNetworkAccess        bool                        `json:"publicNetworkAccess,omitempty"`
	Sku                        Sku                         `json:"sku,omitempty"`
	State                      AutomationAccountState      `json:"state,omitempty"`
}

type AutomationAccountState struct {
	Ok          string `json:"Ok,omitempty"`
	Suspended   string `json:"Suspended,omitempty"`
	Unavailable string `json:"Unavailable,omitempty"`
}
