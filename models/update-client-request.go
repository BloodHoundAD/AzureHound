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

type UpdateClientRequest struct {
	Address  string `json:"address"`
	Username string `json:"username"`
	Hostname string `json:"hostname"`
	Version  string `json:"version"`
	UserSid  string `json:"usersid"`
}

type UpdateClientResponse struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	IPAddress    string    `json:"ip_address"`
	Hostname     string    `json:"hostname"`
	CurrentJobID int       `json:"current_job_id"`
	CurrentJob   ClientJob `json:"current_job"`
}
