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

// Describes a data disk.
type DataDisk struct {
	// Specifies the caching requirements.
	// Possible values are:
	// None
	// ReadOnly
	// ReadWrite
	//
	// Default: None for Standard storage. ReadOnly for Premium storage
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
	// The default value is set to detach
	DeleteOption string `json:"deleteOption,omitempty"`

	// Specifies the detach behavior to be used while detaching a disk or which is already in the process of detachment
	// from the virtual machine.
	// Supported values: ForceDetach
	//
	// ForceDetach is applicable only for managed data disks. If a previous detachment attempt of the data disk did not
	// complete due to an unexpected failure from the virtual machine and the disk is still not released then use
	// force-detach as a last resort option to detach the disk forcibly from the VM. All writes might not have been
	// flushed when using this detach behavior.
	//
	// This feature is still in preview mode and is not supported for VirtualMachineScaleSet. To force-detach a data disk
	// update toBeDetached to 'true' along with setting detachOption: 'ForceDetach'.
	DetachOption string `json:"detachOption,omitempty"`

	// Specifies the Read-Write IOPS for the managed disk when StorageAccountType is UltraSSD_LRS.
	// Returned only for VirtualMachine ScaleSet VM disks. Can be updated only via updates to the
	// VirtualMachine Scale Set.
	DiskIOPSReadWrite int `json:"diskIOPSReadWrite,omitempty"`

	// Specifies the bandwidth in MB per second for the managed disk when StorageAccountType is UltraSSD_LRS.
	// Returned only for VirtualMachine ScaleSet VM disks. Can be updated only via updates to the
	// VirtualMachine Scale Set.
	DiskMBpsReadWrite int `json:"diskMBpsReadWrite,omitempty"`

	// Specifies the size of an empty data disk in gigabytes.
	// This element can be used to overwrite the size of the disk in a virtual machine image.
	// This value cannot be larger than 1023 GB
	DiskSizeGB int `json:"diskSizeGB,omitempty"`

	// The source user image virtual hard disk. The virtual hard disk will be copied before being attached to the
	// virtual machine. If SourceImage is provided, the destination virtual hard drive must not exist.
	Image VirtualHardDisk `json:"image,omitempty"`

	// Specifies the logical unit number of the data disk.
	// This value is used to identify data disks within the VM and therefore must be unique for each data disk attached
	// to a VM.
	Lun int `json:"lun,omitempty"`

	// The managed disk parameters.
	ManagedDisk ManagedDiskParameters `json:"managedDisk,omitempty"`

	// The disk name.
	Name string `json:"name,omitempty"`

	// Specifies whether the data disk is in process of detachment from the VirtualMachine/VirtualMachineScaleset.
	ToBeDetached bool `json:"toBeDetached,omitempty"`

	// The virtual hard disk.
	Vhd VirtualHardDisk `json:"vhd,omitempty"`

	// Specifies whether writeAccelerator should be enabled or disabled on the disk.
	WriteAcceleratorEnabled bool `json:"writeAcceleratorEnabled,omitempty"`
}
