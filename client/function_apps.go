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
	"fmt"

	"github.com/bloodhoundad/azurehound/v2/client/query"
	"github.com/bloodhoundad/azurehound/v2/models/azure"
)

// ListAzureFunctionApps
func (s *azureClient) ListAzureFunctionApps(ctx context.Context, subscriptionId string) <-chan AzureResult[azure.FunctionApp] {
	var (
		out    = make(chan AzureResult[azure.FunctionApp])
		path   = fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Web/sites", subscriptionId)
		params = query.RMParams{ApiVersion: "2022-03-01"}
	)

	go getAzureObjectList[azure.FunctionApp](s.resourceManager, ctx, path, params, out)

	return out
}
