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

// Specifies the hardware settings for the virtual machine.
type HardwareProfile struct {

	// Specifies the size of the virtual machine.
	//
	// Recommended way to get the list of available sizes is using the API:
	// - List all available virtual machine sizes in an availability set
	// - List all available virtual machine sizes in a region
	// - List all available virtual machine sizes for resizing.
	//
	// For more information about virtual machine sizes, see Sizes for virtual machines.
	//
	// The available VM sizes depend on region and availability set.
	VMSize string `json:"vmSize,omitempty"`

	// Specifies the properties for customizing the size of the virtual machine. Minimum api-version: 2021-07-01.
	//
	// This feature is still in preview mode and is not supported for VirtualMachineScaleSet.
	// Please follow the instructions in VM Customization for more details.
	VMSizeProperties VMSizeProperties `json:"vmSizeProperties,omitempty"`
}
