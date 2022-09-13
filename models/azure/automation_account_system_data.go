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

type AutomationAccountSystemData struct {
	CreatedAt          string                        `json:"createdAt,omitempty"`
	CreatedBy          string                        `json:"createdBy,omitempty"`
	CreatedByType      AutomationAccountIdentityType `json:"createdByType,omitempty"`
	LastModifiedAt     string                        `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     string                        `json:"lastModifiedBy,omitempty"`
	LastModifiedByType AutomationAccountIdentityType `json:"lastModifiedByType,omitempty"`
}

type AutomationAccountIdentityType struct {
	Application     string `json:"Application,omitempty"`
	Key             string `json:"Key,omitempty"`
	ManagedIdentity string `json:"ManagedIdentity,omitempty"`
	User            string `json:"User,omitempty"`
}
