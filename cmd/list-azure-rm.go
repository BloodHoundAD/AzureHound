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
	"os"
	"os/signal"
	"time"

	"github.com/bloodhoundad/azurehound/client"
	"github.com/bloodhoundad/azurehound/enums"
	"github.com/bloodhoundad/azurehound/models"
	"github.com/bloodhoundad/azurehound/pipeline"
	"github.com/spf13/cobra"
)

func init() {
	listRootCmd.AddCommand(listAzureRMCmd)
}

var listAzureRMCmd = &cobra.Command{
	Use:               "az-rm",
	Long:              "Lists All Azure RM Entities",
	PersistentPreRunE: persistentPreRunE,
	Run:               listAzureRMCmdImpl,
	SilenceUsage:      true,
}

func listAzureRMCmdImpl(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		exit(fmt.Errorf("unsupported subcommand: %v", args))
	}

	ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, os.Kill)
	defer gracefulShutdown(stop)

	log.V(1).Info("testing connections")
	azClient := connectAndCreateClient()
	log.Info("collecting azure resource management objects...")
	start := time.Now()
	stream := listAllRM(ctx, azClient)
	outputStream(ctx, stream)
	duration := time.Since(start)
	log.Info("collection completed", "duration", duration.String())
}

func listAllRM(ctx context.Context, client client.AzureClient) <-chan interface{} {
	var (
		keyVaults                = make(chan interface{})
		keyVaults2               = make(chan interface{})
		keyVaults3               = make(chan interface{})
		keyVaultRoleAssignments1 = make(chan azureWrapper[models.KeyVaultRoleAssignments])
		keyVaultRoleAssignments2 = make(chan azureWrapper[models.KeyVaultRoleAssignments])
		keyVaultRoleAssignments3 = make(chan azureWrapper[models.KeyVaultRoleAssignments])
		keyVaultRoleAssignments4 = make(chan azureWrapper[models.KeyVaultRoleAssignments])

		mgmtGroups                = make(chan interface{})
		mgmtGroups2               = make(chan interface{})
		mgmtGroups3               = make(chan interface{})
		mgmtGroupRoleAssignments1 = make(chan azureWrapper[models.ManagementGroupRoleAssignments])
		mgmtGroupRoleAssignments2 = make(chan azureWrapper[models.ManagementGroupRoleAssignments])

		resourceGroups                = make(chan interface{})
		resourceGroups2               = make(chan interface{})
		resourceGroupRoleAssignments1 = make(chan azureWrapper[models.ResourceGroupRoleAssignments])
		resourceGroupRoleAssignments2 = make(chan azureWrapper[models.ResourceGroupRoleAssignments])

		subscriptions                = make(chan interface{})
		subscriptions2               = make(chan interface{})
		subscriptions3               = make(chan interface{})
		subscriptions4               = make(chan interface{})
		subscriptions5               = make(chan interface{})
		subscriptionRoleAssignments1 = make(chan interface{})
		subscriptionRoleAssignments2 = make(chan interface{})

		virtualMachines                = make(chan interface{})
		virtualMachines2               = make(chan interface{})
		virtualMachineRoleAssignments1 = make(chan azureWrapper[models.VirtualMachineRoleAssignments])
		virtualMachineRoleAssignments2 = make(chan azureWrapper[models.VirtualMachineRoleAssignments])
		virtualMachineRoleAssignments3 = make(chan azureWrapper[models.VirtualMachineRoleAssignments])
		virtualMachineRoleAssignments4 = make(chan azureWrapper[models.VirtualMachineRoleAssignments])
		virtualMachineRoleAssignments5 = make(chan azureWrapper[models.VirtualMachineRoleAssignments])
	)

	// Enumerate entities
	pipeline.Tee(ctx.Done(), listManagementGroups(ctx, client), mgmtGroups, mgmtGroups2, mgmtGroups3)
	pipeline.Tee(ctx.Done(), listSubscriptions(ctx, client), subscriptions, subscriptions2, subscriptions3, subscriptions4, subscriptions5)
	pipeline.Tee(ctx.Done(), listResourceGroups(ctx, client, subscriptions2), resourceGroups, resourceGroups2)
	pipeline.Tee(ctx.Done(), listKeyVaults(ctx, client, subscriptions3), keyVaults, keyVaults2, keyVaults3)
	pipeline.Tee(ctx.Done(), listVirtualMachines(ctx, client, subscriptions4), virtualMachines, virtualMachines2)

	// Enumerate Relationships
	// ManagementGroups: Descendants, Owners and UserAccessAdmins
	mgmtGroupDescendants := listManagementGroupDescendants(ctx, client, mgmtGroups2)
	pipeline.Tee(ctx.Done(), listManagementGroupRoleAssignments(ctx, client, mgmtGroups3), mgmtGroupRoleAssignments1, mgmtGroupRoleAssignments2)
	mgmtGroupOwners := listManagementGroupOwners(ctx, mgmtGroupRoleAssignments1)
	mgmtGroupUserAccessAdmins := listManagementGroupUserAccessAdmins(ctx, mgmtGroupRoleAssignments2)

	// Subscriptions: Owners and UserAccessAdmins
	pipeline.Tee(ctx.Done(), listSubscriptionRoleAssignments(ctx, client, subscriptions5), subscriptionRoleAssignments1, subscriptionRoleAssignments2)
	subscriptionOwners := listSubscriptionOwners(ctx, client, subscriptionRoleAssignments1)
	subscriptionUserAccessAdmins := listSubscriptionUserAccessAdmins(ctx, client, subscriptionRoleAssignments2)

	// ResourceGroups: Owners and UserAccessAdmins
	pipeline.Tee(ctx.Done(), listResourceGroupRoleAssignments(ctx, client, resourceGroups2), resourceGroupRoleAssignments1, resourceGroupRoleAssignments2)
	resourceGroupOwners := listResourceGroupOwners(ctx, resourceGroupRoleAssignments1)
	resourceGroupUserAccessAdmins := listResourceGroupUserAccessAdmins(ctx, resourceGroupRoleAssignments2)

	// KeyVaults: AccessPolicies, Owners, UserAccessAdmins, Contributors and KVContributors
	pipeline.Tee(ctx.Done(), listKeyVaultRoleAssignments(ctx, client, keyVaults2), keyVaultRoleAssignments1, keyVaultRoleAssignments2, keyVaultRoleAssignments3, keyVaultRoleAssignments4)
	keyVaultAccessPolicies := listKeyVaultAccessPolicies(ctx, client, keyVaults3, []enums.KeyVaultAccessType{enums.GetCerts, enums.GetKeys, enums.GetCerts})
	keyVaultOwners := listKeyVaultOwners(ctx, keyVaultRoleAssignments1)
	keyVaultUserAccessAdmins := listKeyVaultUserAccessAdmins(ctx, keyVaultRoleAssignments2)
	keyVaultContributors := listKeyVaultContributors(ctx, keyVaultRoleAssignments3)
	keyVaultKVContributors := listKeyVaultKVContributors(ctx, keyVaultRoleAssignments4)

	// VirtualMachines: Owners, AvereContributors, Contributors, AdminLogins and UserAccessAdmins
	pipeline.Tee(ctx.Done(), listVirtualMachineRoleAssignments(ctx, client, virtualMachines2), virtualMachineRoleAssignments1, virtualMachineRoleAssignments2, virtualMachineRoleAssignments3, virtualMachineRoleAssignments4, virtualMachineRoleAssignments5)
	virtualMachineOwners := listVirtualMachineOwners(ctx, virtualMachineRoleAssignments1)
	virtualMachineAvereContributors := listVirtualMachineAvereContributors(ctx, virtualMachineRoleAssignments2)
	virtualMachineContributors := listVirtualMachineContributors(ctx, virtualMachineRoleAssignments3)
	virtualMachineAdminLogins := listVirtualMachineAdminLogins(ctx, virtualMachineRoleAssignments4)
	virtualMachineUserAccessAdmins := listVirtualMachineUserAccessAdmins(ctx, virtualMachineRoleAssignments5)

	return pipeline.Mux(ctx.Done(),
		keyVaultAccessPolicies,
		keyVaultContributors,
		keyVaultKVContributors,
		keyVaultOwners,
		keyVaultUserAccessAdmins,
		keyVaults,
		mgmtGroupDescendants,
		mgmtGroupOwners,
		mgmtGroupUserAccessAdmins,
		mgmtGroups,
		resourceGroupOwners,
		resourceGroupUserAccessAdmins,
		resourceGroups,
		subscriptionOwners,
		subscriptionUserAccessAdmins,
		subscriptions,
		virtualMachineAdminLogins,
		virtualMachineAvereContributors,
		virtualMachineContributors,
		virtualMachineOwners,
		virtualMachineUserAccessAdmins,
		virtualMachines,
	)
}
