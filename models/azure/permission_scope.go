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

// Represents the definition of a delegated permission.
//
// Delegated permissions can be requested by client applications needing an access token to the API which defined the
// permissions. Delegated permissions can be requested dynamically, using the scopes parameter in an authorization request
// to the Microsoft identity platform, or statically, through the requiredResourceAccess collection on the application
// object.
// For more detail see https://docs.microsoft.com/en-us/graph/api/resources/permissionscope?view=graph-rest-1.0
type PermissionScope struct {
	// A description of the delegated permissions, intended to be read by an administrator granting the permission on
	// behalf of all users. This text appears in tenant-wide admin consent experiences.
	AdminConsentDescription string `json:"adminConsentDescription,omitempty"`

	// The permission's title, intended to be read by an administrator granting the permission on behalf of all users.
	AdminConsentDisplayName string `json:"adminConsentDisplayName,omitempty"`

	// Unique delegated permission identifier inside the collection of delegated permissions defined for a resource
	// application.
	Id uuid.UUID `json:"id,omitempty"`

	// When creating or updating a permission, this property must be set to true (which is the default). To delete a
	// permission, this property must first be set to false. At that point, in a subsequent call, the permission may be
	// removed.
	IsEnabled bool `json:"isEnabled,omitempty"`

	// Specifies whether this delegated permission should be considered safe for non-admin users to consent to on behalf
	// of themselves, or whether an administrator should be required for consent to the permissions. This will be the
	// default behavior, but each customer can choose to customize the behavior in their organization (by allowing,
	// restricting or limiting user consent to this delegated permission.)
	Type string `json:"type,omitempty"`

	// A description of the delegated permissions, intended to be read by a user granting the permission on their own
	// behalf. This text appears in consent experiences where the user is consenting only on behalf of themselves.
	UserConsentDescription string `json:"userConsentDescription,omitempty"`

	// A title for the permission, intended to be read by a user granting the permission on their own behalf. This text
	// appears in consent experiences where the user is consenting only on behalf of themselves.
	UserConsentDisplayName string `json:"userConsentDisplayName,omitempty"`

	// Specifies the value to include in the scp (scope) claim in access tokens.
	// Must not exceed 120 characters in length.
	// Allowed characters are : ! # $ % & ' ( ) * + , - . / : ; < = > ? @ [ ] ^ + _ ` { | } ~, as well as characters in
	// the ranges 0-9, A-Z and a-z.
	// Any other character, including the space character, are not allowed.
	// May not begin with ..
	Value string `json:"value,omitempty"`
}
