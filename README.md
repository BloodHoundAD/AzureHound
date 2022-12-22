# AzureHound

The BloodHound data collector for Microsoft Azure


![GitHub Workflow Status](https://img.shields.io/github/workflow/status/BloodHoundAD/AzureHound/Build)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/BloodHoundAD/AzureHound)
![GitHub all releases](https://img.shields.io/github/downloads/BloodHoundAD/AzureHound/total)
[![Documentation](https://img.shields.io/static/v1?label=&message=documentation&color=blue)](https://pkg.go.dev/github.com/bloodhoundad/azurehound)

## Get AzureHound

### Release Binaries

Download the appropriate binary for your platform from one of our [Releases](https://github.com/BloodHoundAD/AzureHound/releases).

#### Rolling Release

The rolling release contains pre-built binaries that are automatically kept up-to-date with the `main` branch and can be downloaded from
[here](https://github.com/BloodHoundAD/AzureHound/releases/tag/rolling).

> **Warning:** The rolling release may be unstable.

## Compiling

##### Prerequisites

- [Go 1.18](https://go.dev/dl/) or later

To build this project from source run the following:

``` sh
go build -ldflags="-s -w -X github.com/bloodhoundad/azurehound/constants.Version=`git describe tags --exact-match 2> /dev/null || git rev-parse HEAD`"
```

## Usage

### Quickstart 

**Print all Azure Tenant data to stdout**
``` sh
❯ azurehound list -u "$USERNAME" -p "$PASSWORD" -t "$TENANT"
```

**Print all Azure Tenant data to file**
``` sh
❯ azurehound list -u "$USERNAME" -p "$PASSWORD" -t "$TENANT" -o "mytenant.json"
```

**Configure and start data collection service for BloodHound Enterprise**
``` sh
❯ azurehound configure
(follow prompts)

❯ azurehound start
```

### CLI

```
❯ azurehound --help
AzureHound vx.x.x
Created by the BloodHound Enterprise team - https://bloodhoundenterprise.io

The official tool for collecting Azure data for BloodHound and BloodHound Enterprise

Usage:
  azurehound [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  configure   Configure AzureHound
  help        Help about any command
  list        Lists Azure Objects
  start       Start Azure data collection service for BloodHound Enterprise

Flags:
  -c, --config string          AzureHound configuration file (default: /Users/dlees/.config/azurehound/config.json)
  -h, --help                   help for azurehound
      --json                   Output logs as json
  -j, --jwt string             Use an acquired JWT to authenticate into Azure
      --log-file string        Output logs to this file
      --proxy string           Sets the proxy URL for the AzureHound service
  -r, --refresh-token string   Use an acquired refresh token to authenticate into Azure
  -v, --verbosity int          AzureHound verbosity level (defaults to 0) [Min: -1, Max: 2]
      --version                version for azurehound

Use "azurehound [command] --help" for more information about a command.
```
