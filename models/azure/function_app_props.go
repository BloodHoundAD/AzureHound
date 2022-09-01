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

type FunctionAppProperties struct {
	Name                        string     `json:"name,omitempty"`
	State                       string     `json:"state,omitempty"`
	Hostnames                   []string   `json:"hostNames,omitempty"`
	SelfLink                    string     `json:"selfLink,omitempty"`
	Enabled                     bool       `json:"enabled,omitempty"`
	AdminEnabled                bool       `json:"adminEnabled,omitempty"`
	EnabledHostnames            []string   `json:"enabledHostnames,omitempty"`
	ComputeMode                 string     `json:"computeMode,omitempty"`
	ServerFarmId                string     `json:"serverFarmId,omitempty"`
	HyperV                      bool       `json:"hyperV,omitempty"`
	LastModifiedTimeUTC         string     `json:"lastModifiedTimeUtc,omitempty"`
	StorageRecoveryDefaultState string     `json:"storageRecoveryDefaultState,omitempty"`
	ContentAvailabilityState    string     `json:"contentAvailabilityState,omitempty"`
	RuntimeAvailabilityState    string     `json:"runtimeAvailabilityState,omitempty"`
	VnetRouteAllEnabled         bool       `json:"vnetRouteAllEnabled,omitempty"`
	ContainerAllocationSubnet   string     `json:"containerAllocationSubnet,omitempty"`
	VnetContentShareEnabled     bool       `json:"vnetContentShareEnabled,omitempty"`
	Kind                        string     `json:"kind,omitempty"`
	InboundIPAddress            string     `json:"inboundIpAddress,omitempty"`
	PossibleInboundIpAddresses  string     `json:"possibleInboundIpAddresses,omitempty"`
	FtpUsername                 string     `json:"ftpUsername,omitempty"`
	FtpsHostName                string     `json:"ftpsHostName,omitempty"`
	OutboundIpAddresses         string     `json:"outboundIpAddresses,omitempty"`
	PossibleOutboundIpAddresses string     `json:"possibleOutboundIpAddresses,omitempty"`
	HttpsOnly                   bool       `json:"httpsOnly,omitempty"`
	PrivateEndpointConnections  string     `json:"privateEndpointConnections,omitempty"`
	PublicNetworkAccess         string     `json:"publicNetworkAccess,omitempty"`
	VirtualNetworkSubnetId      string     `json:"virtualNetworkSubnetId,omitempty"`
	KeyVaultReferenceIdentity   string     `json:"keyVaultReferenceIdentity,omitempty"`
	SiteConfig                  SiteConfig `json:"siteConfig,omitempty"`
}
