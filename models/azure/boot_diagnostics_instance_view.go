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

// The instance view of a virtual machine boot diagnostics.
type BootDiagnoticsInstanceView struct {
	// The console screenshot blob URI.
	// NOTE: This will not be set if boot diagnostics is currently enabled with managed storage.
	ConsoleScreenshotBlobUri string `json:"consoleScreenshotBlobUri,omitempty"`

	// The serial console log blob Uri.
	// NOTE: This will not be set if boot diagnostics is currently enabled with managed storage.
	SerialConsoleLogBlobUri string `json:"serialConsoleLogBlobUri,omitempty"`

	// The boot diagnostics status information for the VM.
	// NOTE: It will be set only if there are errors encountered in enabling boot diagnostics.
	Status InstanceViewStatus `json:"status,omitempty"`
}
