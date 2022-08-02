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

// Specifies the Microsoft accounts that are supported for the current application.
type SigninAudience string

const (
	// Users with a Microsoft work or school account in my organization’s Azure AD tenant (single-tenant).
	SigninAudienceMyOrg SigninAudience = "AzureADMyOrg"

	// Users with a Microsoft work or school account in any organization’s Azure AD tenant (multi-tenant).
	SigninAudienceMultiOrg SigninAudience = "AzureADMultipleOrgs"

	// Users with a personal Microsoft account, or a work or school account in any organization’s Azure AD tenant.
	SigninAudienceMultiOrgAndAccount SigninAudience = "AzureADandPersonalMicrosoftAccount"

	// Users with a personal Microsoft account only.
	SigninAudienceAccount SigninAudience = "PersonalMicrosoftAccount"
)
