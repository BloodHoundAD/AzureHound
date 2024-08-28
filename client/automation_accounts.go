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

func (s *azureClient) GetAzureAutomationAccounts(ctx context.Context, subscriptionId string) (azure.AutomationAccountList, error) {
	var (
		path     = fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Automation/automationAccounts", subscriptionId)
		params   = query.RMParams{ApiVersion: "2021-06-22"}
		response azure.AutomationAccountList
	)

	if res, err := s.resourceManager.Get(ctx, path, params, nil); err != nil {
		return response, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return response, err
	} else {
		return response, nil
	}
}

func (s *azureClient) ListAzureAutomationAccounts(ctx context.Context, subscriptionId string) <-chan azure.AutomationAccountResult {
	out := make(chan azure.AutomationAccountResult)

	go func() {
		defer panicrecovery.PanicRecovery()
		defer close(out)

		var (
			errResult = azure.AutomationAccountResult{
				SubscriptionId: subscriptionId,
			}
			nextLink string
		)

		if result, err := s.GetAzureAutomationAccounts(ctx, subscriptionId); err != nil {
			errResult.Error = err
			if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
				return
			}
		} else {
			for _, u := range result.Value {
				if ok := pipeline.Send(ctx.Done(), out, azure.AutomationAccountResult{SubscriptionId: subscriptionId, Ok: u}); !ok {
					return
				}
			}

			nextLink = result.NextLink
			for nextLink != "" {
				var list azure.AutomationAccountList
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
						if ok := pipeline.Send(ctx.Done(), out, azure.AutomationAccountResult{
							SubscriptionId: "/subscriptions/" + subscriptionId,
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
