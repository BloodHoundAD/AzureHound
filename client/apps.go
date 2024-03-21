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

func (s *azureClient) GetAzureADApp(ctx context.Context, objectId string, selectCols []string) (*azure.Application, error) {
	var (
		path     = fmt.Sprintf("/%s/applications/%s", constants.GraphApiVersion, objectId)
		params   = query.Params{Select: selectCols}.AsMap()
		response azure.ApplicationList
	)
	if res, err := s.msgraph.Get(ctx, path, params, nil); err != nil {
		return nil, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return nil, err
	} else {
		return &response.Value[0], nil
	}
}

func (s *azureClient) GetAzureADAppOwners(ctx context.Context, objectId string, filter string, search string, orderBy string, selectCols []string, top int32, count bool) (azure.DirectoryObjectList, error) {
	var (
		path     = fmt.Sprintf("/%s/applications/%s/owners", constants.GraphApiBetaVersion, objectId)
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

func (s *azureClient) GetAzureADAppMemberObjects(ctx context.Context, objectId string, securityEnabledOnly bool) (azure.MemberObjectList, error) {
	var (
		path     = fmt.Sprintf("/%s/directoryObjects/%s/getMemberObjects", constants.GraphApiVersion, objectId)
		response azure.MemberObjectList
		body     = map[string]bool{
			"securityEnabledOnly": securityEnabledOnly,
		}
	)
	if res, err := s.msgraph.Post(ctx, path, body, nil, nil); err != nil {
		return response, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return response, err
	} else {
		return response, nil
	}

}

func (s *azureClient) GetAzureADApps(ctx context.Context, filter, search, orderBy, expand string, selectCols []string, top int32, count bool) (azure.ApplicationList, error) {
	var (
		path     = fmt.Sprintf("/%s/applications", constants.GraphApiVersion)
		params   = query.Params{Filter: filter, Search: search, OrderBy: orderBy, Select: selectCols, Top: top, Count: count, Expand: expand}
		headers  map[string]string
		response azure.ApplicationList
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

func (s *azureClient) ListAzureADApps(ctx context.Context, filter, search, orderBy, expand string, selectCols []string) <-chan azure.ApplicationResult {
	out := make(chan azure.ApplicationResult)

	go func() {
		defer panicrecovery.PanicRecovery()
		defer close(out)

		var (
			errResult = azure.ApplicationResult{}
			nextLink  string
		)

		if list, err := s.GetAzureADApps(ctx, filter, search, orderBy, expand, selectCols, 999, false); err != nil {
			errResult.Error = err
			if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
				return
			}
		} else {
			for _, u := range list.Value {
				if ok := pipeline.Send(ctx.Done(), out, azure.ApplicationResult{Ok: u}); !ok {
					return
				}
			}

			nextLink = list.NextLink
			for nextLink != "" {
				var list azure.ApplicationList
				if url, err := url.Parse(nextLink); err != nil {
					errResult.Error = err
					if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
						return
					}
					return
				} else if req, err := rest.NewRequest(ctx, "GET", url, nil, nil, nil); err != nil {
					errResult.Error = err
					if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
						return
					}
					return
				} else if res, err := s.msgraph.Send(req); err != nil {
					errResult.Error = err
					if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
						return
					}
					return
				} else if err := rest.Decode(res.Body, &list); err != nil {
					errResult.Error = err
					if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
						return
					}
					return
				} else {
					for _, u := range list.Value {
						if ok := pipeline.Send(ctx.Done(), out, azure.ApplicationResult{Ok: u}); !ok {
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

func (s *azureClient) ListAzureADAppOwners(ctx context.Context, objectId string, filter, search, orderBy string, selectCols []string) <-chan azure.AppOwnerResult {
	out := make(chan azure.AppOwnerResult)

	go func() {
		defer panicrecovery.PanicRecovery()
		defer close(out)

		var (
			errResult = azure.AppOwnerResult{}
			nextLink  string
		)

		if list, err := s.GetAzureADAppOwners(ctx, objectId, filter, search, orderBy, selectCols, 999, false); err != nil {
			errResult.Error = err
			if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
				return
			}
		} else {
			for _, u := range list.Value {
				if ok := pipeline.Send(ctx.Done(), out, azure.AppOwnerResult{
					AppId: objectId,
					Ok:    u,
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
					return
				} else if req, err := rest.NewRequest(ctx, "GET", url, nil, nil, nil); err != nil {
					errResult.Error = err
					if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
						return
					}
					return
				} else if res, err := s.msgraph.Send(req); err != nil {
					errResult.Error = err
					if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
						return
					}
					return
				} else if err := rest.Decode(res.Body, &list); err != nil {
					errResult.Error = err
					if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
						return
					}
					return
				} else {
					for _, u := range list.Value {
						if ok := pipeline.Send(ctx.Done(), out, azure.AppOwnerResult{
							AppId: objectId,
							Ok:    u,
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

func (s *azureClient) ListAzureADAppMemberObjects(ctx context.Context, objectId string, securityEnabledOnly bool) <-chan azure.MemberObjectResult {
	out := make(chan azure.MemberObjectResult)

	go func() {
		defer panicrecovery.PanicRecovery()
		defer close(out)

		var (
			errResult = azure.MemberObjectResult{
				ParentId:   objectId,
				ParentType: string(enums.EntityApplication),
			}
			nextLink string
		)
		if list, err := s.GetAzureADAppMemberObjects(ctx, objectId, securityEnabledOnly); err != nil {
			errResult.Error = err
			if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
				return
			}
		} else {
			for _, u := range list.Value {
				if ok := pipeline.Send(ctx.Done(), out, azure.MemberObjectResult{
					ParentId:   objectId,
					ParentType: string(enums.EntityApplication),
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
							ParentType: string(enums.EntityApplication),
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
