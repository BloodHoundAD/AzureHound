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

package models

import "time"

type ClientTask struct {
	ADStructureCollection bool      `json:"ad_structure_collection"`
	ClientId              string    `json:"client_id"`
	CreatedAt             time.Time `json:"created_at"`
	DomainController      string    `json:"domain_controller"`
	EndTime               time.Time `json:"end_time"`
	EventId               int       `json:"event_id"`
	EventTitle            string    `json:"event_title"`
	ExectionTime          time.Time `json:"exection_time"`
	Id                    int       `json:"id"`
	LocalGroupCollection  bool      `json:"local_group_collection"`
	LogPath               string    `json:"log_path"`
	SessionCollection     bool      `json:"session_collection"`
	StartTime             time.Time `json:"start_time"`
	Status                int       `json:"status"`
	UpdatedAt             time.Time `json:"updated_at"`
}
