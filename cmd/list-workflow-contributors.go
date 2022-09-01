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
	listRootCmd.AddCommand(listWorkflowContributor)
}

var listWorkflowContributor = &cobra.Command{
	Use:          "workflow-contributors",
	Long:         "Lists Azure Workflow (Logic Apps) Contributors",
	Run:          listWorkflowContributorImpl,
	SilenceUsage: true,
}

func listWorkflowContributorImpl(cmd *cobra.Command, args []string) {
	ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, os.Kill)
	defer gracefulShutdown(stop)

	log.V(1).Info("testing connections")
	if err := testConnections(); err != nil {
		exit(err)
	} else if azClient, err := newAzureClient(); err != nil {
		exit(err)
	} else {
		log.Info("collecting azure workflow contributors...")
		start := time.Now()
		subscriptions := listSubscriptions(ctx, azClient)
		stream := listWorkflowContributors(ctx, azClient, listWorkflows(ctx, azClient, subscriptions))
		outputStream(ctx, stream)
		duration := time.Since(start)
		log.Info("collection completed", "duration", duration.String())
	}
}

func listWorkflowContributors(ctx context.Context, client client.AzureClient, workflows <-chan interface{}) <-chan interface{} {
	var (
		out     = make(chan interface{})
		ids     = make(chan string)
		streams = pipeline.Demux(ctx.Done(), ids, 25)
		wg      sync.WaitGroup
	)

	go func() {
		defer close(ids)

		for result := range pipeline.OrDone(ctx.Done(), workflows) {
			if workflow, ok := result.(AzureWrapper).Data.(models.Workflow); !ok {
				log.Error(fmt.Errorf("failed type assertion"), "unable to continue enumerating workflow contributors", "result", result)
				return
			} else {
				ids <- workflow.Id
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
					workflowContributors = models.WorkflowContributors{
						WorkflowId: id.(string),
					}
					count = 0
				)
				for item := range client.ListRoleAssignmentsForResource(ctx, id.(string), "") {
					if item.Error != nil {
						log.Error(item.Error, "unable to continue processing contributors for this workflow", "workflowId", id)
					} else {
						roleDefinitionId := path.Base(item.Ok.Properties.RoleDefinitionId)

						if (roleDefinitionId == constants.ContributorRoleID) ||
							(roleDefinitionId == constants.AzLogicAppContributor) {
							workflowContributor := models.WorkflowContributor{
								Contributor: item.Ok,
								WorkflowId:  item.ParentId,
							}
							log.V(2).Info("found workflow contributor", "workflowContributor", workflowContributor)
							count++
							workflowContributors.Contributors = append(workflowContributors.Contributors, workflowContributor)
						}
					}
				}
				out <- AzureWrapper{
					Kind: enums.KindAZWorkflowContributor,
					Data: workflowContributors,
				}
				log.V(1).Info("finished listing workflow contributors", "workflowId", id, "count", count)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
		log.Info("finished listing all workflow contributors")
	}()

	return out
}
