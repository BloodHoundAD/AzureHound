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

// Specifies the Security profile settings for the virtual machine or virtual machine scale set.
type SecurityProfile struct {
	// This property can be used by user in the request to enable or disable the Host Encryption for the virtual machine
	// or virtual machine scale set. This will enable the encryption for all the disks including Resource/Temp disk at
	// host itself.
	// Default: The Encryption at host will be disabled unless this property is set to true for the resource.
	EncryptionAtHost bool `json:"encryptionAtHost,omitempty"`

	// Specifies the SecurityType of the virtual machine. It is set as TrustedLaunch to enable UefiSettings.
	// Default: UefiSettings will not be enabled unless this property is set as TrustedLaunch.
	SecurityType string `json:"securityType,omitempty"`

	// Specifies the security settings like secure boot and vTPM used while creating the virtual machine.
	// Minimum api-version: 2020-12-01
	UefiSettings UefiSettings `json:"uefiSettings,omitempty"`
}
