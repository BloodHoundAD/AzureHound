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

import "github.com/gofrs/uuid"

// Specifies settings for an application that implements a web API.
// For more detail see https://docs.microsoft.com/en-us/graph/api/resources/apiapplication?view=graph-rest-1.0
type ApiApplication struct {
	// When true, allows an application to use claims mapping without specifying a custom signing key.
	AcceptMappedClaims bool `json:"acceptMappedClaims,omitempty"`

	// Used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app.
	// If you set the appID of the client app to this value, the user only consents once to the client app. Azure AD
	// knows that consenting to the client means implicitly consenting to the web API and automatically provisions
	// service principals for both APIs at the same time. Both the client and the web API app must be registered in the
	// same tenant.
	KnownClientApplications []uuid.UUID `json:"knownClientApplications,omitempty"`

	// The definition of the delegated permissions exposed by the web API represented by this application registration.
	// These delegated permissions may be requested by a client application, and may be granted by users or
	// administrators during consent. Delegated permissions are sometimes referred to as OAuth 2.0 scopes.
	OAuth2PermissionScopes []PermissionScope `json:"oauth2PermissionScopes,omitempty"`

	// Lists the client applications that are pre-authorized with the specified delegated permissions to access this
	// application's APIs. Users are not required to consent to any pre-authorized application (for the permissions
	// specified). However, any additional permissions not listed in preAuthorizedApplications (requested through
	// incremental consent for example) will require user consent.
	PreAuthorizedApplications []PreAuthorizedApplication `json:"preAuthorizedApplications,omitempty"`

	// Specifies the access token version expected by this resource.
	// This changes the version and format of the JWT produced independent of the endpoint or client used to request the
	// access token.
	//
	// The endpoint used, v1.0 or v2.0, is chosen by the client and only impacts the version of id_tokens. Resources
	// need to explicitly configure requestedAccessTokenVersion to indicate the supported access token format.
	//
	// Possible values for requestedAccessTokenVersion are 1, 2, or null. If the value is null, this defaults to 1,
	// which corresponds to the v1.0 endpoint.
	//
	// If signInAudience on the application is configured as AzureADandPersonalMicrosoftAccount, the value for this
	//property must be 2
	RequestedAccessTokenVersion int32 `json:"requestedAccessTokenVersion,omitempty"`
}
