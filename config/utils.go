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

package config

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
	"syscall"

	client "github.com/bloodhoundad/azurehound/v2/client/config"
	config "github.com/bloodhoundad/azurehound/v2/config/internal"
	"github.com/bloodhoundad/azurehound/v2/constants"
	"golang.org/x/term"
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

func ReadFromStdInput() error {
	if RefreshToken.Value() == "-" {
		fmt.Print("Enter Refresh Token: ")
		rToken, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return err
		}
		RefreshToken.Set(strings.TrimSpace(string(rToken)))
	}

	if JWT.Value() == "-" {
		fmt.Print("Enter JWT: ")
		jwt, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return err
		}
		JWT.Set((strings.TrimSpace(string(jwt))))
	}

	if AzSecret.Value() == "-" {
		fmt.Print("Enter Application Secret: ")
		azSecret, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return err
		}
		AzSecret.Set((strings.TrimSpace(string(azSecret))))
	}

	if AzKeyPass.Value() == "-" {
		fmt.Print("Enter Key Passphrase: ")
		azKeyPass, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return err
		}
		AzKeyPass.Set((strings.TrimSpace(string(azKeyPass))))
	}

	if AzUsername.Value() == "-" {
		r := bufio.NewReader(os.Stdin)
		fmt.Print("Enter username: ")
		azUser, err := r.ReadString('\n')
		if err != nil {
			return err
		}
		AzUsername.Set(strings.TrimSpace(azUser))
	}

	if AzPassword.Value() == "-" {
		fmt.Print("Enter password: ")
		azPass, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return err
		}
		AzPassword.Set((strings.TrimSpace(string(azPass))))
	}

	//newline to not mess with following logs
	fmt.Println()
	return nil
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

func Options() config.Options {
	return config.Options{
		ConfigFile:  ConfigFile.Value().(string),
		ConfigName:  "config",
		ConfigPaths: SystemConfigDirs(),
		EnvPrefix:   EnvPrefix,
	}
}
