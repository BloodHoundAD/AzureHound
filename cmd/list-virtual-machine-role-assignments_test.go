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
	"fmt"
	"testing"

	"github.com/bloodhoundad/azurehound/v2/client/mocks"
	"github.com/bloodhoundad/azurehound/v2/constants"
	"github.com/bloodhoundad/azurehound/v2/models"
	"github.com/bloodhoundad/azurehound/v2/models/azure"
	"go.uber.org/mock/gomock"
)

func init() {
	setupLogger()
}

func TestListVirtualMachineRoleAssignments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mockClient := mocks.NewMockAzureClient(ctrl)

	mockVirtualMachinesChannel := make(chan interface{})
	mockVirtualMachineRoleAssignmentChannel := make(chan azure.RoleAssignmentResult)
	mockVirtualMachineRoleAssignmentChannel2 := make(chan azure.RoleAssignmentResult)

	mockTenant := azure.Tenant{}
	mockError := fmt.Errorf("I'm an error")
	mockClient.EXPECT().TenantInfo().Return(mockTenant).AnyTimes()
	mockClient.EXPECT().ListRoleAssignmentsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockVirtualMachineRoleAssignmentChannel).Times(1)
	mockClient.EXPECT().ListRoleAssignmentsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockVirtualMachineRoleAssignmentChannel2).Times(1)
	channel := listVirtualMachineRoleAssignments(ctx, mockClient, mockVirtualMachinesChannel)

	go func() {
		defer close(mockVirtualMachinesChannel)
		mockVirtualMachinesChannel <- AzureWrapper{
			Data: models.VirtualMachine{},
		}
		mockVirtualMachinesChannel <- AzureWrapper{
			Data: models.VirtualMachine{},
		}
	}()
	go func() {
		defer close(mockVirtualMachineRoleAssignmentChannel)
		mockVirtualMachineRoleAssignmentChannel <- azure.RoleAssignmentResult{
			Ok: azure.RoleAssignment{
				Properties: azure.RoleAssignmentPropertiesWithScope{
					RoleDefinitionId: constants.VirtualMachineContributorRoleID,
				},
			},
		}
		mockVirtualMachineRoleAssignmentChannel <- azure.RoleAssignmentResult{
			Ok: azure.RoleAssignment{
				Properties: azure.RoleAssignmentPropertiesWithScope{
					RoleDefinitionId: constants.AvereContributorRoleID,
				},
			},
		}
	}()
	go func() {
		defer close(mockVirtualMachineRoleAssignmentChannel2)
		mockVirtualMachineRoleAssignmentChannel2 <- azure.RoleAssignmentResult{
			Ok: azure.RoleAssignment{
				Properties: azure.RoleAssignmentPropertiesWithScope{
					RoleDefinitionId: constants.VirtualMachineAdministratorLoginRoleID,
				},
			},
		}
		mockVirtualMachineRoleAssignmentChannel2 <- azure.RoleAssignmentResult{
			Error: mockError,
		}
	}()

	if result, ok := <-channel; !ok {
		t.Fatalf("failed to receive from channel")
	} else if len(result.Data.RoleAssignments) != 2 {
		t.Errorf("got %v, want %v", len(result.Data.RoleAssignments), 2)
	}

	if result, ok := <-channel; !ok {
		t.Fatalf("failed to receive from channel")
	} else if len(result.Data.RoleAssignments) != 1 {
		t.Errorf("got %v, want %v", len(result.Data.RoleAssignments), 2)
	}
}
