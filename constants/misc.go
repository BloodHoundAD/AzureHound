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

package constants

import "fmt"

// AzureHound version
// This gets populated at build time when the command being run uses the following flag:
// -ldflags "-X constants.Version=`git describe --tags --abbrev=0`"
var Version string = "v0.0.0"

const (
	Name        string = "azurehound"
	DisplayName string = "AzureHound"
	Description string = "The official tool for collecting Azure data for BloodHound and BloodHound Enterprise"
	AuthorRef   string = "Created by the BloodHound Enterprise team - https://bloodhoundenterprise.io"
)

// Returns a properly formatted value for the User-Agent header
func UserAgent() string {
	return fmt.Sprintf("%s/%s", Name, Version)
}

// Azure Services
const (
	GraphApiVersion string = "v1.0"
)
