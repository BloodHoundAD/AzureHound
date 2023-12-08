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

// Specifies whether this web application can request tokens using the OAuth 2.0 implicit flow. Separate properties are
// available to request ID and access tokens as part of the implicit flow. To enable implicit flow, at least one of the
// following properties must be set to true.
type ImplicitGrantSettings struct {
	// Specifies whether this web application can request an ID token using the OAuth 2.0 implicit flow.
	EnableIdTokenIssuance bool `json:"enableIdTokenIssuance,omitempty"`

	// Specifies whether this web application can request an access token using the OAuth 2.0 implicit flow.
	EnableAccessTokenIssuance bool `json:"enableAccessTokenIssuance,omitempty"`
}
