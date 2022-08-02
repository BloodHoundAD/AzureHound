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

// Specifies the legal age group rule that applies to users of the app.
type LegalAgeGroupRule string

const (
	// Enforces the legal minimum This means parental consent is required for minors in the EU and Korea.
	//
	// Default
	LegalAgeGroupRuleAllow LegalAgeGroupRule = "Allow"

	// Enforces the user to specify date of birth to comply with COPPA rules.
	LegalAgeGroupRuleRequireConsentForPrivacyServices LegalAgeGroupRule = "RequireConsentForPrivacyServices"

	// Requires parental consent for ages below 18, regardless of country minor rules.
	LegalAgeGroupRuleRequireConsentForMinors LegalAgeGroupRule = "RequireConsentForMinors"

	// Requires parental consent for ages below 14, regardless of country minor rules.
	LegalAgeGroupRuleRequireConsentForKids LegalAgeGroupRule = "RequireConsentForKids"

	// Blocks minors from using the app.
	LegalAgeGroupRuleBlockMinors LegalAgeGroupRule = "BlockMinors"
)
