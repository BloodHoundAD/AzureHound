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

// Describes an encryption setting for a disk
type DiskEncryptionSettings struct {
	// Specifies the location of the disk encryption key, which is a Key Vault Secret.
	DiskEncryptionKey KeyVaultSecretReference `json:"diskEncryptionKey,omitempty"`

	// Specifies whether disk encryption should be enabled on the virtual machine.
	Enabled bool `json:"enabled,omitempty"`

	// Specifies the location of the key encryption key in Key Vault.
	KeyEncryptionKey KeyVaultKeyReference `json:"keyEncryptionKey,omitempty"`
}
