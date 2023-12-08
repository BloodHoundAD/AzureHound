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

// Specifies information about the marketplace image used to create the virtual machine. This element is only used for
// marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic
// use. In the Azure portal, find the marketplace image that you want to use and then click Want to deploy
// programmatically, Get Started ->. Enter any required information and then click Save.
type Plan struct {
	// The plan ID.
	Name string `json:"name,omitempty"`

	// Specifies the product of the image from the marketplace. This is the same value as Offer under the imageReference
	// element.
	Product string `json:"product,omitempty"`

	// The promotion code.
	PromotionCode string `json:"promotionCode,omitempty"`

	// The publisher ID.
	Publisher string `json:"publisher,omitempty"`
}
