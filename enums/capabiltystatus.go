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

type CapabiltyStatus string

const (
	// Available for normal use.
	CapabiltyStatusEnabled CapabiltyStatus = "Enabled"

	// Available for normal use but is in a grace period.
	CapabiltyStatusWarning CapabiltyStatus = "Warning"

	// Unavailable but any data associated with the capability must be preserved.
	CapabiltyStatusSuspended CapabiltyStatus = "Suspended"

	// Unavailable and any data associated with the capability may be deleted.
	CapabiltyStatusDeleted CapabiltyStatus = "Deleted"

	// Unavailable for all administrators and users but any data associated with the capability must be preserved.
	CapabiltyStatusLockedOut CapabiltyStatus = "LockedOut"
)
