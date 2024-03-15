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

	"github.com/bloodhoundad/azurehound/v2/client"
	"github.com/bloodhoundad/azurehound/v2/enums"
	"github.com/bloodhoundad/azurehound/v2/models"
	"github.com/bloodhoundad/azurehound/v2/pipeline"
	"github.com/spf13/cobra"
)

func init() {
	listRootCmd.AddCommand(listUsersCmd)
}

var listUsersCmd = &cobra.Command{
	Use:          "users",
	Long:         "Lists Azure Active Directory Users",
	Run:          listUsersCmdImpl,
	SilenceUsage: true,
}

func listUsersCmdImpl(cmd *cobra.Command, args []string) {
	ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, os.Kill)
	defer gracefulShutdown(stop)

	log.V(1).Info("testing connections")
	azClient := connectAndCreateClient()
	log.Info("collecting azure active directory users...")
	start := time.Now()
	panicChan := panicChan()
	stream := listUsers(ctx, azClient, panicChan)
	handleBubbledPanic(ctx, panicChan, stop)
	outputStream(ctx, stream)
	duration := time.Since(start)
	log.Info("collection completed", "duration", duration.String())
}

func listUsers(ctx context.Context, client client.AzureClient, panicChan chan error) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer panicRecovery(panicChan)
		defer close(out)
		count := 0
		for item := range client.ListAzureADUsers(ctx, "", "", "", []string{
			"accountEnabled",
			"createdDateTime",
			"displayName",
			"jobTitle",
			"lastPasswordChangeDateTime",
			"mail",
			"onPremisesSecurityIdentifier",
			"onPremisesSyncEnabled",
			"userPrincipalName",
			"userType",
			"id",
		}) {
			if item.Error != nil {
				log.Error(item.Error, "unable to continue processing users")
				return
			} else {
				log.V(2).Info("found user", "user", item)
				count++
				user := models.User{
					User:       item.Ok,
					TenantId:   client.TenantInfo().TenantId,
					TenantName: client.TenantInfo().DisplayName,
				}
				if ok := pipeline.SendAny(ctx.Done(), out, AzureWrapper{
					Kind: enums.KindAZUser,
					Data: user,
				}); !ok {
					return
				}
			}
		}
		log.Info("finished listing all users", "count", count)
	}()

	return out
}
