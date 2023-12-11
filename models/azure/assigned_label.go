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

// Represents a sensitivity label assigned to a Microsoft 365 group. Sensitivity labels allow administrators to enforce
// specific group settings on a group by assigning a classification to the group (such as Confidential, Highly Confidential
// or General). Sensitivity labels are published by administrators in Microsoft 365 Security & Compliance Center as part
// of Microsoft Information Protection capabilities.
// For more detail see https://docs.microsoft.com/en-us/graph/api/resources/assignedlabel?view=graph-rest-1.0
type AssignedLabel struct {
	// The unique identifier of the label.
	LabelId string `json:"labelId,omitempty"`

	// The display name of the label. Read-only.
	DisplayName string `json:"displayName,omitempty"`
}
