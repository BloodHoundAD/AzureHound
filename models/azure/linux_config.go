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

// Specifies the Linux operating system settings on the virtual machine.
// For a list of supported Linux distributions, see Linux on Azure-Endorsed Distributions.
type LinuxConfiguration struct {
	// Specifies whether password authentication should be disabled.
	DisablePasswordAuthentication bool `json:"disablePasswordAuthentication,omitempty"`

	// [Preview Feature] Specifies settings related to VM Guest Patching on Linux.
	PatchSettings LinuxPatchSettings `json:"patchSettings,omitempty"`

	// Indicates whether virtual machine agent should be provisioned on the virtual machine.
	// When this property is not specified in the request body, default behavior is to set it to true. This will ensure
	// that VM Agent is installed on the VM so that extensions can be added to the VM later.
	ProvisionVMAgent bool `json:"provisionVMAgent,omitempty"`

	// Specifies the ssh key configuration for a Linux OS.
	Ssh SshConfiguration `json:"ssh,omitempty"`
}
