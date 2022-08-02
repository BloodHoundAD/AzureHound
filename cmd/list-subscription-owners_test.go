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
	"fmt"
	"testing"

	"github.com/bloodhoundad/azurehound/client/mocks"
	"github.com/bloodhoundad/azurehound/constants"
	"github.com/bloodhoundad/azurehound/models"
	"github.com/bloodhoundad/azurehound/models/azure"
	"github.com/golang/mock/gomock"
)

func init() {
	setupLogger()
}

func TestListSubscriptionOwners(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mockClient := mocks.NewMockAzureClient(ctrl)

	mockSubscriptionsChannel := make(chan interface{})
	mockSubscriptionOwnerChannel := make(chan azure.RoleAssignmentResult)
	mockSubscriptionOwnerChannel2 := make(chan azure.RoleAssignmentResult)

	mockTenant := azure.Tenant{}
	mockError := fmt.Errorf("I'm an error")
	mockClient.EXPECT().TenantInfo().Return(mockTenant).AnyTimes()
	mockClient.EXPECT().ListRoleAssignmentsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockSubscriptionOwnerChannel).Times(1)
	mockClient.EXPECT().ListRoleAssignmentsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockSubscriptionOwnerChannel2).Times(1)
	channel := listSubscriptionOwners(ctx, mockClient, mockSubscriptionsChannel)

	go func() {
		defer close(mockSubscriptionsChannel)
		mockSubscriptionsChannel <- AzureWrapper{
			Data: models.Subscription{},
		}
		mockSubscriptionsChannel <- AzureWrapper{
			Data: models.Subscription{},
		}
	}()
	go func() {
		defer close(mockSubscriptionOwnerChannel)
		mockSubscriptionOwnerChannel <- azure.RoleAssignmentResult{
			Ok: azure.RoleAssignment{
				Properties: azure.RoleAssignmentPropertiesWithScope{
					RoleDefinitionId: constants.OwnerRoleID,
				},
			},
		}
		mockSubscriptionOwnerChannel <- azure.RoleAssignmentResult{
			Ok: azure.RoleAssignment{
				Properties: azure.RoleAssignmentPropertiesWithScope{
					RoleDefinitionId: constants.OwnerRoleID,
				},
			},
		}
	}()
	go func() {
		defer close(mockSubscriptionOwnerChannel2)
		mockSubscriptionOwnerChannel2 <- azure.RoleAssignmentResult{
			Ok: azure.RoleAssignment{
				Properties: azure.RoleAssignmentPropertiesWithScope{
					RoleDefinitionId: constants.OwnerRoleID,
				},
			},
		}
		mockSubscriptionOwnerChannel2 <- azure.RoleAssignmentResult{
			Error: mockError,
		}
	}()

	if result, ok := <-channel; !ok {
		t.Fatalf("failed to receive from channel")
	} else if wrapper, ok := result.(AzureWrapper); !ok {
		t.Errorf("failed type assertion: got %T, want %T", result, AzureWrapper{})
	} else if data, ok := wrapper.Data.(models.SubscriptionOwners); !ok {
		t.Errorf("failed type assertion: got %T, want %T", wrapper.Data, models.SubscriptionOwners{})
	} else if len(data.Owners) != 2 {
		t.Errorf("got %v, want %v", len(data.Owners), 2)
	}

	if result, ok := <-channel; !ok {
		t.Fatalf("failed to receive from channel")
	} else if wrapper, ok := result.(AzureWrapper); !ok {
		t.Errorf("failed type assertion: got %T, want %T", result, AzureWrapper{})
	} else if data, ok := wrapper.Data.(models.SubscriptionOwners); !ok {
		t.Errorf("failed type assertion: got %T, want %T", wrapper.Data, models.SubscriptionOwners{})
	} else if len(data.Owners) != 1 {
		t.Errorf("got %v, want %v", len(data.Owners), 2)
	}
}
