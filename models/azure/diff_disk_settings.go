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

// Describes the parameters of ephemeral disk settings that can be specified for operating system disk.
// NOTE: The ephemeral disk settings can only be specified for managed disk.
type DiffDiskSettings struct {
	// Specifies the ephemeral disk settings for operating system disk.
	Option string `json:"option,omitempty"`

	// Specifies the ephemeral disk placement for operating system disk.
	// Possible values are:
	// - CacheDisk
	// - ResourceDisk
	//
	// Default: CacheDisk if one is configured for the VM size otherwise ResourceDisk is used.
	// Refer to VM size documentation for Windows VM at https://docs.microsoft.com/azure/virtual-machines/windows/sizes
	// and Linux VM at https://docs.microsoft.com/azure/virtual-machines/linux/sizes to check which VM sizes exposes a
	// cache disk.
	Placement string `json:"placement,omitempty"`
}
