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

func TestListManagementGroupUserAccessAdmins(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mockClient := mocks.NewMockAzureClient(ctrl)

	mockManagementGroupsChannel := make(chan interface{})
	mockManagementGroupUserAccessAdminChannel := make(chan azure.RoleAssignmentResult)
	mockManagementGroupUserAccessAdminChannel2 := make(chan azure.RoleAssignmentResult)

	mockTenant := azure.Tenant{}
	mockError := fmt.Errorf("I'm an error")
	mockClient.EXPECT().TenantInfo().Return(mockTenant).AnyTimes()
	mockClient.EXPECT().ListRoleAssignmentsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockManagementGroupUserAccessAdminChannel).Times(1)
	mockClient.EXPECT().ListRoleAssignmentsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockManagementGroupUserAccessAdminChannel2).Times(1)
	channel := listManagementGroupUserAccessAdmins(ctx, mockClient, mockManagementGroupsChannel)

	go func() {
		defer close(mockManagementGroupsChannel)
		mockManagementGroupsChannel <- AzureWrapper{
			Data: models.ManagementGroup{},
		}
		mockManagementGroupsChannel <- AzureWrapper{
			Data: models.ManagementGroup{},
		}
	}()
	go func() {
		defer close(mockManagementGroupUserAccessAdminChannel)
		mockManagementGroupUserAccessAdminChannel <- azure.RoleAssignmentResult{
			Ok: azure.RoleAssignment{
				Properties: azure.RoleAssignmentPropertiesWithScope{
					RoleDefinitionId: constants.UserAccessAdminRoleID,
				},
			},
		}
		mockManagementGroupUserAccessAdminChannel <- azure.RoleAssignmentResult{
			Ok: azure.RoleAssignment{
				Properties: azure.RoleAssignmentPropertiesWithScope{
					RoleDefinitionId: constants.UserAccessAdminRoleID,
				},
			},
		}
	}()
	go func() {
		defer close(mockManagementGroupUserAccessAdminChannel2)
		mockManagementGroupUserAccessAdminChannel2 <- azure.RoleAssignmentResult{
			Ok: azure.RoleAssignment{
				Properties: azure.RoleAssignmentPropertiesWithScope{
					RoleDefinitionId: constants.UserAccessAdminRoleID,
				},
			},
		}
		mockManagementGroupUserAccessAdminChannel2 <- azure.RoleAssignmentResult{
			Error: mockError,
		}
	}()

	if result, ok := <-channel; !ok {
		t.Fatalf("failed to receive from channel")
	} else if wrapper, ok := result.(AzureWrapper); !ok {
		t.Errorf("failed type assertion: got %T, want %T", result, AzureWrapper{})
	} else if data, ok := wrapper.Data.(models.ManagementGroupUserAccessAdmins); !ok {
		t.Errorf("failed type assertion: got %T, want %T", wrapper.Data, models.ManagementGroupUserAccessAdmins{})
	} else if len(data.UserAccessAdmins) != 2 {
		t.Errorf("got %v, want %v", len(data.UserAccessAdmins), 2)
	}

	if result, ok := <-channel; !ok {
		t.Fatalf("failed to receive from channel")
	} else if wrapper, ok := result.(AzureWrapper); !ok {
		t.Errorf("failed type assertion: got %T, want %T", result, AzureWrapper{})
	} else if data, ok := wrapper.Data.(models.ManagementGroupUserAccessAdmins); !ok {
		t.Errorf("failed type assertion: got %T, want %T", wrapper.Data, models.ManagementGroupUserAccessAdmins{})
	} else if len(data.UserAccessAdmins) != 1 {
		t.Errorf("got %v, want %v", len(data.UserAccessAdmins), 2)
	}
}
