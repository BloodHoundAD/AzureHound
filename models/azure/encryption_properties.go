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

type EncryptionProperties struct {
	Identity           ManagedIdentity         `json:"identity,omitempty"`
	KeySource          EncryptionKeySourceType `json:"keySource,omitempty"`
	KeyVaultProperties KeyVaultProperties      `json:"keyVaultProperties,omitempty"`
}

type EncryptionKeySourceType struct {
	Automation string `json:"Microsoft.Automation,omitempty"`
	Keyvault   string `json:"Microsoft.Keyvault,omitempty"`
}

type KeyVaultProperties struct {
	KeyName     string `json:"keyName,omitempty"`
	KeyVersion  string `json:"keyVersion,omitempty"`
	KeyvaultUri string `json:"keyvaultUri,omitempty"`
}
