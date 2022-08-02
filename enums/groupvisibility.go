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

// Specifies the group join policy and group content visibility for groups.
type GroupVisibility string

const (
	// Owner permission is needed to join the group.
	// Non-members cannot view the contents of the group.
	GroupVisibilityPrivate GroupVisibility = "Private"

	// Anyone can join the group without needing owner permission.
	// Anyone can view the contents of the group.
	GroupVisibilityPublic GroupVisibility = "Public"

	// Owner permission is needed to join the group.
	// Non-members cannot view the contents of the group.
	// Non-members cannot see the members of the group.
	// Administrators (global, company, user, and helpdesk) can view the membership of the group.
	// The group appears in the global address book (GAL).
	GroupVisibilityHidden GroupVisibility = "Hiddenmembership"
)
