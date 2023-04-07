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
	"encoding/json"
	"fmt"
	"io"

	"github.com/bloodhoundad/azurehound/v2/constants"
)

func (s *azureClient) GetAzureADDirectoryObject(ctx context.Context, objectId string) (json.RawMessage, error) {
	var (
		path = fmt.Sprintf("/%s/directoryObjects/%s", constants.GraphApiVersion, objectId)
	)
	if res, err := s.msgraph.Get(ctx, path, nil, nil); err != nil {
		return nil, err
	} else if body, err := io.ReadAll(res.Body); err != nil {
		res.Body.Close()
		return nil, err
	} else {
		res.Body.Close()
		return json.RawMessage(body), nil
	}
}
