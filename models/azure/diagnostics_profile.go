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

// Specifies the boot diagnostic settings state.
// Minimum api-version: 2015-06-15.
type DiagnosticsProfile struct {
	// Boot Diagnostics is a debugging feature which allows you to view Console Output and Screenshot to diagnose VM
	// status. You can easily view the output of your console log. Azure also enables you to see a screenshot of the VM
	// from the hypervisor.
	BootDiagnotics BootDiagnotics `json:"bootDiagnotics,omitempty"`
}
