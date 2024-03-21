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
	"github.com/bloodhoundad/azurehound/v2/enums"
	"github.com/bloodhoundad/azurehound/v2/models/azure"
	"github.com/bloodhoundad/azurehound/v2/panicrecovery"
	"github.com/bloodhoundad/azurehound/v2/pipeline"
)

func (s *azureClient) GetAzureADGroup(ctx context.Context, objectId string, selectCols []string) (*azure.Group, error) {
	var (
		path     = fmt.Sprintf("/%s/groups/%s", constants.GraphApiVersion, objectId)
		params   = query.Params{Select: selectCols}.AsMap()
		response azure.GroupList
	)
	if res, err := s.msgraph.Get(ctx, path, params, nil); err != nil {
		return nil, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return nil, err
	} else {
		return &response.Value[0], nil
	}
}

func (s *azureClient) GetAzureADGroupOwners(ctx context.Context, objectId string, filter string, search string, orderBy string, selectCols []string, top int32, count bool) (azure.DirectoryObjectList, error) {
	var (
		path     = fmt.Sprintf("/%s/groups/%s/owners", constants.GraphApiBetaVersion, objectId)
		params   = query.Params{Filter: filter, Search: search, OrderBy: orderBy, Select: selectCols, Top: top, Count: count}.AsMap()
		response azure.DirectoryObjectList
	)
	if res, err := s.msgraph.Get(ctx, path, params, nil); err != nil {
		return response, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return response, err
	} else {
		return response, nil
	}
}

func (s *azureClient) GetAzureADGroupMembers(ctx context.Context, objectId string, filter string, search string, count bool) (azure.MemberObjectList, error) {
	var (
		path     = fmt.Sprintf("/%s/groups/%s/members", constants.GraphApiBetaVersion, objectId)
		params   = query.Params{Filter: filter, Search: search, Count: count}.AsMap()
		response azure.MemberObjectList
	)
	if res, err := s.msgraph.Get(ctx, path, params, nil); err != nil {
		return response, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return response, err
	} else {
		return response, nil
	}
}

func (s *azureClient) GetAzureADGroups(ctx context.Context, filter, search, orderBy, expand string, selectCols []string, top int32, count bool) (azure.GroupList, error) {
	var (
		path     = fmt.Sprintf("/%s/groups", constants.GraphApiVersion)
		params   = query.Params{Filter: filter, Search: search, OrderBy: orderBy, Select: selectCols, Top: top, Count: count, Expand: expand}
		headers  map[string]string
		response azure.GroupList
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

func (s *azureClient) ListAzureADGroups(ctx context.Context, filter, search, orderBy, expand string, selectCols []string) <-chan azure.GroupResult {
	out := make(chan azure.GroupResult)

	go func() {
		defer panicrecovery.PanicRecovery()
		defer close(out)

		var (
			errResult = azure.GroupResult{}
			nextLink  string
		)

		if list, err := s.GetAzureADGroups(ctx, filter, search, orderBy, expand, selectCols, 999, false); err != nil {
			errResult.Error = err
			if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
				return
			}
		} else {
			for _, u := range list.Value {
				if ok := pipeline.Send(ctx.Done(), out, azure.GroupResult{Ok: u}); !ok {
					return
				}
			}

			nextLink = list.NextLink
			for nextLink != "" {
				var list azure.GroupList
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
						if ok := pipeline.Send(ctx.Done(), out, azure.GroupResult{Ok: u}); !ok {
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

func (s *azureClient) ListAzureADGroupOwners(ctx context.Context, objectId string, filter, search, orderBy string, selectCols []string) <-chan azure.GroupOwnerResult {
	out := make(chan azure.GroupOwnerResult)

	go func() {
		defer panicrecovery.PanicRecovery()
		defer close(out)

		var (
			errResult = azure.GroupOwnerResult{}
			nextLink  string
		)

		if list, err := s.GetAzureADGroupOwners(ctx, objectId, filter, search, orderBy, selectCols, 999, false); err != nil {
			errResult.Error = err
			if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
				return
			}
		} else {
			for _, u := range list.Value {
				if ok := pipeline.Send(ctx.Done(), out, azure.GroupOwnerResult{
					GroupId: objectId,
					Ok:      u,
				}); !ok {
					return
				}
			}

			nextLink = list.NextLink
			for nextLink != "" {
				var list azure.DirectoryObjectList
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
						if ok := pipeline.Send(ctx.Done(), out, azure.GroupOwnerResult{
							GroupId: objectId,
							Ok:      u,
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

func (s *azureClient) ListAzureADGroupMembers(ctx context.Context, objectId string, filter, search, orderBy string, selectCols []string) <-chan azure.MemberObjectResult {
	out := make(chan azure.MemberObjectResult)

	go func() {
		defer panicrecovery.PanicRecovery()
		defer close(out)

		var (
			errResult = azure.MemberObjectResult{
				ParentId:   objectId,
				ParentType: string(enums.EntityGroup),
			}
			nextLink string
		)

		if list, err := s.GetAzureADGroupMembers(ctx, objectId, filter, search, false); err != nil {
			errResult.Error = err
			if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
				return
			}
		} else {
			for _, u := range list.Value {
				if ok := pipeline.Send(ctx.Done(), out, azure.MemberObjectResult{
					ParentId:   objectId,
					ParentType: string(enums.EntityGroup),
					Ok:         u,
				}); !ok {
					return
				}
			}

			nextLink = list.NextLink
			for nextLink != "" {
				var list azure.MemberObjectList
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
						if ok := pipeline.Send(ctx.Done(), out, azure.MemberObjectResult{
							ParentId:   objectId,
							ParentType: string(enums.EntityGroup),
							Ok:         u,
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
