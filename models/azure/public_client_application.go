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

// Specifies settings for non-web app or non-web API (for example, mobile or other public clients such as an installed
// application running on a desktop device).
type PublicClientApplication struct {
	// Specifies the URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization
	// codes and access tokens are sent.
	RedirectUris []string `json:"redirectUris,omitempty"`
}
