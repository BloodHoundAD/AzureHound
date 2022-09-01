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
	"net/url"

	"github.com/bloodhoundad/azurehound/client/query"
	"github.com/bloodhoundad/azurehound/client/rest"
	"github.com/bloodhoundad/azurehound/models/azure"
)

func (s *azureClient) GetAzureFunctionApp(ctx context.Context, subscriptionId, groupName, functionAppName, expand string) (*azure.FunctionApp, error) {
	var (
		path     = fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s", subscriptionId, groupName, functionAppName)
		params   = query.Params{ApiVersion: "2022-03-01", Expand: expand}.AsMap()
		headers  map[string]string
		response azure.FunctionApp
	)
	if res, err := s.resourceManager.Get(ctx, path, params, headers); err != nil {
		return nil, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return nil, err
	} else {
		return &response, nil
	}
}

func (s *azureClient) GetAzureFunctionApps(ctx context.Context, subscriptionId string, statusOnly bool) (azure.FunctionAppList, error) {
	var (
		path     = fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Web/sites", subscriptionId)
		params   = query.Params{ApiVersion: "2022-03-01", StatusOnly: statusOnly}.AsMap()
		headers  map[string]string
		response azure.FunctionAppList
	)

	if res, err := s.resourceManager.Get(ctx, path, params, headers); err != nil {
		return response, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return response, err
	} else {
		return response, nil
	}
}

func (s *azureClient) ListAzureFunctionApps(ctx context.Context, subscriptionId string, statusOnly bool) <-chan azure.FunctionAppResult {
	out := make(chan azure.FunctionAppResult)

	go func() {
		defer close(out)

		var (
			errResult = azure.FunctionAppResult{
				SubscriptionId: subscriptionId,
			}
			nextLink string
		)

		if result, err := s.GetAzureFunctionApps(ctx, subscriptionId, statusOnly); err != nil {
			errResult.Error = err
			out <- errResult
		} else {
			for _, u := range result.Value {
				out <- azure.FunctionAppResult{SubscriptionId: subscriptionId, Ok: u}
			}

			nextLink = result.NextLink
			for nextLink != "" {
				var list azure.FunctionAppList
				if url, err := url.Parse(nextLink); err != nil {
					errResult.Error = err
					out <- errResult
					nextLink = ""
				} else if req, err := rest.NewRequest(ctx, "GET", url, nil, nil, nil); err != nil {
					errResult.Error = err
					out <- errResult
					nextLink = ""
				} else if res, err := s.resourceManager.Send(req); err != nil {
					errResult.Error = err
					out <- errResult
					nextLink = ""
				} else if err := rest.Decode(res.Body, &list); err != nil {
					errResult.Error = err
					out <- errResult
					nextLink = ""
				} else {
					for _, u := range list.Value {
						out <- azure.FunctionAppResult{
							SubscriptionId: "/subscriptions/" + subscriptionId,
							Ok:             u,
						}
					}
					nextLink = list.NextLink
				}
			}
		}
	}()
	return out
}
