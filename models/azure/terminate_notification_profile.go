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

package azure

type TerminateNotificationProfile struct {
	// Specifies whether the Terminate Scheduled event is enabled or disabled.
	Enable bool `json:"enable,omitempty"`

	// Configurable length of time a Virtual Machine being deleted will have to potentially approve the
	// Terminate Scheduled Event before the event is auto approved (timed out).
	// The configuration must be specified in ISO 8601 format, the default value is 5 minutes (PT5M)
	NotBeforeTimeout string `json:"notBeforeTimeout,omitempty"`
}
