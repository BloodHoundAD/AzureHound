// Copyright (C) 2022 The BloodHound Enterprise Team
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
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/bloodhoundad/azurehound/client"
	client_config "github.com/bloodhoundad/azurehound/client/config"
	"github.com/bloodhoundad/azurehound/client/rest"
	"github.com/bloodhoundad/azurehound/config"
	"github.com/bloodhoundad/azurehound/enums"
	"github.com/bloodhoundad/azurehound/logger"
	"github.com/bloodhoundad/azurehound/pipeline"
	"github.com/bloodhoundad/azurehound/sinks"
	"github.com/spf13/cobra"
)

func exit(err error) {
	log.Error(err, "encountered unrecoverable error")
	log.GetSink()
	os.Exit(1)
}

func persistentPreRunE(cmd *cobra.Command, args []string) error {
	// need to set config flag value explicitly
	if cmd != nil {
		if configFlag := cmd.Flag(config.ConfigFile.Name).Value.String(); configFlag != "" {
			config.ConfigFile.Set(configFlag)
		}
	}

	if logr, err := logger.GetLogger(); err != nil {
		return err
	} else {
		log = *logr

		config.LoadValues(cmd, config.Options())
		config.SetAzureDefaults()

		if err := config.SetProxyEnvVars(); err != nil {
			return err
		}

		if config.ConfigFileUsed() != "" {
			log.V(1).Info(fmt.Sprintf("Config File: %v", config.ConfigFileUsed()))
		}

		if config.LogFile.Value() != "" {
			log.V(1).Info(fmt.Sprintf("Log File: %v", config.LogFile.Value()))
		}

		return nil
	}
}

func gracefulShutdown(stop context.CancelFunc) {
	stop()
	fmt.Fprintln(os.Stderr, "\nshutting down gracefully, press ctrl+c again to force")
	// TODO timeout context
}

func testConnections() error {
	if _, err := dial(config.AzAuthUrl.Value().(string)); err != nil {
		return fmt.Errorf("unable to connect to %s: %w", config.AzAuthUrl.Value(), err)
	} else if _, err := dial(config.AzGraphUrl.Value().(string)); err != nil {
		return fmt.Errorf("unable to connect to %s: %w", config.AzGraphUrl.Value(), err)
	} else if _, err := dial(config.AzMgmtUrl.Value().(string)); err != nil {
		return fmt.Errorf("unable to connect to %s: %w", config.AzMgmtUrl.Value(), err)
	} else {
		return nil
	}
}

func dial(targetUrl string) (string, error) {
	log.V(2).Info("dialing...", "targetUrl", targetUrl)
	if url, err := url.Parse(targetUrl); err != nil {
		return "", err
	} else if conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:https", url.Host), 5*time.Second); err != nil {
		return "", err
	} else {
		defer conn.Close()
		addr := conn.LocalAddr().(*net.TCPAddr)
		return addr.IP.String(), nil
	}
}

func newAzureClient() (client.AzureClient, error) {
	var (
		certFile   = config.AzCert.Value()
		keyFile    = config.AzKey.Value()
		clientCert string
		clientKey  string
	)

	if file, ok := certFile.(string); ok && file != "" {
		if content, err := ioutil.ReadFile(certFile.(string)); err != nil {
			return nil, fmt.Errorf("unable to read provided certificate: %w", err)
		} else {
			clientCert = string(content)
		}

	}

	if file, ok := keyFile.(string); ok && file != "" {
		if content, err := ioutil.ReadFile(keyFile.(string)); err != nil {
			return nil, fmt.Errorf("unable to read provided key file: %w", err)
		} else {
			clientKey = string(content)
		}
	}

	config := client_config.Config{
		ApplicationId:  config.AzAppId.Value().(string),
		Authority:      config.AzAuthUrl.Value().(string),
		ClientSecret:   config.AzSecret.Value().(string),
		ClientCert:     clientCert,
		ClientKey:      clientKey,
		ClientKeyPass:  config.AzKeyPass.Value().(string),
		Graph:          config.AzGraphUrl.Value().(string),
		JWT:            config.JWT.Value().(string),
		Management:     config.AzMgmtUrl.Value().(string),
		MgmtGroupId:    config.AzMgmtGroupId.Value().([]string),
		Password:       config.AzPassword.Value().(string),
		RefreshToken:   config.RefreshToken.Value().(string),
		Region:         config.AzRegion.Value().(string),
		SubscriptionId: config.AzSubId.Value().([]string),
		Tenant:         config.AzTenant.Value().(string),
		Username:       config.AzUsername.Value().(string),
	}
	return client.NewClient(config)
}

func newSigningHttpClient(signature, tokenId, token string) *http.Client {
	client := rest.NewHTTPClient()
	client.Transport = signingTransport{
		base:      client.Transport,
		tokenId:   tokenId,
		token:     token,
		signature: signature,
	}
	return client
}

type signingTransport struct {
	base      http.RoundTripper
	tokenId   string
	token     string
	signature string
}

func (s signingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	clone := req.Clone(req.Context())

	// token
	digester := hmac.New(sha256.New, []byte(s.token))

	// path
	if _, err := digester.Write([]byte(req.Method + req.URL.Path)); err != nil {
		return nil, err
	}

	// datetime
	datetime := time.Now().Format(time.RFC3339)
	digester = hmac.New(sha256.New, digester.Sum(nil))
	if _, err := digester.Write([]byte(datetime[:13])); err != nil {
		return nil, err
	}

	// body
	body := &bytes.Buffer{}
	digester = hmac.New(sha256.New, digester.Sum(nil))
	if req.Body != nil {
		defer req.Body.Close()
		if contentLength, err := body.ReadFrom(req.Body); err != nil {
			return nil, err
		} else if contentLength != 0 {
			req.Body = ioutil.NopCloser(bytes.NewReader(body.Bytes()))
			clone.Body = ioutil.NopCloser(bytes.NewReader(body.Bytes()))
		}
	}
	if _, err := digester.Write(body.Bytes()); err != nil {
		return nil, err
	}

	signature := digester.Sum(nil)

	clone.Header.Set("Authorization", fmt.Sprintf("%s %s", s.signature, s.tokenId))
	clone.Header.Set("RequestDate", datetime)
	clone.Header.Set("Signature", base64.StdEncoding.EncodeToString(signature))

	return s.base.RoundTrip(clone)
}

func contains(collection []string, value string) bool {
	for _, item := range collection {
		if item == value {
			return true
		}
	}
	return false
}

func unique(collection []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, item := range collection {
		if _, found := keys[item]; !found {
			keys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func stat(path string) (string, fs.FileInfo, error) {
	if info, err := os.Stat(path); err == nil {
		return path, info, nil
	} else {
		p := path + ".exe"
		info, err := os.Stat(p)
		return p, info, err
	}
}

func getExePath() (string, error) {
	exe := os.Args[0]
	if exePath, err := filepath.Abs(exe); err != nil {
		return "", err
	} else if path, info, err := stat(exePath); err != nil {
		return "", err
	} else if info.IsDir() {
		return "", fmt.Errorf("%s is a directory", path)
	} else {
		return path, nil
	}
}

func setupLogger() {
	if logger, err := logger.GetLogger(); err != nil {
		panic(err)
	} else {
		log = *logger
	}
}

type AzureWrapper struct {
	Kind enums.Kind  `json:"kind"`
	Data interface{} `json:"data"`
}

func outputStream(ctx context.Context, stream <-chan interface{}) {
	formatted := pipeline.FormatJson(ctx.Done(), stream)
	if path := config.OutputFile.Value().(string); path != "" {
		if err := sinks.WriteToFile(ctx, path, formatted); err != nil {
			exit(err)
		}
	} else {
		sinks.WriteToConsole(ctx, formatted)
	}
}
