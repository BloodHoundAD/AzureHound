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

// Specifies the required information to reference a compute gallery application version.
type VMGalleryApplication struct {
	// Optional, Specifies the uri to an azure blob that will replace the default configuration for the package if
	// provided.
	ConfigurationReference string `json:"configurationReference,omitempty"`

	// Optional, Specifies the order in which the packages have to be installed.
	Order int `json:"order,omitempty"`

	// Specifies the GalleryApplicationVersion resource id on the form of
	// /subscriptions/{SubscriptionId}/resourceGroups/{ResourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{application}/versions/{version}
	PackageReferenceId string `json:"packageReferenceId,omitempty"`

	// Optional, Specifies a passthrough value for more generic context.
	Tags string `json:"tags,omitempty"`
}
