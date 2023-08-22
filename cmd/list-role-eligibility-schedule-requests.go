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
	listRootCmd.AddCommand(listRoleEligibilityScheduleRequestsCmd)
}

var listRoleEligibilityScheduleRequestsCmd = &cobra.Command{
	Use:          "role-eligibility-schedule-requests",
	Long:         "Lists Azure Active Directory Role Eligibility Requests",
	Run:          listRoleEligibilityScheduleRequestsCmdImpl,
	SilenceUsage: true,
}

func listRoleEligibilityScheduleRequestsCmdImpl(cmd *cobra.Command, args []string) {
	ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, os.Kill)
	defer gracefulShutdown(stop)

	log.V(1).Info("testing connections")
	azClient := connectAndCreateClient()
	log.Info("collecting azure active directory role eligibility requests...")
	start := time.Now()
	roles := listRoles(ctx, azClient)
	stream := listRoleEligibilityScheduleRequests(ctx, azClient, roles)
	outputStream(ctx, stream)
	duration := time.Since(start)
	log.Info("collection completed", "duration", duration.String())
}

func listRoleEligibilityScheduleRequests(ctx context.Context, client client.AzureClient, roles <-chan interface{}) <-chan interface{} {
	var (
		out     = make(chan interface{})
		ids     = make(chan string)
		streams = pipeline.Demux(ctx.Done(), ids, 25)
		wg      sync.WaitGroup
	)

	go func() {
		defer close(ids)

		for result := range pipeline.OrDone(ctx.Done(), roles) {
			if role, ok := result.(AzureWrapper).Data.(models.Role); !ok {
				log.Error(fmt.Errorf("failed type assertion"), "unable to continue enumerating role eligibility schedule requests", "result", result)
				return
			} else {
				ids <- role.Id
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
					roleEligibilityScheduleRequests = models.RoleEligibilityScheduleRequests{
						RoleDefinitionId: id,
						TenantId:         client.TenantInfo().TenantId,
					}
					count  = 0
					filter = fmt.Sprintf("roleDefinitionId eq '%s'", id)
				)
				for item := range client.ListAzureADRoleEligibilityScheduleRequests(ctx, filter, "", "", "", nil) {
					if item.Error != nil {
						log.Error(item.Error, "unable to continue processing role eligibility schedule requests for this role", "roleDefinitionId", id)
					} else {
						log.V(2).Info("found role eligibility schedule request", "roleEligibilityScheduleRequest", item)
						count++
						roleEligibilityScheduleRequests.RoleEligibilityScheduleRequests = append(roleEligibilityScheduleRequests.RoleEligibilityScheduleRequests, item.Ok)
					}
				}
				out <- AzureWrapper{
					Kind: enums.KindAZRoleEligibilityScheduleRequest,
					Data: roleEligibilityScheduleRequests,
				}
				log.V(1).Info("finished listing role eligibility schedule requests", "roleDefinitionId", id, "count", count)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
		log.Info("finished listing all role eligibility schedule requests")
	}()

	return out
}
