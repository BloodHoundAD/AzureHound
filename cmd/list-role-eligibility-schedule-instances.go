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
	listRootCmd.AddCommand(listRoleEligibilityScheduleInstancesCmd)
}

var listRoleEligibilityScheduleInstancesCmd = &cobra.Command{
	Use:          "role-eligibility-schedule-instances",
	Long:         "Lists Azure Active Directory Role Eligibility Instances",
	Run:          listRoleEligibilityScheduleInstancesCmdImpl,
	SilenceUsage: true,
}

func listRoleEligibilityScheduleInstancesCmdImpl(cmd *cobra.Command, args []string) {
	ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, os.Kill)
	defer gracefulShutdown(stop)

	log.V(1).Info("testing connections")
	azClient := connectAndCreateClient()
	log.Info("collecting azure active directory role eligibility instances...")
	start := time.Now()
	roles := listRoles(ctx, azClient)
	stream := listRoleEligibilityScheduleInstances(ctx, azClient, roles)
	outputStream(ctx, stream)
	duration := time.Since(start)
	log.Info("collection completed", "duration", duration.String())
}

func listRoleEligibilityScheduleInstances(ctx context.Context, client client.AzureClient, roles <-chan interface{}) <-chan interface{} {
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
				log.Error(fmt.Errorf("failed type assertion"), "unable to continue enumerating role eligibility schedule instances", "result", result)
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
					roleEligibilityScheduleInstances = models.RoleEligibilityScheduleInstances{
						RoleDefinitionId: id,
						TenantId:         client.TenantInfo().TenantId,
					}
					count  = 0
					filter = fmt.Sprintf("roleDefinitionId eq '%s'", id)
				)
				for item := range client.ListAzureADRoleEligibilityScheduleInstances(ctx, filter, "", "", "", nil) {
					if item.Error != nil {
						log.Error(item.Error, "unable to continue processing role eligibility schedule instances for this role", "roleDefinitionId", id)
					} else {
						log.V(2).Info("found role eligibility schedule instance", "roleEligibilityScheduleInstance", item)
						count++
						roleEligibilityScheduleInstances.RoleEligibilityScheduleInstances = append(roleEligibilityScheduleInstances.RoleEligibilityScheduleInstances, item.Ok)
					}
				}
				out <- AzureWrapper{
					Kind: enums.KindAZRoleEligibilityScheduleInstance,
					Data: roleEligibilityScheduleInstances,
				}
				log.V(1).Info("finished listing role eligibility schedule instances", "roleDefinitionId", id, "count", count)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
		log.Info("finished listing all role eligibility schedule instances")
	}()

	return out
}
