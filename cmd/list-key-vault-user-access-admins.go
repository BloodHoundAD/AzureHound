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
	listRootCmd.AddCommand(listKeyVaultUserAccessAdminsCmd)
}

var listKeyVaultUserAccessAdminsCmd = &cobra.Command{
	Use:          "key-vault-user-access-admins",
	Long:         "Lists Azure Key Vault User Access Admins",
	Run:          listKeyVaultUserAccessAdminsCmdImpl,
	SilenceUsage: true,
}

func listKeyVaultUserAccessAdminsCmdImpl(cmd *cobra.Command, args []string) {
	ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, os.Kill)
	defer gracefulShutdown(stop)

	log.V(1).Info("testing connections")
	if err := testConnections(); err != nil {
		exit(err)
	} else if azClient, err := newAzureClient(); err != nil {
		exit(err)
	} else {
		log.Info("collecting azure key vault user access admins...")
		start := time.Now()
		subscriptions := listSubscriptions(ctx, azClient)
		stream := listKeyVaultUserAccessAdmins(ctx, azClient, listKeyVaults(ctx, azClient, subscriptions))
		outputStream(ctx, stream)
		duration := time.Since(start)
		log.Info("collection completed", "duration", duration.String())
	}
}

func listKeyVaultUserAccessAdmins(ctx context.Context, client client.AzureClient, keyVaults <-chan interface{}) <-chan interface{} {
	var (
		out     = make(chan interface{})
		ids     = make(chan string)
		streams = pipeline.Demux(ctx.Done(), ids, 25)
		wg      sync.WaitGroup
	)

	go func() {
		defer close(ids)

		for result := range pipeline.OrDone(ctx.Done(), keyVaults) {
			if keyVault, ok := result.(AzureWrapper).Data.(models.KeyVault); !ok {
				log.Error(fmt.Errorf("failed type assertion"), "unable to continue enumerating key vault user access admins", "result", result)
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
					keyVaultUserAccessAdmins = models.KeyVaultUserAccessAdmins{
						KeyVaultId: id.(string),
					}
					count = 0
				)
				for item := range client.ListRoleAssignmentsForResource(ctx, id.(string), "") {
					if item.Error != nil {
						log.Error(item.Error, "unable to continue processing user access admins for this key vault", "keyVaultId", id)
					} else {
						roleDefinitionId := path.Base(item.Ok.Properties.RoleDefinitionId)

						if roleDefinitionId == constants.UserAccessAdminRoleID {
							keyVaultUserAccessAdmin := models.KeyVaultUserAccessAdmin{
								UserAccessAdmin: item.Ok,
								KeyVaultId:      item.ParentId,
							}
							log.V(2).Info("found key vault user access admin", "keyVaultUserAccessAdmin", keyVaultUserAccessAdmin)
							count++
							keyVaultUserAccessAdmins.UserAccessAdmins = append(keyVaultUserAccessAdmins.UserAccessAdmins, keyVaultUserAccessAdmin)
						}
					}
				}
				out <- AzureWrapper{
					Kind: enums.KindAZKeyVaultUserAccessAdmin,
					Data: keyVaultUserAccessAdmins,
				}
				log.V(1).Info("finished listing key vault user access admins", "keyVaultId", id, "count", count)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
		log.Info("finished listing all key vault user access admins")
	}()
	return out
}
