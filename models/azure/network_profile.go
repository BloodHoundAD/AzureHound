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

// Specifies the network interfaces or the networking configuration of the virtual machine.
type NetworkProfile struct {
	// Specifies the Microsoft.Network API version used when creating networking resources in the Network Interface
	// Configurations.
	NetworkApiVersion string `json:"networkApiVersion,omitempty"`

	// Specifies the networking configurations that will be used to create the virtual machine networking resources.
	NetworkInterfaceConfigurations []VirtualMachineNetworkInterfaceConfiguration `json:"networkInterfaceConfigurations,omitempty"`

	// Specifies the list of resource Ids for the network interfaces associated with the virtual machine.
	NetworkInterfaces []NetworkInterfaceReference `json:"networkInterfaces,omitempty"`
}
