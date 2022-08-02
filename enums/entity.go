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

type Entity = string

const (
	EntityUser                    Entity = "#microsoft.graph.user"
	EntityInvitation              Entity = "#microsoft.graph.invitation"
	EntityAppTemplate             Entity = "#microsoft.graph.applicationTemplate"
	EntityAuthMethodConfig        Entity = "#microsoft.graph.authenticationMethodConfiguration"
	EntityIdentityProvider        Entity = "#microsoft.graph.identityProvider"
	EntityApplication             Entity = "#microsoft.graph.application"
	EntityCertBasedAuthConfig     Entity = "#microsoft.graph.certificateBasedAuthConfiguration"
	EntityOrgContact              Entity = "#microsoft.graph.orgContact"
	EntityContract                Entity = "#microsoft.graph.contract"
	EntityDevice                  Entity = "#microsoft.graph.device"
	EntityDirectoryObject         Entity = "#microsoft.graph.directoryObject"
	EntityDirectoryRole           Entity = "#microsoft.graph.directoryRole"
	EntityDirectoryRoleTemplate   Entity = "#microsoft.graph.directoryRoleTemplate"
	EntityDomainDNSRecord         Entity = "#microsoft.graph.domainDnsRecord"
	EntityDomain                  Entity = "#microsoft.graph.domain"
	EntityGroup                   Entity = "#microsoft.graph.group"
	EntityGroupSetting            Entity = "#microsoft.graph.groupSetting"
	EntityGroupSettingTemplate    Entity = "#microsoft.graph.groupSettingTemplate"
	EntityOrgBrandingLocalization Entity = "#microsoft.graph.organizationalBrandingLocalization"
	EntityOAuth2PermissionGrant   Entity = "#microsoft.graph.oAuth2PermissionGrant"
	EntityOrganization            Entity = "#microsoft.graph.organization"
	EntityResourcePermissionGrant Entity = "#microsoft.graph.resourceSpecificPermissionGrant"
	EntityScopedRoleMembership    Entity = "#microsoft.graph.scopedRoleMembership"
	EntityServicePrincipal        Entity = "#microsoft.graph.servicePrincipal"
	EntitySubscribedSku           Entity = "#microsoft.graph.subscribedSku"
	EntityPlace                   Entity = "#microsoft.graph.place"
	EntityDrive                   Entity = "#microsoft.graph.drive"
	EntitySharedDriveItem         Entity = "#microsoft.graph.sharedDriveItem"
	EntitySite                    Entity = "#microsoft.graph.site"
	EntitySchemaExt               Entity = "#microsoft.graph.schemaExtension"
	EntityGroupLifecyclePolicy    Entity = "#microsoft.graph.groupLifecyclePolicy"
	EntityAgreementAcceptance     Entity = "#microsoft.graph.agreementAcceptance"
	EntityAgreement               Entity = "#microsoft.graph.agreement"
	EntityDataPolicyOperation     Entity = "#microsoft.graph.dataPolicyOperation"
	EntitySubscription            Entity = "#microsoft.graph.subscription"
	EntityExternalConnection      Entity = "#microsoft.graph.externalConnection"
	EntityChat                    Entity = "#microsoft.graph.chat"
	EntityTeam                    Entity = "#microsoft.graph.team"
	EntityTeamsTemplate           Entity = "#microsoft.graph.teamsTemplate"
)
