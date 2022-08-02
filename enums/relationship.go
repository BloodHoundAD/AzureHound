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

type Relationship string

// relationshiperated relationships
const (
	RelationshipAZAvereContributor Relationship = "AZAvereContributor"
	RelationshipAZContains         Relationship = "AZContains"
	RelationshipAZContributor      Relationship = "AZContributor"
	RelationshipAZGetCertificates  Relationship = "AZGetCertificates"
	RelationshipAZGetKeys          Relationship = "AZGetKeys"
	RelationshipAZGetSecrets       Relationship = "AZGetSecrets"
	RelationshipAZHasRole          Relationship = "AZHasRole"
	RelationshipAZMemberOf         Relationship = "AZMemberOf"
	RelationshipAZOwner            Relationship = "AZOwner"
	RelationshipAZRunsAs           Relationship = "AZRunsAs"
	RelationshipAZVMContributor    Relationship = "AZVMContributor"
)

// Post-processed relationships
const (
	RelationshipAZAddMembers              Relationship = "AZAddMembers"
	RelationshipAZAddSecret               Relationship = "AZAddSecret"
	RelationshipAZExecuteCommand          Relationship = "AZExecuteCommand"
	RelationshipAZGlobalAdmin             Relationship = "AZGlobalAdmin"
	RelationshipAZGrant                   Relationship = "AZGrant"
	RelationshipAZGrantSelf               Relationship = "AZGrantSelf"
	RelationshipAZPrivilegedRoleAdmin     Relationship = "AZPrivilegedRoleAdmin"
	RelationshipAZAZResetPassword         Relationship = "AZAZResetPassword"
	RelationshipAZUserAccessAdministrator Relationship = "AZUserAccessAdministrator"
)
