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

// Specifies settings related to VM Guest Patching on Linux.
type LinuxPatchSettings struct {
	// Specifies the mode of VM Guest Patch Assessment for the IaaS virtual machine.
	// Possible values are:
	// ImageDefault - You control the timing of patch assessments on a virtual machine.
	// AutomaticByPlatform - The platform will trigger periodic patch assessments. The property provisionVMAgent must be true.
	AssessmentMode string `json:"assessmentMode,omitempty"`

	// Specifies the mode of VM Guest Patching to IaaS virtual machine or virtual machines associated to virtual machine scale set with OrchestrationMode as Flexible.
	// Possible values are:
	// ImageDefault - The virtual machine's default patching configuration is used.
	// AutomaticByPlatform - The virtual machine will be automatically updated by the platform. The property provisionVMAgent must be true
	PatchMode string `json:"patchMode,omitempty"`
}
