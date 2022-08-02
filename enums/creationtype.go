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

package enums

type CreationType string

const (
	// User was created as an external account.
	CreationTypeInvitation CreationType = "Invitation"

	// User was created as a local account for an Azure AD B2C tenant.
	CreationTypeLocalAccount CreationType = "LocalAccount"

	// User was created through self-service sign-up by an internal user using email verification.
	CreationTypeEmailVerified CreationType = "EmailVerified"

	// User was created through self-service sign-up by an external user signing up through a link that is part of a
	// user flow.
	CreationTypeSelfServiceSignUp CreationType = "SelfServiceSignUp"
)
