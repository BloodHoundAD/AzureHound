// Copyright (C) 2022 The BloodHound Enterprise Team
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

import (
	"github.com/gofrs/uuid"
	"github.com/bloodhoundad/azurehound/enums"
)

// Represents a plan assigned to user and organization entities.
type AssignedPlan struct {
	// The date and time at which the plan was assigned using ISO 8601 format.
	AssignedDateTime string `json:"assignedDateTime"`

	// Condition of the capability assignment.
	CapabilityStatus enums.CapabiltyStatus `json:"capabilityStatus"`

	// The name of the service.
	Service string `json:"service"`

	// A GUID that identifies the service plan.
	ServicePlanId uuid.UUID `json:"servicePlanId"`
}
