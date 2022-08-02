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

package config

import (
	"errors"
	"fmt"
	"net/url"
	"os"

	client "github.com/bloodhoundad/azurehound/client/config"
	config "github.com/bloodhoundad/azurehound/config/internal"
	"github.com/bloodhoundad/azurehound/constants"
)

var Init = config.Init
var LoadValues = config.LoadValues

func SetAzureDefaults() {
	if AzAuthUrl.Value() == "" {
		region := AzRegion.Value().(string)
		url := client.AuthorityUrl(region, constants.AzureCloud().ActiveDirectoryAuthority)
		AzAuthUrl.Set(url)
	}

	if AzGraphUrl.Value() == "" {
		region := AzRegion.Value().(string)
		url := client.GraphUrl(region, constants.AzureCloud().MicrosoftGraphUrl)
		AzGraphUrl.Set(url)
	}

	if AzMgmtUrl.Value() == "" {
		region := AzRegion.Value().(string)
		url := client.ResourceManagerUrl(region, constants.AzureCloud().ResourceManagerUrl)
		AzMgmtUrl.Set(url)
	}
}

func ValidateURL(input string) error {
	if parsedURL, err := url.Parse(input); err != nil {
		return err
	} else if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return fmt.Errorf("invalid URL")
	} else {
		return nil
	}
}

func SetProxyEnvVars() error {
	if proxy, ok := Proxy.Value().(string); !ok || proxy == "" {
		return nil
	} else if err := ValidateURL(proxy); err != nil {
		return err
	} else if proxyURL, err := url.Parse(proxy); err != nil {
		return err
	} else if proxyURL.Scheme == "https" {
		os.Setenv("HTTPS_PROXY", proxy)
		return nil
	} else if proxyURL.Scheme == "http" {
		os.Setenv("HTTP_PROXY", proxy)
		return nil
	} else {
		return errors.New("unsupported url scheme")
	}
}

func Options() config.Options {
	return config.Options{
		ConfigFile:  ConfigFile.Value().(string),
		ConfigName:  "config",
		ConfigPaths: SystemConfigDirs(),
		EnvPrefix:   EnvPrefix,
	}
}
