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

type LocaleInfo struct {
	// A locale code for the user, which includes the user's perferred language and country/region as defined
	// in ISO 639-1 and ISO 3166-1 alpha-2. E.g. "en-us"
	Locale string `json:"locale,omitempty"`

	// A name representing the user's locale in natural language. E.g. "English (United States)"
	DisplayName string `json:"displayName,omitempty"`
}
