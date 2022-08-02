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
	"os"
	"os/signal"
	"time"

	"github.com/bloodhoundad/azurehound/client"
	"github.com/bloodhoundad/azurehound/enums"
	"github.com/bloodhoundad/azurehound/models"
	"github.com/spf13/cobra"
)

func init() {
	listRootCmd.AddCommand(listRolesCmd)
}

var listRolesCmd = &cobra.Command{
	Use:          "roles",
	Long:         "Lists Azure Active Directory Roles",
	Run:          listRolesCmdImpl,
	SilenceUsage: true,
}

func listRolesCmdImpl(cmd *cobra.Command, args []string) {
	ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, os.Kill)
	defer gracefulShutdown(stop)

	log.V(1).Info("testing connections")
	if err := testConnections(); err != nil {
		exit(err)
	} else if azClient, err := newAzureClient(); err != nil {
		exit(err)
	} else {
		log.Info("collecting azure active directory roles...")
		start := time.Now()
		stream := listRoles(ctx, azClient)
		outputStream(ctx, stream)
		duration := time.Since(start)
		log.Info("collection completed", "duration", duration.String())
	}
}

func listRoles(ctx context.Context, client client.AzureClient) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)
		count := 0
		for item := range client.ListAzureADRoles(ctx, "", "") {
			if item.Error != nil {
				log.Error(item.Error, "unable to continue processing roles")
				return
			} else {
				log.V(2).Info("found role", "role", item)
				count++
				out <- AzureWrapper{
					Kind: enums.KindAZRole,
					Data: models.Role{
						Role:       item.Ok,
						TenantId:   client.TenantInfo().TenantId,
						TenantName: client.TenantInfo().DisplayName,
					},
				}
			}
		}
		log.Info("finished listing all roles", "count", count)
	}()

	return out
}
