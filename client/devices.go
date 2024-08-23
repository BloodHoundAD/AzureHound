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

func (s *azureClient) GetAzureDeviceRegisteredOwners(ctx context.Context, objectId string, params query.GraphParams) (azure.DirectoryObjectList, error) {
	var (
		path     = fmt.Sprintf("/%s/devices/%s/registeredOwners", constants.GraphApiBetaVersion, objectId)
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

func (s *azureClient) GetAzureDevices(ctx context.Context, params query.GraphParams) (azure.DeviceList, error) {
	var (
		path     = fmt.Sprintf("/%s/devices", constants.GraphApiVersion)
		response azure.DeviceList
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

func (s *azureClient) ListAzureDevices(ctx context.Context, params query.GraphParams) <-chan azure.DeviceResult {
	out := make(chan azure.DeviceResult)

	go func() {
		defer panicrecovery.PanicRecovery()
		defer close(out)

		var (
			errResult = azure.DeviceResult{}
			nextLink  string
		)

		if list, err := s.GetAzureDevices(ctx, params); err != nil {
			errResult.Error = err
			if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
				return
			}
		} else {
			for _, u := range list.Value {
				if ok := pipeline.Send(ctx.Done(), out, azure.DeviceResult{Ok: u}); !ok {
					return
				}
			}

			nextLink = list.NextLink
			for nextLink != "" {
				var list azure.DeviceList
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
						if ok := pipeline.Send(ctx.Done(), out, azure.DeviceResult{Ok: u}); !ok {
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

func (s *azureClient) ListAzureDeviceRegisteredOwners(ctx context.Context, objectId string, params query.GraphParams) <-chan azure.DeviceRegisteredOwnerResult {
	out := make(chan azure.DeviceRegisteredOwnerResult)

	go func() {
		defer panicrecovery.PanicRecovery()
		defer close(out)

		var (
			errResult = azure.DeviceRegisteredOwnerResult{
				DeviceId: objectId,
			}
			nextLink string
		)

		if list, err := s.GetAzureDeviceRegisteredOwners(ctx, objectId, params); err != nil {
			errResult.Error = err
			if ok := pipeline.Send(ctx.Done(), out, errResult); !ok {
				return
			}
		} else {
			for _, u := range list.Value {
				if ok := pipeline.Send(ctx.Done(), out, azure.DeviceRegisteredOwnerResult{
					DeviceId: objectId,
					Ok:       u,
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
						if ok := pipeline.Send(ctx.Done(), out, azure.DeviceRegisteredOwnerResult{
							DeviceId: objectId,
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
