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

type VMIPConfigProperties struct {
	// Specifies an array of references to backend address pools of application gateways. A virtual machine can
	// reference backend address pools of multiple application gateways. Multiple virtual machines cannot use the same
	// application gateway.
	ApplicationGatewayBackendAddressPools []SubResource `json:"applicationGatewayBackendAddressPools,omitempty"`

	// Specifies an array of references to application security group.
	ApplicationSecurityGroups []SubResource `json:"applicationSecurityGroups,omitempty"`

	// Specifies an array of references to backend address pools of load balancers. A virtual machine can reference
	// backend address pools of one public and one internal load balancer. [Multiple virtual machines cannot use the
	// same basic sku load balancer].
	LoadBalancerBackendAddressPools []SubResource `json:"loadBalancerBackendAddressPools,omitempty"`

	// Specifies the primary network interface in case the virtual machine has more than 1 network interface.
	Primary bool `json:"primary,omitempty"`

	// Available from Api-Version 2017-03-30 onwards, it represents whether the specific ipconfiguration is IPv4 or IPv6.
	// Default is taken as IPv4. Possible values are: 'IPv4' and 'IPv6'.
	PrivateIPAddressVersion string `json:"privateIpAddressVersion,omitempty"`

	// The publicIPAddressConfiguration.
	PublicIPAddressConfiguration VMPublicIPConfig `json:"publicIpAddressConfiguration,omitempty"`

	// Specifies the identifier of the subnet.
	Subnet SubResource `json:"subnet,omitempty"`
}
