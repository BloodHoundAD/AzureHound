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

// Specifies settings related to VM Guest Patching on Windows.
type WindowsPatchSettings struct {
	// Specifies the mode of VM Guest patch assessment for the IaaS virtual machine.
	// Possible values are:
	// ImageDefault - You control the timing of patch assessments on a virtual machine.
	// AutomaticByPlatform - The platform will trigger periodic patch assessments. The property provisionVMAgent must be true.
	AssessmentMode string `json:"assessmentMode,omitempty"`

	// Enables customers to patch their Azure VMs without requiring a reboot.
	// For enableHotpatching, the 'provisionVMAgent' must be set to true and 'patchMode' must be set to 'AutomaticByPlatform'.
	EnableHotpatching bool `json:"enableHotpatching,omitempty"`

	// Specifies the mode of VM Guest Patching to IaaS virtual machine or virtual machines associated to virtual machine
	// scale set with OrchestrationMode as Flexible.
	// Possible values are:
	// Manual - You control the application of patches to a virtual machine. You do this by applying patches manually inside the VM. In this mode, automatic updates are disabled; the property WindowsConfiguration.enableAutomaticUpdates must be false
	// AutomaticByOS - The virtual machine will automatically be updated by the OS. The property WindowsConfiguration.enableAutomaticUpdates must be true.
	// AutomaticByPlatform - the virtual machine will automatically updated by the platform. The properties provisionVMAgent and WindowsConfiguration.enableAutomaticUpdates must be true
	PatchMode string `json:"patchMode,omitempty"`
}
