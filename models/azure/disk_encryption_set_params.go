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

// Describes the parameter of customer managed disk encryption set resource id that can be specified for disk.
// NOTE: The disk encryption set resource id can only be specified for managed disk.
// Please refer https://aka.ms/mdssewithcmkoverview for more details.
type DiskEncryptionSetParameters struct {
	// Resource ID.
	Id string `json:"id,omitempty"`
}
