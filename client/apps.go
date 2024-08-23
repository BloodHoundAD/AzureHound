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

func (s *azureClient) GetAzureADAppOwners(ctx context.Context, objectId string, params query.GraphParams) (azure.DirectoryObjectList, error) {
	var (
		path     = fmt.Sprintf("/%s/applications/%s/owners", constants.GraphApiBetaVersion, objectId)
		response azure.DirectoryObjectList
	)

	if params.Top == 0 {
		params.Top = 99
	}

	if res, err := s.msgraph.Get(ctx, path, params, nil); err != nil {
		return response, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return response, err
	} else {
		return response, nil
	}
}

func (s *azureClient) GetAzureADApps(ctx context.Context, params query.GraphParams) (azure.ApplicationList, error) {
	var (
		path     = fmt.Sprintf("/%s/applications", constants.GraphApiVersion)
		response azure.ApplicationList
	)

	if params.Top == 0 {
		params.Top = 99
	}

	if res, err := s.msgraph.Get(ctx, path, params, nil); err != nil {
		return response, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return response, err
	} else {
		return response, nil
	}
}

func (s *azureClient) ListAzureADApps(ctx context.Context, params query.GraphParams) <-chan azure.ApplicationResult {
	out := make(chan azure.ApplicationResult)

	go func() {
		defer panicrecovery.PanicRecovery()
		defer close(out)

		var (
			errResult = azure.ApplicationResult{}
			nextLink  string
		)

		if list, err := s.GetAzureADApps(ctx, params); err != nil {
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

func (s *azureClient) ListAzureADAppOwners(ctx context.Context, objectId string, params query.GraphParams) <-chan azure.AppOwnerResult {
	out := make(chan azure.AppOwnerResult)

	go func() {
		defer panicrecovery.PanicRecovery()
		defer close(out)

		var (
			errResult = azure.AppOwnerResult{}
			nextLink  string
		)

		if list, err := s.GetAzureADAppOwners(ctx, objectId, params); err != nil {
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
