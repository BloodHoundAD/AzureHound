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

// The scope of a role assignment determines the set of resources for which the principal has been granted access.
// An app scope is a scope defined and understood by a specific application, unlike directory scopes which are shared
// scopes stored in the directory and understood by multiple applications.
//
// This may be in both the following principal and scope scenarios:
//
//	A single principal and a single scope
//	Multiple principals and multiple scopes.
type AppScope struct {
	Entity

	// Provides the display name of the app-specific resource represented by the app scope.
	// Provided for display purposes since appScopeId is often an immutable, non-human-readable id.
	// Read-only.
	DisplayName string `json:"display_name,omitempty"`

	// Describes the type of app-specific resource represented by the app scope.
	// Provided for display purposes, so a user interface can convey to the user the kind of app specific resource
	// represented by the app scope.
	// Read-only.
	Type string `json:"type,omitempty"`
}
