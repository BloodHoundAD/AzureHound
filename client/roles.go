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
	"github.com/bloodhoundad/azurehound/constants"
	"github.com/bloodhoundad/azurehound/models/azure"
)

func (s *azureClient) GetAzureADRole(ctx context.Context, roleId string, selectCols []string) (*azure.Role, error) {
	var (
		path     = fmt.Sprintf("/%s/roleManagement/directory/roleDefinitions/%s", constants.GraphApiVersion, roleId)
		params   = query.Params{Select: selectCols}.AsMap()
		response azure.RoleList
	)
	if res, err := s.msgraph.Get(ctx, path, params, nil); err != nil {
		return nil, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return nil, err
	} else {
		return &response.Value[0], nil
	}
}

func (s *azureClient) GetAzureADRoles(ctx context.Context, filter, expand string) (azure.RoleList, error) {
	var (
		path     = fmt.Sprintf("/%s/roleManagement/directory/roleDefinitions", constants.GraphApiVersion)
		params   = query.Params{Filter: filter, Expand: expand}
		headers  map[string]string
		response azure.RoleList
	)

	if res, err := s.msgraph.Get(ctx, path, params.AsMap(), headers); err != nil {
		return response, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return response, err
	} else {
		return response, nil
	}
}

func (s *azureClient) ListAzureADRoles(ctx context.Context, filter, expand string) <-chan azure.RoleResult {
	out := make(chan azure.RoleResult)

	go func() {
		defer close(out)

		var (
			errResult = azure.RoleResult{}
			nextLink  string
		)

		if users, err := s.GetAzureADRoles(ctx, filter, expand); err != nil {
			errResult.Error = err
			out <- errResult
		} else {
			for _, u := range users.Value {
				out <- azure.RoleResult{Ok: u}
			}

			nextLink = users.NextLink
			for nextLink != "" {
				var users azure.RoleList
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
				} else if err := rest.Decode(res.Body, &users); err != nil {
					errResult.Error = err
					out <- errResult
					nextLink = ""
				} else {
					for _, u := range users.Value {
						out <- azure.RoleResult{Ok: u}
					}
					nextLink = users.NextLink
				}
			}
		}
	}()
	return out
}
