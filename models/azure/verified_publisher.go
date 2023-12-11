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

// Represents the verified publisher of the application.
// For more detail see https://docs.microsoft.com/en-us/graph/api/resources/verifiedpublisher?view=graph-rest-1.0
type VerifiedPublisher struct {
	// The verified publisher name from the app publisher's Partner Center account.
	DisplayName string `json:"displayName,omitempty"`

	// The ID of the verified publisher from the app publisher's Partner Center account.
	VerifiedPublisherId string `json:"verifiedPublisherId,omitempty"`

	// The timestamp when the verified publisher was first added or most recently updated.
	AddedDateTime string `json:"addedDateTime,omitempty"`
}
