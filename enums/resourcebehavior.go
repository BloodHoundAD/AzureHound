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

// Specifies group behaviors for a Microsoft 365 group
type ResourceBehavior string

const (
	// Only group members can post conversations to the group.
	// If unset ny user in the organization can post conversations to the group.
	ResourceBehaviorAllowOnlyMembersToPost ResourceBehavior = "AllowOnlyMembersToPost"

	// This group is hidden in Outlook experiences.
	// If unset all groups are visible and discoverable in Outlook experiences.
	ResourceBehaviorHideGroupInOutlook ResourceBehavior = "HideGroupInOutlook"

	// Group members are subscribed to receive group conversations.
	// If unset Group members do not receive group conversations.
	ResourceBehaviorSubscribeNewGroupMembers ResourceBehavior = "SubscribeNewGroupMembers"

	// Welcome emails are not sent to new members.
	// If unset A welcome email is sent to a new member on joining the group.
	ResourceBehaviorWelcomeEmailDisabled ResourceBehavior = "WelcomeEmailDisabled"
)
