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

// Specifies the storage settings for the virtual machine disks.
type StorageProfile struct {
	// Specifies the parameters that are used to add a data disk to a virtual machine.
	// For more information about disks, see About disks and VHDs for Azure virtual machines.
	DataDisks []DataDisk `json:"dataDisks,omitempty"`

	// Specifies information about the image to use. You can specify information about platform images, marketplace
	// images, or virtual machine images. This element is required when you want to use a platform image, marketplace
	// image, or virtual machine image, but is not used in other creation operations.
	ImageReference ImageReference `json:"imageReference,omitempty"`

	// Specifies information about the operating system disk used by the virtual machine.
	// For more information about disks, see About disks and VHDs for Azure virtual machines.
	OSDisk OSDisk `json:"osDisk,omitempty"`
}
