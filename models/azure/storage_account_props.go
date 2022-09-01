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

type StorageAcccountProperties struct {
	DnsEndpointType              string                          `json:"dnsEndpointType,omitempty"`
	DefaultToOAuthAuthentication bool                            `json:"defaultToOAuthAuthentication,omitempty"`
	PublicNetworkAccess          string                          `json:"availabilitySet,omitempty"`
	KeyCreationTime              string                          `json:"keyCreationTime,omitempty"`
	AllowCrossTenantReplication  bool                            `json:"allowCrossTenantReplication,omitempty"`
	PrivateEndpointConnections   []PrivateEndpointConnectionItem `json:"privateEndpointConnections"`
	MinimumTlsVersion            string                          `json:"minimumTlsVersion,omitempty"`
	AllowBlobPublicAccess        bool                            `json:"allowBlobPublicAccess,omitempty"`
	AllowSharedKeyAccess         bool                            `json:"allowSharedKeyAccess,omitempty"`
	NetworkAcls                  NetworkRuleSet                  `json:"networkAcls"`
	SupportsHttpsTrafficOnly     bool                            `json:"supportsHttpsTrafficOnly,omitempty"`
	Encryption                   string                          `json:"encryption,omitempty"`
	AccessTier                   string                          `json:"accessTier,omitempty"`
	ProvisioningState            string                          `json:"provisioningState,omitempty"`
	creationTime                 string                          `json:"creationTime,omitempty"`
	primaryEndpoints             SAPrimaryEndpoints              `json:"primaryEndpoints,omitempty"`
	primaryLocation              string                          `json:"primaryLocation,omitempty"`
	statusOfPrimary              string                          `json:"statusOfPrimary,omitempty"`
}
