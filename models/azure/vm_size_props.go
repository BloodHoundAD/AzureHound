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

// Specifies VM Size Property settings on the virtual machine.
type VMSizeProperties struct {
	// Specifies the number of vCPUs available for the VM.
	// When this property is not specified in the request body the default behavior is to set it to the value of vCPUs
	// available for that VM size exposed in api response of List all available virtual machine sizes in a region.
	VCPUsAvailable int `json:"vCPUsAvailable,omitempty"`

	// Specifies the vCPU to physical core ratio.
	// When this property is not specified in the request body the default behavior is set to the value of vCPUsPerCore
	// for the VM Size exposed in api response of List all available virtual machine sizes in a region.
	// Setting this property to 1 also means that hyper-threading is disabled.
	VCPUsPerCore int `json:"vCPUsPerCore,omitempty"`
}
