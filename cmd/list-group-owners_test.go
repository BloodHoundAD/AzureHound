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

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bloodhoundad/azurehound/v2/client/mocks"
	"github.com/bloodhoundad/azurehound/v2/models"
	"github.com/bloodhoundad/azurehound/v2/models/azure"
	"go.uber.org/mock/gomock"
)

func init() {
	setupLogger()
}

func TestListGroupOwners(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mockClient := mocks.NewMockAzureClient(ctrl)

	mockGroupsChannel := make(chan interface{})
	mockGroupOwnerChannel := make(chan azure.GroupOwnerResult)
	mockGroupOwnerChannel2 := make(chan azure.GroupOwnerResult)

	mockTenant := azure.Tenant{}
	mockError := fmt.Errorf("I'm an error")
	mockClient.EXPECT().TenantInfo().Return(mockTenant).AnyTimes()
	mockClient.EXPECT().ListAzureADGroupOwners(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(mockGroupOwnerChannel).Times(1)
	mockClient.EXPECT().ListAzureADGroupOwners(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(mockGroupOwnerChannel2).Times(1)
	channel := listGroupOwners(ctx, mockClient, mockGroupsChannel)

	go func() {
		defer close(mockGroupsChannel)
		mockGroupsChannel <- AzureWrapper{
			Data: models.Group{},
		}
		mockGroupsChannel <- AzureWrapper{
			Data: models.Group{},
		}
	}()
	go func() {
		defer close(mockGroupOwnerChannel)
		mockGroupOwnerChannel <- azure.GroupOwnerResult{
			Ok: json.RawMessage{},
		}
		mockGroupOwnerChannel <- azure.GroupOwnerResult{
			Ok: json.RawMessage{},
		}
	}()
	go func() {
		defer close(mockGroupOwnerChannel2)
		mockGroupOwnerChannel2 <- azure.GroupOwnerResult{
			Ok: json.RawMessage{},
		}
		mockGroupOwnerChannel2 <- azure.GroupOwnerResult{
			Error: mockError,
		}
	}()

	if result, ok := <-channel; !ok {
		t.Fatalf("failed to receive from channel")
	} else if wrapper, ok := result.(AzureWrapper); !ok {
		t.Errorf("failed type assertion: got %T, want %T", result, AzureWrapper{})
	} else if data, ok := wrapper.Data.(models.GroupOwners); !ok {
		t.Errorf("failed type assertion: got %T, want %T", wrapper.Data, models.GroupOwners{})
	} else if len(data.Owners) != 2 {
		t.Errorf("got %v, want %v", len(data.Owners), 2)
	}

	if result, ok := <-channel; !ok {
		t.Fatalf("failed to receive from channel")
	} else if wrapper, ok := result.(AzureWrapper); !ok {
		t.Errorf("failed type assertion: got %T, want %T", result, AzureWrapper{})
	} else if data, ok := wrapper.Data.(models.GroupOwners); !ok {
		t.Errorf("failed type assertion: got %T, want %T", wrapper.Data, models.GroupOwners{})
	} else if len(data.Owners) != 1 {
		t.Errorf("got %v, want %v", len(data.Owners), 2)
	}
}
