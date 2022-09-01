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
	"github.com/bloodhoundad/azurehound/config"
	"github.com/bloodhoundad/azurehound/enums"
	"github.com/bloodhoundad/azurehound/pipeline"
	"github.com/spf13/cobra"
)

func init() {
	config.Init(listRootCmd, append(config.AzureConfig, config.OutputFile))
	rootCmd.AddCommand(listRootCmd)
}

var listRootCmd = &cobra.Command{
	Use:               "list",
	Short:             "Lists Azure Objects",
	Run:               listCmdImpl,
	PersistentPreRunE: persistentPreRunE,
	SilenceUsage:      true,
}

func listCmdImpl(cmd *cobra.Command, args []string) {
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
		log.Info("collecting azure objects...")
		start := time.Now()
		stream := listAll(ctx, azClient)
		outputStream(ctx, stream)
		duration := time.Since(start)
		log.Info("collection completed", "duration", duration.String())
	}
}

func listAll(ctx context.Context, client client.AzureClient) <-chan interface{} {
	var (
		apps  = make(chan interface{})
		apps2 = make(chan interface{})

		devices  = make(chan interface{})
		devices2 = make(chan interface{})

		groups  = make(chan interface{})
		groups2 = make(chan interface{})
		groups3 = make(chan interface{})

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

		roles  = make(chan interface{})
		roles2 = make(chan interface{})

		servicePrincipals  = make(chan interface{})
		servicePrincipals2 = make(chan interface{})

		subscriptions   = make(chan interface{})
		subscriptions2  = make(chan interface{})
		subscriptions3  = make(chan interface{})
		subscriptions4  = make(chan interface{})
		subscriptions5  = make(chan interface{})
		subscriptions6  = make(chan interface{})
		subscriptions7  = make(chan interface{})
		subscriptions8  = make(chan interface{})
		subscriptions9  = make(chan interface{})
		subscriptions10 = make(chan interface{})

		tenants = make(chan interface{})

		virtualMachines  = make(chan interface{})
		virtualMachines2 = make(chan interface{})
		virtualMachines3 = make(chan interface{})
		virtualMachines4 = make(chan interface{})
		virtualMachines5 = make(chan interface{})
		virtualMachines6 = make(chan interface{})

		storageAccounts  = make(chan interface{})
		storageAccounts2 = make(chan interface{})
		storageAccounts3 = make(chan interface{})

		automationAccounts  = make(chan interface{})
		automationAccounts2 = make(chan interface{})

		workflows  = make(chan interface{})
		workflows2 = make(chan interface{})

		functionApps  = make(chan interface{})
		functionApps2 = make(chan interface{})
	)

	// Enumerate Apps, AppOwners and AppMembers
	pipeline.Tee(ctx.Done(), listApps(ctx, client), apps, apps2)
	appOwners := listAppOwners(ctx, client, apps2)

	// Enumerate Devices and DeviceOwners
	pipeline.Tee(ctx.Done(), listDevices(ctx, client), devices, devices2)
	deviceOwners := listDeviceOwners(ctx, client, devices2)

	// Enumerate Groups, GroupOwners and GroupMembers
	pipeline.Tee(ctx.Done(), listGroups(ctx, client), groups, groups2, groups3)
	groupOwners := listGroupOwners(ctx, client, groups2)
	groupMembers := listGroupMembers(ctx, client, groups3)

	// Enumerate Subscriptions, SubscriptionOwners and SubscriptionUserAccessAdmins
	pipeline.Tee(ctx.Done(), listSubscriptions(ctx, client), subscriptions, subscriptions2, subscriptions3, subscriptions4, subscriptions5, subscriptions6, subscriptions7, subscriptions8, subscriptions9, subscriptions10)
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

	// Enumerate ServicePrincipals and ServicePrincipalOwners
	pipeline.Tee(ctx.Done(), listServicePrincipals(ctx, client), servicePrincipals, servicePrincipals2)
	servicePrincipalOwners := listServicePrincipalOwners(ctx, client, servicePrincipals2)

	// Enumerate Tenants
	pipeline.Tee(ctx.Done(), listTenants(ctx, client), tenants)

	// Enumerate Users
	users := listUsers(ctx, client)

	// Enumerate Roles and RoleAssignments
	pipeline.Tee(ctx.Done(), listRoles(ctx, client), roles, roles2)
	roleAssignments := listRoleAssignments(ctx, client, roles2)

	// Enumerate VirtualMachines, VirtualMachineOwners, VirtualMachineAvereContributors, VirtualMachineContributors,
	// VirtualMachineAdminLogins and VirtualMachineUserAccessAdmins
	pipeline.Tee(ctx.Done(), listVirtualMachines(ctx, client, subscriptions4), virtualMachines, virtualMachines2, virtualMachines3, virtualMachines4, virtualMachines5, virtualMachines6)
	virtualMachineOwners := listVirtualMachineOwners(ctx, client, virtualMachines2)
	virtualMachineAvereContributors := listVirtualMachineAvereContributors(ctx, client, virtualMachines3)
	virtualMachineContributors := listVirtualMachineContributors(ctx, client, virtualMachines4)
	virtualMachineAdminLogins := listVirtualMachineAdminLogins(ctx, client, virtualMachines5)
	virtualMachineUserAccessAdmins := listVirtualMachineUserAccessAdmins(ctx, client, virtualMachines6)

	//Enumerate storage accounts
	pipeline.Tee(ctx.Done(), listStorageAccounts(ctx, client, subscriptions7), storageAccounts, storageAccounts2, storageAccounts3)
	storageContainers := listStorageContainers(ctx, client, storageAccounts2)
	storageAccountRoleAssignments := listStorageAccountRoleAssignments(ctx, client, storageAccounts3)

	//Enumerage automation accounts
	pipeline.Tee(ctx.Done(), listAutomationAccounts(ctx, client, subscriptions8), automationAccounts, automationAccounts2)
	automationAccountRoleAssignments := listAutomationAccountRoleAssignments(ctx, client, automationAccounts2)

	//Enumerate workflows / logic apps
	pipeline.Tee(ctx.Done(), listWorkflows(ctx, client, subscriptions9), workflows, workflows2)
	workflowRoleAssignments := listWorkflowRoleAsignments(ctx, client, workflows2)

	//Enumerate function apps
	pipeline.Tee(ctx.Done(), listFunctionApps(ctx, client, subscriptions10), functionApps, functionApps2)
	functionAppRoleAssignments := listFunctionAppRoleAssignments(ctx, client, functionApps2)

	return pipeline.Mux(ctx.Done(),
		appOwners,
		apps,
		deviceOwners,
		devices,
		groupMembers,
		groupOwners,
		groups,
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
		roleAssignments,
		roles,
		servicePrincipalOwners,
		servicePrincipals,
		subscriptionOwners,
		subscriptionUserAccessAdmins,
		subscriptions,
		tenants,
		users,
		virtualMachineAdminLogins,
		virtualMachineAvereContributors,
		virtualMachineContributors,
		virtualMachineOwners,
		virtualMachineUserAccessAdmins,
		virtualMachines,
		storageAccounts,
		storageContainers,
		storageAccountRoleAssignments,
		automationAccounts,
		automationAccountRoleAssignments,
		workflows,
		workflowRoleAssignments,
		functionApps,
		functionAppRoleAssignments,
	)
}
