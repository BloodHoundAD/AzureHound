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

type BootDiagnotics struct {
	// Whether boot diagnostics should be enabled on the virtual machine.
	Enabled bool `json:"enabled,omitempty"`

	// Uri of the storage account to use for placing the console output and screenshot.
	// If storageUri is not specified while enabling boot diagnostics, managed storage will be used.
	StorageUri string `json:"storageUri,omitempty"`
}
