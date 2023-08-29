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
	listRootCmd.AddCommand(listGroupEligibilityScheduleInstancesCmd)
}

var listGroupEligibilityScheduleInstancesCmd = &cobra.Command{
	Use:          "group-eligibility-schedule-instances",
	Long:         "Lists Azure Active Directory Group Eligibility Instances",
	Run:          listGroupEligibilityScheduleInstancesCmdImpl,
	SilenceUsage: true,
}

func listGroupEligibilityScheduleInstancesCmdImpl(cmd *cobra.Command, args []string) {
	ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, os.Kill)
	defer gracefulShutdown(stop)

	log.V(1).Info("testing connections")
	azClient := connectAndCreateClient()
	log.Info("collecting azure active directory group eligibility instances...")
	start := time.Now()
	groups := listGroups(ctx, azClient)
	stream := listGroupEligibilityScheduleInstances(ctx, azClient, groups)
	outputStream(ctx, stream)
	duration := time.Since(start)
	log.Info("collection completed", "duration", duration.String())
}

func listGroupEligibilityScheduleInstances(ctx context.Context, client client.AzureClient, groups <-chan interface{}) <-chan interface{} {
	var (
		out     = make(chan interface{})
		ids     = make(chan string)
		streams = pipeline.Demux(ctx.Done(), ids, 25)
		wg      sync.WaitGroup
	)

	go func() {
		defer close(ids)

		for result := range pipeline.OrDone(ctx.Done(), groups) {
			if group, ok := result.(AzureWrapper).Data.(models.Group); !ok {
				log.Error(fmt.Errorf("failed type assertion"), "unable to continue enumerating group eligibility schedule instances", "result", result)
				return
			} else {
				ids <- group.Id
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
					groupEligibilityScheduleInstances = models.GroupEligibilityScheduleInstances{
						GroupId:  id,
						TenantId: client.TenantInfo().TenantId,
					}
					count  = 0
					filter = fmt.Sprintf("groupId eq '%s'", id)
				)
				for item := range client.ListAzureADGroupEligibilityScheduleInstances(ctx, filter, "", "", "", nil) {
					if item.Error != nil {
						log.Error(item.Error, "unable to continue processing group eligibility schedule instances for this group", "groupId", id)
					} else {
						log.V(2).Info("found group eligibility schedule instance", "groupEligibilityScheduleInstance", item)
						count++
						groupEligibilityScheduleInstances.GroupEligibilityScheduleInstances = append(groupEligibilityScheduleInstances.GroupEligibilityScheduleInstances, item.Ok)
					}
				}
				out <- AzureWrapper{
					Kind: enums.KindAZGroupEligibilityScheduleInstance,
					Data: groupEligibilityScheduleInstances,
				}
				log.V(1).Info("finished listing group eligibility schedule instances", "groupId", id, "count", count)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
		log.Info("finished listing all group eligibility schedule instances")
	}()

	return out
}
