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

// Contains a key credential associated with an application.
// For more detail see https://docs.microsoft.com/en-us/graph/api/resources/keycredential?view=graph-rest-1.0
type KeyCredential struct {
	// Custom key identifier
	// Base64Url encoded.
	CustomKeyIdentifier string `json:"customKeyIdentifier,omitempty"`

	// Friendly name for the key.
	// Optional.
	DisplayName string `json:"displayName,omitempty"`

	// The date and time at which the credential expires.
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time.
	// For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	EndDateTime string `json:"endDateTime,omitempty"`

	// The certificate's raw data in byte array converted to Base64 string;
	// For example, [System.Convert]::ToBase64String($Cert.GetRawCertData()).
	// Base64Url encoded.
	Key []byte `json:"key,omitempty"`

	// The unique identifier (GUID) for the key.
	KeyId uuid.UUID `json:"keyId,omitempty"`

	// The date and time at which the credential becomes valid.The Timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC time.
	// For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	StartDateTime string `json:"startDateTime,omitempty"`

	// The type of key credential; for example, Symmetric.
	Type string `json:"type,omitempty"`

	// A string that describes the purpose for which the key can be used; for example, Verify.
	Usage string `json:"usage,omitempty"`
}
