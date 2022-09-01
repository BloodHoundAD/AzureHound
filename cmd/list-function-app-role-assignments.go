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
	listRootCmd.AddCommand(listFunctionAppRoleAssignment)
}

var listFunctionAppRoleAssignment = &cobra.Command{
	Use:          "function-app-role-assignments",
	Long:         "Lists Azure Function App Role Assignments",
	Run:          listFunctionAppRoleAssignmentImpl,
	SilenceUsage: true,
}

func listFunctionAppRoleAssignmentImpl(cmd *cobra.Command, args []string) {
	ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, os.Kill)
	defer gracefulShutdown(stop)

	log.V(1).Info("testing connections")
	if err := testConnections(); err != nil {
		exit(err)
	} else if azClient, err := newAzureClient(); err != nil {
		exit(err)
	} else {
		log.Info("collecting azure function app role assignments...")
		start := time.Now()
		subscriptions := listSubscriptions(ctx, azClient)
		stream := listFunctionAppRoleAssignments(ctx, azClient, listFunctionApps(ctx, azClient, subscriptions))
		outputStream(ctx, stream)
		duration := time.Since(start)
		log.Info("collection completed", "duration", duration.String())
	}
}

func listFunctionAppRoleAssignments(ctx context.Context, client client.AzureClient, functionApps <-chan interface{}) <-chan interface{} {
	var (
		out     = make(chan interface{})
		ids     = make(chan string)
		streams = pipeline.Demux(ctx.Done(), ids, 25)
		wg      sync.WaitGroup
	)

	go func() {
		defer close(ids)

		for result := range pipeline.OrDone(ctx.Done(), functionApps) {
			if functionApp, ok := result.(AzureWrapper).Data.(models.FunctionApp); !ok {
				log.Error(fmt.Errorf("failed type assertion"), "unable to continue enumerating function app role assignments", "result", result)
				return
			} else {
				ids <- functionApp.Id
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
					functionAppOwners = models.FunctionAppOwners{
						FunctionAppId: id.(string),
					}
					functionAppContributors = models.FunctionAppContributors{
						FunctionAppId: id.(string),
					}
					count = 0
				)
				for item := range client.ListRoleAssignmentsForResource(ctx, id.(string), "") {
					if item.Error != nil {
						log.Error(item.Error, "unable to continue processing role assignments for this function app", "functionAppId", id)
					} else {
						roleDefinitionId := path.Base(item.Ok.Properties.RoleDefinitionId)

						if roleDefinitionId == constants.OwnerRoleID {
							functionAppOwner := models.FunctionAppOwner{
								Owner:         item.Ok,
								FunctionAppId: item.ParentId,
							}
							log.V(2).Info("found function app owner", "functionAppOwner", functionAppOwner)
							count++
							functionAppOwners.Owners = append(functionAppOwners.Owners, functionAppOwner)
						} else if (roleDefinitionId == constants.ContributorRoleID) ||
							(roleDefinitionId == constants.AzWebsiteContributor) {
							functionAppContributor := models.FunctionAppContributor{
								Contributor:   item.Ok,
								FunctionAppId: item.ParentId,
							}
							log.V(2).Info("found function app contributor", "functionAppContributor", functionAppContributor)
							count++
							functionAppContributors.Contributors = append(functionAppContributors.Contributors, functionAppContributor)
						}
					}
				}
				out <- []AzureWrapper{
					{
						Kind: enums.KindAZFunctionAppOwner,
						Data: functionAppOwners,
					},
					{
						Kind: enums.KindAZFunctionAppContributor,
						Data: functionAppContributors,
					},
				}
				log.V(1).Info("finished listing function app owners", "functionAppId", id, "count", count)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
		log.Info("finished listing all function app owners")
	}()

	return out
}
