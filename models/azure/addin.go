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

// Defines custom behavior that a consuming service can use to call an app in specific contexts.For example, applications
// that can render file streams may configure addIns for its "FileHandler" functionality. This will let services like
// Microsoft 365 call the application in the context of a document the user is working on.
type AddIn struct {
	Id         uuid.UUID  `json:"id,omitempty"`
	Properties []KeyValue `json:"properties,omitempty"`
	Type       string     `json:"type,omitempty"`
}
