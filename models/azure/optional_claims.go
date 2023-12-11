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

// Contains an optional claim associated with an application . The idToken, accessToken, and saml2Token properties of
// the optionalClaims resource is a collection of optionalClaim. If supported by a specific claim, you can also modify
// the behavior of the optionalClaim using the additionalProperties property.
// For more detail see https://docs.microsoft.com/en-us/graph/api/resources/optionalclaim?view=graph-rest-1.0
type OptionalClaim struct {
	// Additional properties of the claim. If a property exists in this collection, it modifies the behavior of the
	// optional claim specified in the name property.
	AdditionalProperties []string `json:"additionalProperties,omitempty"`

	// If the value is true, the claim specified by the client is necessary to ensure a smooth authorization experience
	// for the specific task requested by the end user. The default value is false.
	Essential bool `json:"essential,omitempty"`

	// The name of the optional claim.
	Name string `json:"name,omitempty"`

	// The source (directory object) of the claim.
	// There are predefined claims and user-defined claims from extension properties. If the source value is null, the
	// claim is a predefined optional claim. If the source value is user, the value in the name property is the
	// extension property from the user object.
	Source string `json:"source,omitempty"`
}

// Declares the optional claims requested by an application. An application can configure optional claims to be returned
// in each of three types of tokens (ID token, access token, SAML 2 token) it can receive from the security token
// service. An application can configure a different set of optional claims to be returned in each token type.
//
// Application developers can configure optional claims in their Azure AD apps to specify which claims they want in
// tokens sent to their application by the Microsoft security token service.
// For more detail see https://docs.microsoft.com/en-us/graph/api/resources/optionalclaims?view=graph-rest-1.0
type OptionalClaims struct {
	// The optional claims returned in the JWT ID token.
	IdToken []OptionalClaim `json:"idToken,omitempty"`

	// The optional claims returned in the JWT access token.
	AccessToken []OptionalClaim `json:"accessToken,omitempty"`

	// The optional claims returned in the SAML token.
	Saml2Token []OptionalClaim `json:"saml2Token,omitempty"`
}
