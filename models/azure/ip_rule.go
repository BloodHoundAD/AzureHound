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

// A rule governing the accessibility of a vault from a specific ip address or ip range.
type IPRule struct {
	// An IPv4 address range in CIDR notation, such as '124.56.78.91' (simple IP address)
	// or '124.56.78.0/24' (all addresses that start with 124.56.78).
	Value string `json:"value,omitempty"`
}
