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

type PrivilegedAccessGroupEligibilityScheduleInstance struct {
	Entity

	// When this instance starts.
	StartDateTime string `json:"startDateTime"`

	// When the schedule instance ends.
	EndDateTime string `json:"endDateTime"`

	// The identifier of the principal whose membership or ownership eligibility to the group is managed through PIM for groups.
	PrincipalId string `json:"principalId"`

	// The identifier of the membership or ownership eligibility relationship to the group.
	// The possible values are: owner, member.
	AccessId string `json:"accessId"`

	// The identifier of the group representing the scope of the membership or ownership eligibility through PIM for groups.
	GroupId string `json:"groupId"`

	// Indicates whether the assignment is derived from a group assignment. It can further imply whether the calling
	// principal can manage the assignment schedule.
	// The possible values are: direct, group, unknownFutureValue.
	MemberType string `json:"memberType"`

	// The identifier of the privilegedAccessGroupEligibilitySchedule from which this instance was created.
	EligibilityScheduleId string `json:"eligibilityScheduleId"`
}

type PrivilegedAccessGroupEligibilityScheduleInstanceList struct {
	Count    int                                                `json:"@odata.count,omitempty"`    // The total count of all results
	NextLink string                                             `json:"@odata.nextLink,omitempty"` // The URL to use for getting the next set of values.
	Value    []PrivilegedAccessGroupEligibilityScheduleInstance `json:"value"`                     // A list of role assignments.
}

type PrivilegedAccessGroupEligibilityScheduleInstanceResult struct {
	Error error
	Ok    PrivilegedAccessGroupEligibilityScheduleInstance
}
