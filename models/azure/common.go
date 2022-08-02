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

import "encoding/json"

type Response struct {
	Context  string            `json:"@odata.context,omitempty"`
	Count    int               `json:"@odata.count,omitempty"`
	NextLink string            `json:"@odata.nextLink,omitempty"`
	Value    []json.RawMessage `json:"value"`
}

type ErrorResponse struct {
	Error ODataError `json:"error"`
}

type ErrorAdditionalInfo struct {
	Info map[string]string `json:"info,omitempty"`
	Type string            `json:"type,omitempty"`
}

type ODataError struct {
	AdditionalInfo []ErrorAdditionalInfo `json:"additionalInfo,omitempty"`
	Code           string                `json:"code"`
	Details        []ODataError          `json:"details,omitempty"`
	Message        string                `json:"message"`
	InnerError     *ODataError           `json:"innererror,omitempty"`
	Target         string                `json:"target,omitempty"`
}
