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

// Enables or disables a capability on the virtual machine or virtual machine scale set.
type AdditionalCapabilities struct {
	// The flag that enables or disables hibernation capability on the VM.
	HibernationEnabled bool `json:"hibernationEnabled,omitempty"`

	// The flag that enables or disables a capability to have one or more managed data disks with UltraSSD_LRS storage
	// account type on the VM or VMSS. Managed disks with storage account type UltraSSD_LRS can be added to a virtual
	// machine or virtual machine scale set only if this property is enabled.
	UltraSSDEnabled bool `json:"ultraSSDEnabled,omitempty"`
}
