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
	"strings"

	"github.com/bloodhoundad/azurehound/v2/client/query"
	"github.com/bloodhoundad/azurehound/v2/client/rest"
	"github.com/bloodhoundad/azurehound/v2/constants"
	"github.com/bloodhoundad/azurehound/v2/models/azure"
	"github.com/bloodhoundad/azurehound/v2/panicrecovery"
	"github.com/bloodhoundad/azurehound/v2/pipeline"
)

func (s *azureClient) GetAzureADRoleAssignment(ctx context.Context, objectId string, selectCols []string) (*azure.UnifiedRoleAssignment, error) {
	var (
		path     = fmt.Sprintf("/%s/roleManagement/directory/roleAssignments/%s", constants.GraphApiVersion, objectId)
		params   = query.Params{Select: selectCols}.AsMap()
		response azure.UnifiedRoleAssignment
	)
	if res, err := s.msgraph.Get(ctx, path, params, nil); err != nil {
		return nil, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return nil, err
	} else {
		return &response, nil
	}
}

func (s *azureClient) GetAzureADRoleAssignments(ctx context.Context, filter, search, orderBy, expand string, selectCols []string, top int32, count bool) (azure.UnifiedRoleAssignmentList, error) {
	var (
		path     = fmt.Sprintf("/%s/roleManagement/directory/roleAssignments", constants.GraphApiVersion)
		params   = query.Params{Filter: filter, Search: search, OrderBy: orderBy, Select: selectCols, Top: top, Count: count, Expand: expand}
		headers  map[string]string
		response azure.UnifiedRoleAssignmentList
	)

	count = count || search != "" || (filter != "" && orderBy != "") || strings.Contains(filter, "endsWith")
	if count {
		headers = make(map[string]string)
		headers["ConsistencyLevel"] = "eventual"
	}
	if res, err := s.msgraph.Get(ctx, path, params.AsMap(), headers); err != nil {
		return response, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return response, err
	} else {
		return response, nil
	}
}

func (s *azureClient) ListAzureADRoleAssignments(ctx context.Context, filter, search, orderBy, expand string, selectCols []string) <-chan azure.UnifiedRoleAssignmentResult {
	out := make(chan azure.UnifiedRoleAssignmentResult)

	go func() {
		defer panicrecovery.PanicRecovery()
		defer close(out)

		var (
			errResult = azure.UnifiedRoleAssignmentResult{}
			nextLink  string
		)

		if list, err := s.GetAzureADRoleAssignments(ctx, filter, search, orderBy, expand, selectCols, 999, false); err != nil {
			errResult.Error = err
			if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
				return
			}
		} else {
			for _, u := range list.Value {
				if ok := pipeline.Send(ctx.Done(), out, azure.UnifiedRoleAssignmentResult{Ok: u}); !ok {
					return
				}
			}

			nextLink = list.NextLink
			for nextLink != "" {
				var list azure.UnifiedRoleAssignmentList
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
				} else if res, err := s.msgraph.Send(req); err != nil {
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
						if ok := pipeline.Send(ctx.Done(), out, azure.UnifiedRoleAssignmentResult{Ok: u}); !ok {
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

func (s *azureClient) GetRoleAssignmentsForResource(ctx context.Context, resourceId string, filter string) (azure.RoleAssignmentList, error) {
	var (
		path     = fmt.Sprintf("%s/providers/Microsoft.Authorization/roleAssignments", resourceId)
		params   = query.Params{ApiVersion: "2015-07-01", Filter: filter}.AsMap()
		headers  map[string]string
		response azure.RoleAssignmentList
	)

	if res, err := s.resourceManager.Get(ctx, path, params, headers); err != nil {
		return response, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return response, err
	} else {
		return response, nil
	}

}

func (s *azureClient) ListRoleAssignmentsForResource(ctx context.Context, resourceId string, filter string) <-chan azure.RoleAssignmentResult {
	out := make(chan azure.RoleAssignmentResult)

	go func() {
		defer panicrecovery.PanicRecovery()
		defer close(out)

		var (
			errResult = azure.RoleAssignmentResult{ParentId: resourceId}
			nextLink  string
		)

		if result, err := s.GetRoleAssignmentsForResource(ctx, resourceId, filter); err != nil {
			errResult.Error = err
			if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
				return
			}
		} else {
			for _, u := range result.Value {
				if ok := pipeline.Send(ctx.Done(), out, azure.RoleAssignmentResult{
					ParentId: resourceId,
					Ok:       u,
				}); !ok {
					return
				}
			}

			nextLink = result.NextLink
			for nextLink != "" {
				var list azure.RoleAssignmentList
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
						if ok := pipeline.Send(ctx.Done(), out, azure.RoleAssignmentResult{
							ParentId: resourceId,
							Ok:       u,
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

func (s *azureClient) GetResourceRoleAssignments(ctx context.Context, subscriptionId string, filter string, expand string) (azure.RoleAssignmentList, error) {
	var (
		path     = fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleAssignments", subscriptionId)
		params   = query.Params{ApiVersion: "2015-07-01", Filter: filter, Expand: expand}.AsMap()
		headers  map[string]string
		response azure.RoleAssignmentList
	)

	if res, err := s.resourceManager.Get(ctx, path, params, headers); err != nil {
		return response, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return response, err
	} else {
		return response, nil
	}
}

func (s *azureClient) ListResourceRoleAssignments(ctx context.Context, subscriptionId string, filter string, expand string) <-chan azure.RoleAssignmentResult {
	out := make(chan azure.RoleAssignmentResult)

	go func() {
		defer panicrecovery.PanicRecovery()
		defer close(out)

		var (
			errResult = azure.RoleAssignmentResult{ParentId: subscriptionId}
			nextLink  string
		)

		if result, err := s.GetResourceRoleAssignments(ctx, subscriptionId, filter, expand); err != nil {
			errResult.Error = err
			if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
				return
			}
		} else {
			for _, u := range result.Value {
				if ok := pipeline.Send(ctx.Done(), out, azure.RoleAssignmentResult{
					ParentId: subscriptionId,
					Ok:       u,
				}); !ok {
					return
				}
			}

			nextLink = result.NextLink
			for nextLink != "" {
				var list azure.RoleAssignmentList
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
						if ok := pipeline.Send(ctx.Done(), out, azure.RoleAssignmentResult{
							ParentId: subscriptionId,
							Ok:       u,
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
