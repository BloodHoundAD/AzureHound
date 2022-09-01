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

import "strings"

type StorageContainer struct {
	Entity

	Location   string                     `json:"location,omitempty"`
	Name       string                     `json:"name,omitempty"`
	Etag       string                     `json:"etag,omitempty"`
	Properties StorageContainerProperties `json:"properties,omitempty"`
}

func (s StorageContainer) ResourceGroupName() string {
	parts := strings.Split(s.Id, "/")
	if len(parts) > 4 {
		return parts[4]
	} else {
		return ""
	}
}

func (s StorageContainer) ResourceGroupId() string {
	parts := strings.Split(s.Id, "/")
	if len(parts) > 5 {
		return strings.Join(parts[:5], "/")
	} else {
		return ""
	}
}

func (s StorageContainer) StorageAccountName() string {
	parts := strings.Split(s.Id, "/")
	if len(parts) > 8 {
		return parts[8]
	} else {
		return ""
	}
}

func (s StorageContainer) StorageAccountId() string {
	parts := strings.Split(s.Id, "/")
	if len(parts) > 9 {
		return strings.Join(parts[:9], "/")
	} else {
		return ""
	}
}

type StorageContainerList struct {
	NextLink string             `json:"nextLink,omitempty"` // The URL to use for getting the next set of values.
	Value    []StorageContainer `json:"value"`              // A list of storage containers.
}

type StorageContainerResult struct {
	SubscriptionId string
	Error          error
	Ok             StorageContainer
}
