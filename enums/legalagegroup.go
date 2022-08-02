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

type LegalAgeGroup string

const (
	// The user is considered a minor based on the age-related regulations of their country or region and the
	// administrator of the account has obtained appropriate consent from a parent or guardian.
	LegalAgeGroupMinorWithParentalConsent LegalAgeGroup = "minorWithParentalConsent"

	// The user is considered an adult based on the age-related regulations of their country or region.
	LegalAgeGroupAdult LegalAgeGroup = "adult"

	// The user is from a country or region that has statutory regulations and the user's age is more than the upper
	// limit of kid age and less than the lower limit of adult age as defined by the user's country or region.
	LegalAgeGroupNotAdult LegalAgeGroup = "notAdult"

	// The user is a minor but is from a country or region that has no age-related regulations.
	LegalAgeGroupMinorNoParentalConsentRequired LegalAgeGroup = "minorNoParentalConsentRequired"
)
