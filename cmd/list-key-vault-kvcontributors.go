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
	"os"
	"os/signal"
	"time"

	"github.com/bloodhoundad/azurehound/v2/constants"
	"github.com/bloodhoundad/azurehound/v2/enums"
	"github.com/bloodhoundad/azurehound/v2/internal"
	"github.com/bloodhoundad/azurehound/v2/models"
	"github.com/bloodhoundad/azurehound/v2/pipeline"
	"github.com/spf13/cobra"
)

func init() {
	listRootCmd.AddCommand(listKeyVaultKVContributorsCmd)
}

var listKeyVaultKVContributorsCmd = &cobra.Command{
	Use:          "key-vault-kvcontributors",
	Long:         "Lists Azure Key Vault KVContributors",
	Run:          listKeyVaultKVContributorsCmdImpl,
	SilenceUsage: true,
}

func listKeyVaultKVContributorsCmdImpl(cmd *cobra.Command, args []string) {
	ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, os.Kill)
	defer gracefulShutdown(stop)

	log.V(1).Info("testing connections")
	azClient := connectAndCreateClient()
	log.Info("collecting azure key vault kvcontributors...")
	start := time.Now()
	panicChan := panicChan()
	subscriptions := listSubscriptions(ctx, azClient, panicChan)
	keyVaults := listKeyVaults(ctx, azClient, panicChan, subscriptions)
	kvRoleAssignments := listKeyVaultRoleAssignments(ctx, azClient, panicChan, keyVaults)
	handleBubbledPanic(ctx, panicChan, stop)
	stream := listKeyVaultKVContributors(ctx, kvRoleAssignments)
	outputStream(ctx, stream)
	duration := time.Since(start)
	log.Info("collection completed", "duration", duration.String())
}

func listKeyVaultKVContributors(
	ctx context.Context,
	kvRoleAssignments <-chan azureWrapper[models.KeyVaultRoleAssignments],
) <-chan any {
	return pipeline.Map(ctx.Done(), kvRoleAssignments, func(ra azureWrapper[models.KeyVaultRoleAssignments]) any {
		filteredAssignments := internal.Filter(ra.Data.RoleAssignments, kvRoleAssignmentFilter(constants.KeyVaultContributorRoleID))

		kvContributors := internal.Map(filteredAssignments, func(ra models.KeyVaultRoleAssignment) models.KeyVaultKVContributor {
			return models.KeyVaultKVContributor{
				KVContributor: ra.RoleAssignment,
				KeyVaultId:    ra.KeyVaultId,
			}
		})

		return NewAzureWrapper(enums.KindAZKeyVaultKVContributor, models.KeyVaultKVContributors{
			KeyVaultId:     ra.Data.KeyVaultId,
			KVContributors: kvContributors,
		})
	})
}
