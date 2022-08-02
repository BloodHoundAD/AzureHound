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

type ServicePrincipalType string

const (
	// A service principal that represents an application or service.
	// The appId property identifies the associated app registration, and matches the appId of an application, possibly
	// from a different tenant. If the associated app registration is missing, tokens are not issued for the service
	// principal.
	ServicePrincipalTypeApplication ServicePrincipalType = "Application"

	// A service principal that represents a managed identity. Service principals representing managed identities can be
	// granted access and permissions, but cannot be updated or modified directly.
	ServicePrincipalTypeManagedIdentities ServicePrincipalType = "ManagedIdentities"

	// A service principal that represents an app created before app registrations, or through legacy experiences.
	// Legacy service principal can have credentials, service principal names, reply URLs, and other properties which
	// are editable by an authorized user, but does not have an associated app registration. The appId value does not
	// associate the service principal with an app registration. The service principal can only be used in the tenant
	// where it was created.
	ServicePrincipalTypeLegacy ServicePrincipalType = "Legacy"

	// For internal use.
	ServicePrincipalTypeSocialIDP ServicePrincipalType = "SocialIdp"
)
