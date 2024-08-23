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
	"github.com/bloodhoundad/azurehound/v2/constants"
	"github.com/bloodhoundad/azurehound/v2/models/azure"
	"github.com/bloodhoundad/azurehound/v2/panicrecovery"
	"github.com/bloodhoundad/azurehound/v2/pipeline"
)

func (s *azureClient) GetAzureADRoleAssignments(ctx context.Context, params query.GraphParams) (azure.UnifiedRoleAssignmentList, error) {
	var (
		path     = fmt.Sprintf("/%s/roleManagement/directory/roleAssignments", constants.GraphApiVersion)
		response azure.UnifiedRoleAssignmentList
	)

	if params.Top == 0 {
		params.Top = 999
	}

	if res, err := s.msgraph.Get(ctx, path, params, nil); err != nil {
		return response, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return response, err
	} else {
		return response, nil
	}
}

func (s *azureClient) ListAzureADRoleAssignments(ctx context.Context, params query.GraphParams) <-chan azure.UnifiedRoleAssignmentResult {
	out := make(chan azure.UnifiedRoleAssignmentResult)

	go func() {
		defer panicrecovery.PanicRecovery()
		defer close(out)

		var (
			errResult = azure.UnifiedRoleAssignmentResult{}
			nextLink  string
		)

		if list, err := s.GetAzureADRoleAssignments(ctx, params); err != nil {
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

func (s *azureClient) GetRoleAssignmentsForResource(ctx context.Context, resourceId string, filter, tenantId string) (azure.RoleAssignmentList, error) {
	var (
		path     = fmt.Sprintf("%s/providers/Microsoft.Authorization/roleAssignments", resourceId)
		params   = query.RMParams{ApiVersion: "2015-07-01", Filter: filter, TenantId: tenantId}
		response azure.RoleAssignmentList
	)

	if res, err := s.resourceManager.Get(ctx, path, params, nil); err != nil {
		return response, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return response, err
	} else {
		return response, nil
	}

}

func (s *azureClient) ListRoleAssignmentsForResource(ctx context.Context, resourceId string, filter, tenantId string) <-chan azure.RoleAssignmentResult {
	out := make(chan azure.RoleAssignmentResult)

	go func() {
		defer panicrecovery.PanicRecovery()
		defer close(out)

		var (
			errResult = azure.RoleAssignmentResult{ParentId: resourceId}
			nextLink  string
		)

		if result, err := s.GetRoleAssignmentsForResource(ctx, resourceId, filter, tenantId); err != nil {
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
