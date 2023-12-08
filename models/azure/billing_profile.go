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

// Specifies the billing related details of a Azure Spot VM or VMSS.
type BillingProfile struct {

	// Specifies the maximum price you are willing to pay for a Azure Spot VM/VMSS. This price is in US Dollars.
	// This price will be compared with the current Azure Spot price for the VM size. Also, the prices are compared at
	// the time of create/update of Azure Spot VM/VMSS and the operation will only succeed if the maxPrice is greater
	// than the current Azure Spot price.
	//
	// The maxPrice will also be used for evicting a Azure Spot VM/VMSS if the current Azure Spot price goes beyond the
	// maxPrice after creation of VM/VMSS.
	//
	// Possible values are:
	// - Any decimal value greater than zero. Example: 0.01538
	// -1 – indicates default price to be up-to on-demand.
	//
	// You can set the maxPrice to -1 to indicate that the Azure Spot VM/VMSS should not be evicted for price reasons.
	// Also, the default max price is -1 if it is not provided by you.
	//
	// Minimum api-version: 2019-03-01.
	MaxPrice float64 `json:"maxPrice,omitempty"`
}
