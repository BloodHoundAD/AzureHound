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

// Specifies Windows operating system settings on the virtual machine.
type WindowsConfiguration struct {
	// Specifies additional base-64 encoded XML formatted information that can be included in the Unattend.xml file,
	// which is used by Windows Setup.
	AdditionalUnattendContent []AdditionalUnattendContent `json:"additionalUnattendContent,omitempty"`

	// Indicates whether Automatic Updates is enabled for the Windows virtual machine. Default value is true.
	// For virtual machine scale sets, this property can be updated and updates will take effect on OS reprovisioning.
	EnableAutomaticUpdates bool `json:"enableAutomaticUpdates,omitempty"`

	// [Preview Feature] Specifies settings related to VM Guest Patching on Windows.
	PatchSettings WindowsPatchSettings `json:"patchSettings,omitempty"`

	// Indicates whether virtual machine agent should be provisioned on the virtual machine.
	// When this property is not specified in the request body, default behavior is to set it to true. This will ensure
	// that VM Agent is installed on the VM so that extensions can be added to the VM later.
	ProvisionVMAgent bool `json:"provisionVMAgent,omitempty"`

	// Specifies the time zone of the virtual machine. e.g. "Pacific Standard Time"
	TimeZone string `json:"timeZone,omitempty"`

	// Specifies the Windows Remote Management listeners. This enables remote Windows PowerShell.
	WinRM WinRMConfiguration `json:"winRM,omitempty"`
}
