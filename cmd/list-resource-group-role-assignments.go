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

	"github.com/bloodhoundad/azurehound/client"
	"github.com/bloodhoundad/azurehound/enums"
	"github.com/bloodhoundad/azurehound/models"
	"github.com/bloodhoundad/azurehound/pipeline"
	"github.com/spf13/cobra"
)

func init() {
	listRootCmd.AddCommand(listResourceGroupRoleAssignmentsCmd)
}

var listResourceGroupRoleAssignmentsCmd = &cobra.Command{
	Use:          "resource-group-role-assignments",
	Long:         "Lists Resource Group Role Assignments",
	Run:          listResourceGroupRoleAssignmentsCmdImpl,
	SilenceUsage: true,
}

func listResourceGroupRoleAssignmentsCmdImpl(cmd *cobra.Command, args []string) {
	ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, os.Kill)
	defer gracefulShutdown(stop)

	log.V(1).Info("testing connections")
	if err := testConnections(); err != nil {
		exit(err)
	} else if azClient, err := newAzureClient(); err != nil {
		exit(err)
	} else {
		log.Info("collecting azure resource group role assignments...")
		start := time.Now()
		subscriptions := listSubscriptions(ctx, azClient)
		resourceGroups := listResourceGroups(ctx, azClient, subscriptions)
		stream := listResourceGroupRoleAssignments(ctx, azClient, resourceGroups)
		outputStream(ctx, stream)
		duration := time.Since(start)
		log.Info("collection completed", "duration", duration.String())
	}
}

func listResourceGroupRoleAssignments(ctx context.Context, client client.AzureClient, resourceGroups <-chan interface{}) <-chan azureWrapper[models.ResourceGroupRoleAssignments] {
	var (
		out     = make(chan azureWrapper[models.ResourceGroupRoleAssignments])
		ids     = make(chan string)
		streams = pipeline.Demux(ctx.Done(), ids, 25)
		wg      sync.WaitGroup
	)

	go func() {
		defer close(ids)

		for result := range pipeline.OrDone(ctx.Done(), resourceGroups) {
			if resourceGroup, ok := result.(AzureWrapper).Data.(models.ResourceGroup); !ok {
				log.Error(fmt.Errorf("failed type assertion"), "unable to continue enumerating resource group role assignments", "result", result)
				return
			} else {
				ids <- resourceGroup.Id
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
					resourceGroupRoleAssignments = models.ResourceGroupRoleAssignments{
						ResourceGroupId: id,
					}
					count = 0
				)
				for item := range client.ListRoleAssignmentsForResource(ctx, id, "") {
					if item.Error != nil {
						log.Error(item.Error, "unable to continue processing role assignments for this resourceGroup", "resourceGroupId", id)
					} else {
						resourceGroupRoleAssignment := models.ResourceGroupRoleAssignment{
							ResourceGroupId: item.ParentId,
							RoleAssignment:  item.Ok,
						}
						log.V(2).Info("found resourceGroup role assignment", "resourceGroupRoleAssignment", resourceGroupRoleAssignment)
						count++
						resourceGroupRoleAssignments.RoleAssignments = append(resourceGroupRoleAssignments.RoleAssignments, resourceGroupRoleAssignment)
					}
				}
				out <- NewAzureWrapper(enums.KindAZResourceGroupRoleAssignment, resourceGroupRoleAssignments)
				log.V(1).Info("finished listing resourceGroup role assignments", "resourceGroupId", id, "count", count)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
		log.Info("finished listing all resource group role assignments")
	}()

	return out
}
