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
	"os"
	"os/signal"
	"time"

	"github.com/bloodhoundad/azurehound/client"
	"github.com/bloodhoundad/azurehound/enums"
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
	if err := testConnections(); err != nil {
		exit(err)
	} else if azClient, err := newAzureClient(); err != nil {
		exit(err)
	} else {
		log.Info("collecting azure resource management objects...")
		start := time.Now()
		stream := listAllRM(ctx, azClient)
		outputStream(ctx, stream)
		duration := time.Since(start)
		log.Info("collection completed", "duration", duration.String())
	}
}

func listAllRM(ctx context.Context, client client.AzureClient) <-chan interface{} {
	var (
		keyVaults  = make(chan interface{})
		keyVaults2 = make(chan interface{})
		keyVaults3 = make(chan interface{})
		keyVaults4 = make(chan interface{})

		mgmtGroups  = make(chan interface{})
		mgmtGroups2 = make(chan interface{})
		mgmtGroups3 = make(chan interface{})
		mgmtGroups4 = make(chan interface{})

		resourceGroups  = make(chan interface{})
		resourceGroups2 = make(chan interface{})
		resourceGroups3 = make(chan interface{})

		subscriptions  = make(chan interface{})
		subscriptions2 = make(chan interface{})
		subscriptions3 = make(chan interface{})
		subscriptions4 = make(chan interface{})
		subscriptions5 = make(chan interface{})
		subscriptions6 = make(chan interface{})

		virtualMachines  = make(chan interface{})
		virtualMachines2 = make(chan interface{})
		virtualMachines3 = make(chan interface{})
		virtualMachines4 = make(chan interface{})
		virtualMachines5 = make(chan interface{})
		virtualMachines6 = make(chan interface{})
	)

	// Enumerate Subscriptions, SubscriptionOwners and SubscriptionUserAccessAdmins
	pipeline.Tee(ctx.Done(), listSubscriptions(ctx, client), subscriptions, subscriptions2, subscriptions3, subscriptions4, subscriptions5, subscriptions6)
	subscriptionOwners := listSubscriptionOwners(ctx, client, subscriptions5)
	subscriptionUserAccessAdmins := listSubscriptionUserAccessAdmins(ctx, client, subscriptions6)

	// Enumerate KeyVaults, KeyVaultOwners, KeyVaultAccessPolicies and KeyVaultUserAccessAdmins
	pipeline.Tee(ctx.Done(), listKeyVaults(ctx, client, subscriptions2), keyVaults, keyVaults2, keyVaults3, keyVaults4)
	keyVaultOwners := listKeyVaultOwners(ctx, client, keyVaults2)
	keyVaultAccessPolicies := listKeyVaultAccessPolicies(ctx, client, keyVaults3, []enums.KeyVaultAccessType{enums.GetCerts, enums.GetKeys, enums.GetCerts})
	keyVaultUserAccessAdmins := listKeyVaultUserAccessAdmins(ctx, client, keyVaults4)

	// Enumerate ManagementGroups, ManagementGroupOwners and ManagementGroupDescendants
	pipeline.Tee(ctx.Done(), listManagementGroups(ctx, client), mgmtGroups, mgmtGroups2, mgmtGroups3, mgmtGroups4)
	mgmtGroupOwners := listManagementGroupOwners(ctx, client, mgmtGroups2)
	mgmtGroupDescendants := listManagementGroupDescendants(ctx, client, mgmtGroups3)
	mgmtGroupUserAccessAdmins := listManagementGroupUserAccessAdmins(ctx, client, mgmtGroups4)

	// Enumerate ResourceGroups, ResourceGroupOwners and ResourceGroupUserAccessAdmins
	pipeline.Tee(ctx.Done(), listResourceGroups(ctx, client, subscriptions3), resourceGroups, resourceGroups2, resourceGroups3)
	resourceGroupOwners := listResourceGroupOwners(ctx, client, resourceGroups2)
	resourceGroupUserAccessAdmins := listResourceGroupUserAccessAdmins(ctx, client, resourceGroups3)

	// Enumerate VirtualMachines, VirtualMachineOwners, VirtualMachineAvereContributors, VirtualMachineContributors,
	// VirtualMachineAdminLogins and VirtualMachineUserAccessAdmins
	pipeline.Tee(ctx.Done(), listVirtualMachines(ctx, client, subscriptions4), virtualMachines, virtualMachines2, virtualMachines3, virtualMachines4, virtualMachines5, virtualMachines6)
	virtualMachineOwners := listVirtualMachineOwners(ctx, client, virtualMachines2)
	virtualMachineAvereContributors := listVirtualMachineAvereContributors(ctx, client, virtualMachines3)
	virtualMachineContributors := listVirtualMachineContributors(ctx, client, virtualMachines4)
	virtualMachineAdminLogins := listVirtualMachineAdminLogins(ctx, client, virtualMachines5)
	virtualMachineUserAccessAdmins := listVirtualMachineUserAccessAdmins(ctx, client, virtualMachines6)

	return pipeline.Mux(ctx.Done(),
		keyVaultAccessPolicies,
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
