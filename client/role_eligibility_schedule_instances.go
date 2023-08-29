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
)

func (s *azureClient) GetAzureADRoleEligibilityScheduleInstance(ctx context.Context, objectId string, selectCols []string) (*azure.UnifiedRoleEligibilityScheduleInstance, error) {
	var (
		path     = fmt.Sprintf("/%s/roleManagement/directory/RoleEligibilityScheduleInstances/%s", constants.GraphApiVersion, objectId)
		params   = query.Params{Select: selectCols}.AsMap()
		response azure.UnifiedRoleEligibilityScheduleInstance
	)
	if res, err := s.msgraph.Get(ctx, path, params, nil); err != nil {
		return nil, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return nil, err
	} else {
		return &response, nil
	}
}

func (s *azureClient) GetAzureADRoleEligibilityScheduleInstances(ctx context.Context, filter, search, orderBy, expand string, selectCols []string, top int32, count bool) (azure.UnifiedRoleEligibilityScheduleInstanceList, error) {
	var (
		path     = fmt.Sprintf("/%s/roleManagement/directory/RoleEligibilityScheduleInstances", constants.GraphApiVersion)
		params   = query.Params{Filter: filter, Search: search, OrderBy: orderBy, Select: selectCols, Top: top, Count: count, Expand: expand}
		headers  map[string]string
		response azure.UnifiedRoleEligibilityScheduleInstanceList
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

func (s *azureClient) ListAzureADRoleEligibilityScheduleInstances(ctx context.Context, filter, search, orderBy, expand string, selectCols []string) <-chan azure.UnifiedRoleEligibilityScheduleInstanceResult {
	out := make(chan azure.UnifiedRoleEligibilityScheduleInstanceResult)

	go func() {
		defer close(out)

		var (
			errResult = azure.UnifiedRoleEligibilityScheduleInstanceResult{}
			nextLink  string
		)

		if list, err := s.GetAzureADRoleEligibilityScheduleInstances(ctx, filter, search, orderBy, expand, selectCols, 999, false); err != nil {
			errResult.Error = err
			out <- errResult
		} else {
			for _, u := range list.Value {
				out <- azure.UnifiedRoleEligibilityScheduleInstanceResult{Ok: u}
			}

			nextLink = list.NextLink
			for nextLink != "" {
				var list azure.UnifiedRoleEligibilityScheduleInstanceList
				if url, err := url.Parse(nextLink); err != nil {
					errResult.Error = err
					out <- errResult
					nextLink = ""
				} else if req, err := rest.NewRequest(ctx, "GET", url, nil, nil, nil); err != nil {
					errResult.Error = err
					out <- errResult
					nextLink = ""
				} else if res, err := s.msgraph.Send(req); err != nil {
					errResult.Error = err
					out <- errResult
					nextLink = ""
				} else if err := rest.Decode(res.Body, &list); err != nil {
					errResult.Error = err
					out <- errResult
					nextLink = ""
				} else {
					for _, u := range list.Value {
						out <- azure.UnifiedRoleEligibilityScheduleInstanceResult{Ok: u}
					}
					nextLink = list.NextLink
				}
			}
		}
	}()
	return out
}
