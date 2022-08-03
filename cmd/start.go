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
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"runtime"
	"sort"
	"time"

	"github.com/bloodhoundad/azurehound/client/rest"
	"github.com/bloodhoundad/azurehound/config"
	"github.com/bloodhoundad/azurehound/constants"
	"github.com/bloodhoundad/azurehound/models"
	"github.com/bloodhoundad/azurehound/pipeline"
	"github.com/spf13/cobra"
)

const (
	BHEAuthSignature string = "bhesignature"
)

func init() {
	configs := append(config.AzureConfig, config.BloodHoundEnterpriseConfig...)
	config.Init(startCmd, configs)
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:               "start",
	Short:             "Start Azure data collection",
	Run:               startCmdImpl,
	PersistentPreRunE: persistentPreRunE,
	SilenceUsage:      true,
}

func startCmdImpl(cmd *cobra.Command, args []string) {
	start(cmd.Context())
}

func start(ctx context.Context) {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	sigChan := make(chan os.Signal)
	go func() {
		stacktrace := make([]byte, 8192)
		for range sigChan {
			length := runtime.Stack(stacktrace, true)
			fmt.Println(string(stacktrace[:length]))
		}
	}()
	defer gracefulShutdown(stop)

	log.V(1).Info("testing connections")
	if err := testConnections(); err != nil {
		exit(err)
	} else if azClient, err := newAzureClient(); err != nil {
		exit(err)
	} else if bheInstance, err := url.Parse(config.BHEUrl.Value().(string)); err != nil {
		exit(err)
	} else if bheClient, err := newSigningHttpClient(BHEAuthSignature, config.BHETokenId.Value().(string), config.BHEToken.Value().(string), config.Proxy.Value().(string)); err != nil {
		exit(err)
	} else {

		if err := updateClient(ctx, *bheInstance, bheClient); err != nil {
			exit(err)
		}

		log.Info("connected successfully! waiting for tasks...")
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		var currentTask *models.ClientTask

		for {
			select {
			case <-ticker.C:
				if currentTask != nil {
					log.V(1).Info("curently performing collection; continuing...")
				} else {
					log.V(2).Info("checking for available collection tasks")
					if availableTasks, err := getAvailableTasks(ctx, *bheInstance, bheClient); err != nil {
						log.Error(err, "unable to fetch available tasks for azurehound")
					} else {

						// Get only the tasks that have reached their execution time
						executableTasks := []models.ClientTask{}
						now := time.Now()
						for _, task := range availableTasks {
							if task.ExectionTime.Before(now) || task.ExectionTime.Equal(now) {
								executableTasks = append(executableTasks, task)
							}
						}

						// Sort tasks in ascending order by execution time
						sort.Slice(executableTasks, func(i, j int) bool {
							return executableTasks[i].ExectionTime.Before(executableTasks[j].ExectionTime)
						})

						if len(executableTasks) == 0 {
							log.V(2).Info("there are no tasks for azurehound to complete at this time")
						} else {

							// Notify BHE instance of task start
							currentTask = &executableTasks[0]
							startTask(ctx, *bheInstance, bheClient, currentTask.Id)
							start := time.Now()

							// Batch data out for ingestion
							stream := listAll(ctx, azClient)
							batches := pipeline.Batch(ctx.Done(), stream, 999, 10*time.Second)
							ingest(ctx, *bheInstance, bheClient, batches)

							// Notify BHE instance of task end
							duration := time.Since(start)
							endTask(ctx, *bheInstance, bheClient)
							log.Info("finished collection task", "id", currentTask.Id, "duration", duration.String())

							currentTask = nil
						}
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}
}

func ingest(ctx context.Context, bheUrl url.URL, bheClient *http.Client, in <-chan []interface{}) {
	endpoint := bheUrl.ResolveReference(&url.URL{Path: "/api/v1/ingest"})

	for data := range pipeline.OrDone(ctx.Done(), in) {
		body := models.IngestRequest{
			Meta: models.Meta{
				Type: "azure",
			},
			Data: data,
		}

		if req, err := rest.NewRequest(ctx, "POST", endpoint, body, nil, nil); err != nil {
			log.Error(err, "unable to create request")
		} else if _, err := bheClient.Do(req); err != nil {
			log.Error(err, "unable to send data to bloodhound enterprise", "bheUrl", bheUrl)
		}
	}
}

func getAvailableTasks(ctx context.Context, bheUrl url.URL, bheClient *http.Client) ([]models.ClientTask, error) {
	var (
		endpoint = bheUrl.ResolveReference(&url.URL{Path: "/api/v1/clients/availabletasks"})
		response []models.ClientTask
	)

	if req, err := rest.NewRequest(ctx, "GET", endpoint, nil, nil, nil); err != nil {
		return nil, err
	} else if res, err := bheClient.Do(req); err != nil {
		return nil, err
	} else if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	} else {
		return response, nil
	}
}

func startTask(ctx context.Context, bheUrl url.URL, bheClient *http.Client, taskId int) error {
	log.Info("beginning collection task", "id", taskId)
	var (
		endpoint = bheUrl.ResolveReference(&url.URL{Path: "/api/v1/clients/starttask"})
		body     = map[string]int{
			"id": taskId,
		}
	)

	if req, err := rest.NewRequest(ctx, "POST", endpoint, body, nil, nil); err != nil {
		return err
	} else if _, err := bheClient.Do(req); err != nil {
		return err
	} else {
		return nil
	}
}

func endTask(ctx context.Context, bheUrl url.URL, bheClient *http.Client) error {
	endpoint := bheUrl.ResolveReference(&url.URL{Path: "/api/v1/clients/endtask"})

	if req, err := rest.NewRequest(ctx, "POST", endpoint, nil, nil, nil); err != nil {
		return err
	} else if _, err := bheClient.Do(req); err != nil {
		return err
	} else {
		return nil
	}
}

func updateClient(ctx context.Context, bheUrl url.URL, bheClient *http.Client) error {
	endpoint := bheUrl.ResolveReference(&url.URL{Path: "/api/v1/clients/update"})
	if addr, err := dial(bheUrl.String()); err != nil {
		return err
	} else {
		// hostname is nice to have but we don't really need it
		hostname, _ := os.Hostname()

		body := models.UpdateClientRequest{
			Address:  addr,
			Hostname: hostname,
			Version:  constants.Version,
		}

		log.V(2).Info("updating client info", "info", body)

		if req, err := rest.NewRequest(ctx, "PUT", endpoint, body, nil, nil); err != nil {
			return err
		} else if res, err := bheClient.Do(req); err != nil {
			return err
		} else {
			var body json.RawMessage
			defer res.Body.Close()
			if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
				return err
			} else if res.StatusCode < 200 || res.StatusCode >= 400 {
				return fmt.Errorf(string(body))
			} else {
				return nil
			}
		}
	}
}
