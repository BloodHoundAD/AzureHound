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

// Contains the password profile associated with a user.
type PasswordProfile struct {
	// true if the user must change her password on the next login; otherwise false. If not set, default is false.
	// NOTE: For Azure B2C tenants, set to false and instead use custom policies and user flows to force password reset
	// at first sign in.
	ForceChangePasswordNextSignIn bool `json:"forceChangePasswordNextSignIn,omitempty"`

	// If true, at next sign-in, the user must perform a multi-factor authentication (MFA) before being forced to change
	// their password. The behavior is identical to forceChangePasswordNextSignIn except that the user is required to
	// first perform a multi-factor authentication before password change. After a password change, this property will
	// be automatically reset to false. If not set, default is false.
	ForceChangePasswordNextSignInWithMfa bool `json:"forceChangePasswordNextSignInWithMfa,omitempty"`

	// The password for the user. This property is required when a user is created.
	// It can be updated, but the user will be required to change the password on the next login.
	// The password must satisfy minimum requirements as specified by the user’s passwordPolicies property.
	// By default, a strong password is required.
	Password string `json:"password,omitempty"`
}
