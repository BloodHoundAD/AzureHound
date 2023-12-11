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

// Specifies information about the operating system disk used by the virtual machine.
// For more information about disks, see About disks and VHDs for Azure virtual machines.
type OSDisk struct {
	// Specifies the caching requirements.
	// Possible values are:
	// None
	// ReadOnly
	// ReadWrite
	//
	// Default: None for Standard storage. ReadOnly for Premium storage.
	Caching string `json:"caching,omitempty"`

	// Specifies how the virtual machine should be created.
	// Possible values are:
	// Attach - This value is used when you are using a specialized disk to create the virtual machine.
	// FromImage - This value is used when you are using an image to create the virtual machine. If you are using a platform image, you also use the imageReference element described above. If you are using a marketplace image, you also use the plan element previously described.
	CreateOption string `json:"createOption,omitempty"`

	// Specifies whether data disk should be deleted or detached upon VM deletion.
	// Possible values:
	// Delete - If this value is used, the data disk is deleted when VM is deleted.
	// Detach - If this value is used, the data disk is retained after VM is deleted.
	// The default value is set to detach. For an ephemeral OS Disk, the default value is set to Delete. User cannot change the delete option for ephemeral OS Disk.
	DeleteOption string `json:"deleteOption,omitempty"`

	// Specifies the ephemeral Disk Settings for the operating system disk used by the virtual machine.
	DiffDiskSettings DiffDiskSettings `json:"diffDiskSettings,omitempty"`

	// Specifies the size of an empty data disk in gigabytes.
	// This element can be used to overwrite the size of the disk in a virtual machine image.
	// This value cannot be larger than 1023 GB
	DiskSizeGB int `json:"diskSizeGB,omitempty"`

	// Specifies the encryption settings for the OS Disk.
	// Minimum api-version: 2015-06-15
	EncryptionSettings DiskEncryptionSettings `json:"encryptionSettings,omitempty"`

	// The source user image virtual hard disk. The virtual hard disk will be copied before being attached to the
	// virtual machine. If SourceImage is provided, the destination virtual hard drive must not exist.
	Image VirtualHardDisk `json:"image,omitempty"`

	// The managed disk parameters.
	ManagedDisk ManagedDiskParameters `json:"managedDisk,omitempty"`

	// The disk name.
	Name string `json:"name,omitempty"`

	// This property allows you to specify the type of the OS that is included in the disk if creating a VM from user-image or a specialized VHD.
	// Possible values are:
	// - Windows
	// - Linux
	OSType string `json:"osType,omitempty"`

	// The virtual hard disk.
	Vhd VirtualHardDisk `json:"vhd,omitempty"`

	// Specifies whether writeAccelerator should be enabled or disabled on the disk.
	WriteAcceleratorEnabled bool `json:"writeAcceleratorEnabled,omitempty"`
}
