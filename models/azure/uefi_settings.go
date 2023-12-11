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

// Specifies the security settings like secure boot and vTPM used while creating the virtual machine.
// Minimum api-version: 2020-12-01
type UefiSettings struct {
	// Specifies whether secure boot should be enabled on the virtual machine.
	SecureBootEnabled bool `json:"secureBootEnabled,omitempty"`

	// Specifies whether vTPM should be enabled on the virtual machine.
	VTpmEnabled bool `json:"vTpmEnabled,omitempty"`
}
