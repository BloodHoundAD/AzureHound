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

// Specifies information about the image to use. You can specify information about platform images, marketplace images,
// or virtual machine images. This element is required when you want to use a platform image, marketplace image, or
// virtual machine image, but is not used in other creation operations.
// NOTE: Image reference publisher and offer can only be set when you create the scale set.
type ImageReference struct {
	// Specifies in decimal numbers, the version of platform image or marketplace image used to create the virtual
	// machine. This readonly field differs from 'version', only if the value specified in 'version' field is 'latest'.
	ExactVersion string `json:"exactVersion"`

	// Resource ID.
	Id string `json:"id"`

	// Specifies the offer of the platform image or marketplace image used to create the virtual machine.
	Offer string `json:"offer"`

	// The image publisher
	Publisher string `json:"publisher"`

	// Specified the shared gallery image unique id for vm deployment.
	// This can be fetched from shared gallery image GET call.
	SharedGalleryImageId string `json:"sharedGalleryImageId"`

	// The image SKU.
	Sku string `json:"sku"`

	// Specifies the version of the platform image or marketplace image used to create the virtual machine.
	// The allowed formats are Major.Minor.Build or 'latest'. Major, Minor, and Build are decimal numbers.
	// Specify 'latest' to use the latest version of an image available at deploy time. Even if you use 'latest',
	// the VM image will not automatically update after deploy time even if a new version becomes available.
	Version string `json:"version"`
}
