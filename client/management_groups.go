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

func (s *azureClient) GetAzureManagementGroup(ctx context.Context, groupId, filter, expand string, recurse bool) (*azure.ManagementGroup, error) {
	var (
		path     = fmt.Sprintf("/providers/Microsoft.Management/managementGroups/%s", groupId)
		params   = query.Params{ApiVersion: "2020-05-01", Filter: filter, Expand: expand, Recurse: recurse}.AsMap()
		headers  map[string]string
		response azure.ManagementGroup
	)
	if res, err := s.resourceManager.Get(ctx, path, params, headers); err != nil {
		return nil, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return nil, err
	} else {
		return &response, nil
	}
}

func (s *azureClient) GetAzureManagementGroups(ctx context.Context) (azure.ManagementGroupList, error) {
	var (
		path     = "/providers/Microsoft.Management/managementGroups"
		params   = query.Params{ApiVersion: "2020-05-01"}.AsMap()
		headers  map[string]string
		response azure.ManagementGroupList
	)

	if res, err := s.resourceManager.Get(ctx, path, params, headers); err != nil {
		return response, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return response, err
	} else {
		return response, nil
	}
}

func (s *azureClient) GetAzureManagementGroupDescendants(ctx context.Context, groupId string, top int32) (azure.DescendantInfoList, error) {
	var (
		path     = fmt.Sprintf("/providers/Microsoft.Management/managementGroups/%s/descendants", groupId)
		params   = query.Params{ApiVersion: "2020-05-01", Top: top}.AsMap()
		headers  map[string]string
		response azure.DescendantInfoList
	)

	if res, err := s.resourceManager.Get(ctx, path, params, headers); err != nil {
		return response, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return response, err
	} else {
		return response, nil
	}
}

func (s *azureClient) ListAzureManagementGroups(ctx context.Context) <-chan azure.ManagementGroupResult {
	out := make(chan azure.ManagementGroupResult)

	go func() {
		defer panicrecovery.PanicRecovery()
		defer close(out)

		var (
			errResult = azure.ManagementGroupResult{}
			nextLink  string
		)

		if result, err := s.GetAzureManagementGroups(ctx); err != nil {
			errResult.Error = err
			if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
				return
			}
		} else {
			for _, u := range result.Value {
				if ok := pipeline.Send(ctx.Done(), out, azure.ManagementGroupResult{Ok: u}); !ok {
					return
				}
			}

			nextLink = result.NextLink
			for nextLink != "" {
				var list azure.ManagementGroupList
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
						if ok := pipeline.Send(ctx.Done(), out, azure.ManagementGroupResult{Ok: u}); !ok {
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

func (s *azureClient) ListAzureManagementGroupDescendants(ctx context.Context, groupId string) <-chan azure.DescendantInfoResult {
	out := make(chan azure.DescendantInfoResult)

	go func() {
		defer panicrecovery.PanicRecovery()
		defer close(out)

		var (
			errResult = azure.DescendantInfoResult{}
			nextLink  string
		)

		if result, err := s.GetAzureManagementGroupDescendants(ctx, groupId, 3000); err != nil {
			errResult.Error = err
			if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
				return
			}
		} else {
			for _, u := range result.Value {
				if ok := pipeline.Send(ctx.Done(), out, azure.DescendantInfoResult{Ok: u}); !ok {
					return
				}
			}

			nextLink = result.NextLink
			for nextLink != "" {
				var list azure.DescendantInfoList
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
						if ok := pipeline.Send(ctx.Done(), out, azure.DescendantInfoResult{Ok: u}); !ok {
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
