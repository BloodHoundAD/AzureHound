// Copyright (C) 2022 The BloodHound Enterprise Team
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

	"github.com/bloodhoundad/azurehound/client/mocks"
	"github.com/bloodhoundad/azurehound/models"
	"github.com/bloodhoundad/azurehound/models/azure"
	"github.com/golang/mock/gomock"
)

func init() {
	setupLogger()
}

func TestListAppOwners(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mockClient := mocks.NewMockAzureClient(ctrl)

	mockAppsChannel := make(chan interface{})
	mockAppOwnerChannel := make(chan azure.AppOwnerResult)
	mockAppOwnerChannel2 := make(chan azure.AppOwnerResult)

	mockTenant := azure.Tenant{}
	mockError := fmt.Errorf("I'm an error")
	mockClient.EXPECT().TenantInfo().Return(mockTenant).AnyTimes()
	mockClient.EXPECT().ListAzureADAppOwners(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(mockAppOwnerChannel).Times(1)
	mockClient.EXPECT().ListAzureADAppOwners(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(mockAppOwnerChannel2).Times(1)
	channel := listAppOwners(ctx, mockClient, mockAppsChannel)

	go func() {
		defer close(mockAppsChannel)
		mockAppsChannel <- AzureWrapper{
			Data: models.App{},
		}
		mockAppsChannel <- AzureWrapper{
			Data: models.App{},
		}
	}()
	go func() {
		defer close(mockAppOwnerChannel)
		mockAppOwnerChannel <- azure.AppOwnerResult{
			Ok: json.RawMessage{},
		}
		mockAppOwnerChannel <- azure.AppOwnerResult{
			Ok: json.RawMessage{},
		}
	}()
	go func() {
		defer close(mockAppOwnerChannel2)
		mockAppOwnerChannel2 <- azure.AppOwnerResult{
			Ok: json.RawMessage{},
		}
		mockAppOwnerChannel2 <- azure.AppOwnerResult{
			Error: mockError,
		}
	}()

	if result, ok := <-channel; !ok {
		t.Fatalf("failed to receive from channel")
	} else if wrapper, ok := result.(AzureWrapper); !ok {
		t.Errorf("failed type assertion: got %T, want %T", result, AzureWrapper{})
	} else if data, ok := wrapper.Data.(models.AppOwners); !ok {
		t.Errorf("failed type assertion: got %T, want %T", wrapper.Data, models.AppOwners{})
	} else if len(data.Owners) != 2 {
		t.Errorf("got %v, want %v", len(data.Owners), 2)
	}

	if result, ok := <-channel; !ok {
		t.Fatalf("failed to receive from channel")
	} else if wrapper, ok := result.(AzureWrapper); !ok {
		t.Errorf("failed type assertion: got %T, want %T", result, AzureWrapper{})
	} else if data, ok := wrapper.Data.(models.AppOwners); !ok {
		t.Errorf("failed type assertion: got %T, want %T", wrapper.Data, models.AppOwners{})
	} else if len(data.Owners) != 1 {
		t.Errorf("got %v, want %v", len(data.Owners), 2)
	}
}
