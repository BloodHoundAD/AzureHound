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
	"path"
	"sync"
	"time"

	"github.com/bloodhoundad/azurehound/client"
	"github.com/bloodhoundad/azurehound/constants"
	"github.com/bloodhoundad/azurehound/enums"
	"github.com/bloodhoundad/azurehound/models"
	"github.com/bloodhoundad/azurehound/pipeline"
	"github.com/spf13/cobra"
)

func init() {
	listRootCmd.AddCommand(listKeyVaultOwnersCmd)
}

var listKeyVaultOwnersCmd = &cobra.Command{
	Use:          "key-vault-owners",
	Long:         "Lists Azure Key Vault Owners",
	Run:          listKeyVaultOwnersCmdImpl,
	SilenceUsage: true,
}

func listKeyVaultOwnersCmdImpl(cmd *cobra.Command, args []string) {
	ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, os.Kill)
	defer gracefulShutdown(stop)

	log.V(1).Info("testing connections")
	if err := testConnections(); err != nil {
		exit(err)
	} else if azClient, err := newAzureClient(); err != nil {
		exit(err)
	} else {
		log.Info("collecting azure key vault owners...")
		start := time.Now()
		subscriptions := listSubscriptions(ctx, azClient)
		stream := listKeyVaultOwners(ctx, azClient, listKeyVaults(ctx, azClient, subscriptions))
		outputStream(ctx, stream)
		duration := time.Since(start)
		log.Info("collection completed", "duration", duration.String())
	}
}

func listKeyVaultOwners(ctx context.Context, client client.AzureClient, keyVaults <-chan interface{}) <-chan interface{} {
	var (
		out     = make(chan interface{})
		ids     = make(chan string)
		streams = pipeline.Demux(ctx.Done(), ids, 50)
		wg      sync.WaitGroup
	)

	go func() {
		defer close(ids)

		for result := range pipeline.OrDone(ctx.Done(), keyVaults) {
			if keyVault, ok := result.(AzureWrapper).Data.(models.KeyVault); !ok {
				log.Error(fmt.Errorf("failed KeyVault type assertion"), "unable to continue enumerating key vault owners", "result", result)
				return
			} else {
				ids <- keyVault.Id
			}
		}
	}()

	wg.Add(len(streams))
	for i := range streams {
		stream := streams[i]
		go func() {
			defer wg.Done()
			for id := range stream {
				var (
					keyVaultOwners = models.KeyVaultOwners{
						KeyVaultId: id.(string),
					}
					count = 0
				)
				for item := range client.ListRoleAssignmentsForResource(ctx, id.(string), "") {
					if item.Error != nil {
						log.Error(item.Error, "unable to continue processing owners for this key vault", "keyVaultId", id)
					} else {
						roleDefinitionId := path.Base(item.Ok.Properties.RoleDefinitionId)

						if roleDefinitionId == constants.OwnerRoleID {
							keyVaultOwner := models.KeyVaultOwner{
								Owner:      item.Ok,
								KeyVaultId: item.ParentId,
							}
							log.V(2).Info("found key vault owner", "keyVaultOwner", keyVaultOwner)
							count++
							keyVaultOwners.Owners = append(keyVaultOwners.Owners, keyVaultOwner)
						}
					}
				}
				out <- AzureWrapper{
					Kind: enums.KindAZKeyVaultOwner,
					Data: keyVaultOwners,
				}
				log.V(1).Info("finished listing key vault owners", "keyVaultId", id, "count", count)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
		log.Info("finished listing all key vault owners")
	}()

	return out
}
