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

package enums

type SkuName string

const (
	SKU_Premium_LRS     SkuName = "Premium_LRS"
	SKU_Premium_ZRS     SkuName = "Premium_ZRS"
	SKU_Standard_GRS    SkuName = "Standard_GRS"
	SKU_Standard_GZRS   SkuName = "Standard_GZRS"
	SKU_Standard_LRS    SkuName = "Standard_LRS"
	SKU_Standard_RAGRS  SkuName = "Standard_RAGRS"
	SKU_Standard_RAGZRS SkuName = "Standard_RAGZRS"
	SKU_Standard_ZRS    SkuName = "Standard_ZRS"
	SKU_Basic           SkuName = "Basic"
	SKU_Free            SkuName = "Free"
	SKU_NotSpecified    SkuName = "NotSpecified"
	SKU_Premium         SkuName = "Premium"
	SKU_Shared          SkuName = "Shared"
	SKU_Standard        SkuName = "Standard"
)
