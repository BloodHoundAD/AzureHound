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

// Specifies settings for a web application.
type WebApplication struct {
	// Home page or landing page of the application.
	HomePageUrl string `json:"homePageUrl,omitempty"`

	// Specifies whether this web application can request tokens using the OAuth 2.0 implicit flow.
	ImplicitGrantSettings ImplicitGrantSettings `json:"implicitGrantSettings,omitempty"`

	// Specifies the URL that will be used by Microsoft's authorization service to logout a user using front-channel,
	// back-channel or SAML logout protocols.
	LogoutUrl string `json:"logoutUrl,omitempty"`

	// Specifies the URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization
	// codes and access tokens are sent.
	RedirectUris []string `json:"redirectUris,omitempty"`
}
