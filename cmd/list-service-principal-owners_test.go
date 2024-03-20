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

func TestListServicePrincipalOwners(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mockClient := mocks.NewMockAzureClient(ctrl)

	mockServicePrincipalsChannel := make(chan interface{})
	mockServicePrincipalOwnerChannel := make(chan azure.ServicePrincipalOwnerResult)
	mockServicePrincipalOwnerChannel2 := make(chan azure.ServicePrincipalOwnerResult)

	mockTenant := azure.Tenant{}
	mockError := fmt.Errorf("I'm an error")
	mockClient.EXPECT().TenantInfo().Return(mockTenant).AnyTimes()
	mockClient.EXPECT().ListAzureADServicePrincipalOwners(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(mockServicePrincipalOwnerChannel).Times(1)
	mockClient.EXPECT().ListAzureADServicePrincipalOwners(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(mockServicePrincipalOwnerChannel2).Times(1)
	channel := listServicePrincipalOwners(ctx, mockClient, mockServicePrincipalsChannel)

	go func() {
		defer close(mockServicePrincipalsChannel)
		mockServicePrincipalsChannel <- AzureWrapper{
			Data: models.ServicePrincipal{},
		}
		mockServicePrincipalsChannel <- AzureWrapper{
			Data: models.ServicePrincipal{},
		}
	}()
	go func() {
		defer close(mockServicePrincipalOwnerChannel)
		mockServicePrincipalOwnerChannel <- azure.ServicePrincipalOwnerResult{
			Ok: json.RawMessage{},
		}
		mockServicePrincipalOwnerChannel <- azure.ServicePrincipalOwnerResult{
			Ok: json.RawMessage{},
		}
	}()
	go func() {
		defer close(mockServicePrincipalOwnerChannel2)
		mockServicePrincipalOwnerChannel2 <- azure.ServicePrincipalOwnerResult{
			Ok: json.RawMessage{},
		}
		mockServicePrincipalOwnerChannel2 <- azure.ServicePrincipalOwnerResult{
			Error: mockError,
		}
	}()

	if result, ok := <-channel; !ok {
		t.Fatalf("failed to receive from channel")
	} else if wrapper, ok := result.(AzureWrapper); !ok {
		t.Errorf("failed type assertion: got %T, want %T", result, AzureWrapper{})
	} else if data, ok := wrapper.Data.(models.ServicePrincipalOwners); !ok {
		t.Errorf("failed type assertion: got %T, want %T", wrapper.Data, models.ServicePrincipalOwners{})
	} else if len(data.Owners) != 2 {
		t.Errorf("got %v, want %v", len(data.Owners), 2)
	}

	if result, ok := <-channel; !ok {
		t.Fatalf("failed to receive from channel")
	} else if wrapper, ok := result.(AzureWrapper); !ok {
		t.Errorf("failed type assertion: got %T, want %T", result, AzureWrapper{})
	} else if data, ok := wrapper.Data.(models.ServicePrincipalOwners); !ok {
		t.Errorf("failed type assertion: got %T, want %T", wrapper.Data, models.ServicePrincipalOwners{})
	} else if len(data.Owners) != 1 {
		t.Errorf("got %v, want %v", len(data.Owners), 1)
	}
}
