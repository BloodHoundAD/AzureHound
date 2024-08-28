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

package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bloodhoundad/azurehound/v2/client/query"
	"github.com/bloodhoundad/azurehound/v2/constants"
	"github.com/bloodhoundad/azurehound/v2/models/azure"
)

// ListAzureADApps https://learn.microsoft.com/en-us/graph/api/application-list?view=graph-rest-beta
func (s *azureClient) ListAzureADApps(ctx context.Context, params query.GraphParams) <-chan azureResult[azure.Application] {
	var (
		out = make(chan azureResult[azure.Application])
		path      = fmt.Sprintf("/%s/applications", constants.GraphApiVersion)

	)

	if params.Top == 0 {
		params.Top = 99
	}

	go getAzureObjectList[azure.Application](s.msgraph, ctx, path, params, out)

	return out
}

// ListAzureADAppOwners https://learn.microsoft.com/en-us/graph/api/application-list-owners?view=graph-rest-beta
func (s *azureClient) ListAzureADAppOwners(ctx context.Context, objectId string, params query.GraphParams) <-chan azureResult[json.RawMessage] {

	var (
		out  = make(chan azureResult[json.RawMessage])
		path = fmt.Sprintf("/%s/applications/%s/owners", constants.GraphApiBetaVersion, objectId)
	)

	if params.Top == 0 {
		params.Top = 99
	}

	go getAzureObjectList[json.RawMessage](s.msgraph, ctx, path, params, out)

	return out
}
