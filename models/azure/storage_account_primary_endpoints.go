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

type Endpoints struct {
	Blob               string                           `json:"blob,omitempty"`
	DFS                string                           `json:"dfs,omitempty"`
	File               string                           `json:"file,omitempty"`
	InternetEndpoints  StorageAccountInternetEndpoints  `json:"internetEndpoints,omitempty"`
	MicrosoftEndpoints StorageAccountMicrosoftEndpoints `json:"microsoftEndpoints,omitempty"`
	Queue              string                           `json:"queue,omitempty"`
	Table              string                           `json:"table,omitempty"`
	Web                string                           `json:"web,omitempty"`
}

type StorageAccountInternetEndpoints struct {
	Blob string `json:"blob,omitempty"`
	DFS  string `json:"dfs,omitempty"`
	File string `json:"file,omitempty"`
	Web  string `json:"web,omitempty"`
}

type StorageAccountMicrosoftEndpoints struct {
	Blob  string `json:"blob,omitempty"`
	DFS   string `json:"dfs,omitempty"`
	File  string `json:"file,omitempty"`
	Queue string `json:"queue,omitempty"`
	Table string `json:"table,omitempty"`
	Web   string `json:"web,omitempty"`
}
