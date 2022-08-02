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

// Represents a company's privacy profile, which includes a privacy statement URL and a contact person for questions
// regarding the privacy statement.
type PrivacyProfile struct {
	// A valid smtp email address for the privacy statement contact.
	// Not required.
	ContactEmail string `json:"contactEmail,omitempty"`

	// The URL that directs to the company's privacy statement.
	// A valid URL format that begins with http:// or https://.
	// Maximum length is 255 characters.
	// Not required.
	StatementUrl string `json:"statementUrl,omitempty"`
}
