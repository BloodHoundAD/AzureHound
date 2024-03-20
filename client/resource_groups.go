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

	"github.com/bloodhoundad/azurehound/v2/client/query"
	"github.com/bloodhoundad/azurehound/v2/client/rest"
	"github.com/bloodhoundad/azurehound/v2/models/azure"
	"github.com/bloodhoundad/azurehound/v2/panicrecovery"
	"github.com/bloodhoundad/azurehound/v2/pipeline"
)

func (s *azureClient) GetAzureResourceGroup(ctx context.Context, subscriptionId, groupName string) (*azure.ResourceGroup, error) {
	var (
		path     = fmt.Sprintf("/subscriptions/%s/resourcegroups/%s", subscriptionId, groupName)
		params   = query.Params{ApiVersion: "2021-04-01"}.AsMap()
		headers  map[string]string
		response azure.ResourceGroup
	)
	if res, err := s.resourceManager.Get(ctx, path, params, headers); err != nil {
		return nil, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return nil, err
	} else {
		return &response, nil
	}
}

func (s *azureClient) GetAzureResourceGroups(ctx context.Context, subscriptionId string, filter string, top int32) (azure.ResourceGroupList, error) {
	var (
		path     = fmt.Sprintf("/subscriptions/%s/resourcegroups", subscriptionId)
		params   = query.Params{ApiVersion: "2021-04-01", Filter: filter, Top: top}.AsMap()
		headers  map[string]string
		response azure.ResourceGroupList
	)

	if res, err := s.resourceManager.Get(ctx, path, params, headers); err != nil {
		return response, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return response, err
	} else {
		return response, nil
	}
}

func (s *azureClient) ListAzureResourceGroups(ctx context.Context, subscriptionId, filter string) <-chan azure.ResourceGroupResult {
	out := make(chan azure.ResourceGroupResult)

	go func() {
		defer panicrecovery.PanicRecovery()
		defer close(out)

		var (
			objectId  = fmt.Sprintf("/subscriptions/%s", subscriptionId)
			errResult = azure.ResourceGroupResult{SubscriptionId: objectId}
			nextLink  string
		)

		if result, err := s.GetAzureResourceGroups(ctx, subscriptionId, filter, 1000); err != nil {
			errResult.Error = err
			if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
				return
			}
		} else {
			for _, u := range result.Value {
				if ok := pipeline.Send(ctx.Done(), out, azure.ResourceGroupResult{
					SubscriptionId: objectId,
					Ok:             u,
				}); !ok {
					return
				}
			}

			nextLink = result.NextLink
			for nextLink != "" {
				var list azure.ResourceGroupList
				if url, err := url.Parse(nextLink); err != nil {
					errResult.Error = err
					if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
						return
					}
					nextLink = ""
				} else if req, err := rest.NewRequest(ctx, "GET", url, nil, nil, nil); err != nil {
					errResult.Error = err
					if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
						return
					}
					nextLink = ""
				} else if res, err := s.resourceManager.Send(req); err != nil {
					errResult.Error = err
					if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
						return
					}
					nextLink = ""
				} else if err := rest.Decode(res.Body, &list); err != nil {
					errResult.Error = err
					if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
						return
					}
					nextLink = ""
				} else {
					for _, u := range list.Value {
						if ok := pipeline.Send(ctx.Done(), out, azure.ResourceGroupResult{
							SubscriptionId: objectId,
							Ok:             u,
						}); !ok {
							return
						}
					}
					nextLink = list.NextLink
				}
			}
		}
	}()
	return out
}
