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
	"sync"
	"time"

	"github.com/bloodhoundad/azurehound/v2/client"
	"github.com/bloodhoundad/azurehound/v2/enums"
	"github.com/bloodhoundad/azurehound/v2/models"
	"github.com/bloodhoundad/azurehound/v2/pipeline"
	"github.com/spf13/cobra"
)

func init() {
	listRootCmd.AddCommand(listStorageAccountsCmd)
}

var listStorageAccountsCmd = &cobra.Command{
	Use:          "storage-accounts",
	Long:         "Lists Azure Storage Accounts",
	Run:          listStorageAccountsCmdImpl,
	SilenceUsage: true,
}

func listStorageAccountsCmdImpl(cmd *cobra.Command, args []string) {
	ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, os.Kill)
	defer gracefulShutdown(stop)

	log.V(1).Info("testing connections")
	azClient := connectAndCreateClient()
	log.Info("collecting azure storage accounts...")
	start := time.Now()
	panicChan := panicChan()
	stream := listStorageAccounts(ctx, azClient, panicChan, listSubscriptions(ctx, azClient, panicChan))
	handleBubbledPanic(ctx, panicChan, stop)
	outputStream(ctx, stream)
	duration := time.Since(start)
	log.Info("collection completed", "duration", duration.String())
}

func listStorageAccounts(ctx context.Context, client client.AzureClient, panicChan chan error, subscriptions <-chan interface{}) <-chan interface{} {
	var (
		out     = make(chan interface{})
		ids     = make(chan string)
		streams = pipeline.Demux(ctx.Done(), ids, 25)
		wg      sync.WaitGroup
	)

	go func() {
		defer panicRecovery(panicChan)
		defer close(ids)
		for result := range pipeline.OrDone(ctx.Done(), subscriptions) {
			if subscription, ok := result.(AzureWrapper).Data.(models.Subscription); !ok {
				log.Error(fmt.Errorf("failed type assertion"), "unable to continue enumerating storage accounts", "result", result)
				return
			} else {
				if ok := pipeline.Send(ctx.Done(), ids, subscription.SubscriptionId); !ok {
					return
				}
			}
		}
	}()

	wg.Add(len(streams))
	for i := range streams {
		stream := streams[i]
		go func() {
			defer panicRecovery(panicChan)
			defer wg.Done()
			for id := range stream {
				count := 0
				for item := range client.ListAzureStorageAccounts(ctx, id) {
					if item.Error != nil {
						log.Error(item.Error, "unable to continue processing storage accounts for this subscription", "subscriptionId", id)
					} else {
						resourceGroupId := item.Ok.ResourceGroupId()
						resourceGroupName := item.Ok.ResourceGroupName()
						storageAccount := models.StorageAccount{
							StorageAccount:    item.Ok,
							SubscriptionId:    item.SubscriptionId,
							ResourceGroupId:   resourceGroupId,
							ResourceGroupName: resourceGroupName,
							TenantId:          client.TenantInfo().TenantId,
						}
						log.V(2).Info("found storage account", "storageAccount", storageAccount)
						count++
						if ok := pipeline.SendAny(ctx.Done(), out, AzureWrapper{
							Kind: enums.KindAZStorageAccount,
							Data: storageAccount,
						}); !ok {
							return
						}
					}
				}
				log.V(1).Info("finished listing storage accounts", "subscriptionId", id, "count", count)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
		log.Info("finished listing all storage accounts")
	}()

	return out
}
