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
	listRootCmd.AddCommand(listStoraceAccountDataReader)
}

var listStoraceAccountDataReader = &cobra.Command{
	Use:          "storage-account-data-readers",
	Long:         "Lists Azure Storage Account DataReaders",
	Run:          listStoraceAccountDataReaderImpl,
	SilenceUsage: true,
}

func listStoraceAccountDataReaderImpl(cmd *cobra.Command, args []string) {
	ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, os.Kill)
	defer gracefulShutdown(stop)

	log.V(1).Info("testing connections")
	if err := testConnections(); err != nil {
		exit(err)
	} else if azClient, err := newAzureClient(); err != nil {
		exit(err)
	} else {
		log.Info("collecting azure storage account data-readers...")
		start := time.Now()
		subscriptions := listSubscriptions(ctx, azClient)
		stream := listStorageAccountDataReaders(ctx, azClient, listStorageAccounts(ctx, azClient, subscriptions))
		outputStream(ctx, stream)
		duration := time.Since(start)
		log.Info("collection completed", "duration", duration.String())
	}
}

func listStorageAccountDataReaders(ctx context.Context, client client.AzureClient, storageAccounts <-chan interface{}) <-chan interface{} {
	var (
		out     = make(chan interface{})
		ids     = make(chan string)
		streams = pipeline.Demux(ctx.Done(), ids, 25)
		wg      sync.WaitGroup
	)

	go func() {
		defer close(ids)

		for result := range pipeline.OrDone(ctx.Done(), storageAccounts) {
			if storageAccount, ok := result.(AzureWrapper).Data.(models.StorageAccount); !ok {
				log.Error(fmt.Errorf("failed type assertion"), "unable to continue enumerating storage account data-readers", "result", result)
				return
			} else {
				ids <- storageAccount.Id
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
					storageAccountDataReaders = models.StorageAccountDataReaders{
						StorageAccountId: id.(string),
					}
					count = 0
				)
				for item := range client.ListRoleAssignmentsForResource(ctx, id.(string), "") {
					if item.Error != nil {
						log.Error(item.Error, "unable to continue processing data-readers for this storage account", "storageAccountId", id)
					} else {
						roleDefinitionId := path.Base(item.Ok.Properties.RoleDefinitionId)

						if roleDefinitionId == constants.DataReaderRoleID {
							storageAccountDataReader := models.StorageAccountDataReader{
								DataReader:       item.Ok,
								StorageAccountId: item.ParentId,
							}
							log.V(2).Info("found storage account data-reader", "storageAccountDataReader", storageAccountDataReader)
							count++
							storageAccountDataReaders.DataReaders = append(storageAccountDataReaders.DataReaders, storageAccountDataReader)
						}
					}
				}
				out <- AzureWrapper{
					Kind: enums.KindAZSADataReader,
					Data: storageAccountDataReaders,
				}
				log.V(1).Info("finished listing storage account data-readers", "storageAccountId", id, "count", count)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
		log.Info("finished listing all storage account data-readers")
	}()

	return out
}