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

package constants

// Azure AD built-in roles
// See https://docs.microsoft.com/en-us/azure/active-directory/roles/permissions-reference for more info.
const (
	// Can create and manage all aspects of app registrations and enterprise apps.
	ApplicationAdministratorRoleID string = "9b895d92-2cd3-44c7-9d02-a6ac2d5ea5c3"

	// Can create application registrations independent of the 'Users can register applications' setting.
	ApplicationDeveloperRoleID string = "cf1c38e5-3621-4004-a7cb-879624dced7c"

	// Can create attack payloads that an administrator can initiate later.
	AttackPayloadAuthorRoleID string = "9c6df0f2-1e7c-4dc3-b195-66dfbd24aa8f"

	// Can create and manage all aspects of attack simulation campaigns.
	AttackSimulationAdministratorRoleID string = "c430b396-e693-46cc-96f3-db01bf8bb62a"

	// Assign custom security attribute keys and values to supported Azure AD objects.
	AttributeAssignmentAdministratorRoleID string = "58a13ea3-c632-46ae-9ee0-9c0d43cd7f3d"

	// Read custom security attribute keys and values for supported Azure AD objects.
	AttributeAssignmentReaderRoleID string = "ffd52fa5-98dc-465c-991d-fc073eb59f8f"

	// Define and manage the definition of custom security attributes.
	AttributeDefinitionAdministratorRoleID string = "8424c6f0-a189-499e-bbd0-26c1753c96d4"

	// Read the definition of custom security attributes.
	AttributeDefinitionReaderRoleID string = "1d336d2c-4ae8-42ef-9711-b3604ce3fc2c"

	// Can access to view, set and reset authentication method information for any non-admin user.
	AuthenticationAdministratorRoleID string = "c4e39bd9-1100-46d3-8c65-fb160da0071f"

	// Can create and manage the authentication methods policy, tenant-wide MFA settings, password protection policy, and verifiable credentials.
	AuthenticationPolicyAdministratorRoleID string = "0526716b-113d-4c15-b2c8-68e3c22b9f80"

	// Users assigned to this role are added to the local administrators group on Azure AD-joined devices.
	AzureADJoinedDeviceLocalAdministratorRoleID string = "9f06204d-73c1-4d4c-880a-6edb90606fd8"

	// Can manage Azure DevOps organization policy and settings.
	AzureDevOpsAdministratorRoleID string = "e3973bdf-4987-49ae-837a-ba8e231c7286"

	// Can manage all aspects of the Azure Information Protection product.
	AzureInformationProtectionAdministratorRoleID string = "7495fdc4-34c4-4d15-a289-98788ce399fd"

	// Can manage secrets for federation and encryption in the Identity Experience Framework (IEF).
	B2CIEFKeysetAdministratorRoleID string = "aaf43236-0c0d-4d5f-883a-6955382ac081"

	// Can create and manage trust framework policies in the Identity Experience Framework (IEF).
	B2CIEFPolicyAdministratorRoleID string = "3edaf663-341e-4475-9f94-5c398ef6c070"

	// Can perform common billing related tasks like updating payment information.
	BillingAdministratorRoleID string = "b0f54661-2d74-4c50-afa3-1ec803f12efe"

	// Can manage all aspects of the Cloud App Security product.
	CloudAppSecurityAdministratorRoleID string = "892c5842-a9a6-463a-8041-72aa08ca3cf6"

	// Can create and manage all aspects of app registrations and enterprise apps except App Proxy.
	CloudApplicationAdministratorRoleID string = "158c047a-c907-4556-b7ef-446551a6b5f7"

	// Limited access to manage devices in Azure AD.
	CloudDeviceAdministratorRoleID string = "7698a772-787b-4ac8-901f-60d6b08affd2"

	// Can read and manage compliance configuration and reports in Azure AD and Microsoft 365.
	ComplianceAdministratorRoleID string = "17315797-102d-40b4-93e0-432062caca18"

	// Creates and manages compliance content.
	ComplianceDataAdministratorRoleID string = "e6d1a23a-da11-4be4-9570-befc86d067a7"

	// Can manage Conditional Access capabilities.
	ConditionalAccessAdministratorRoleID string = "b1be1c3e-b65d-4f19-8427-f6fa0d97feb9"

	// Can approve Microsoft support requests to access customer organizational data.
	CustomerLockBoxAccessApproverRoleID string = "5c4f9dcd-47dc-4cf7-8c9a-9e4207cbfc91"

	// Can access and manage Desktop management tools and services.
	DesktopAnalyticsAdministratorRoleID string = "38a96431-2bdf-4b4c-8b6e-5d3d8abac1a4"

	// Deprecated - Do Not Use.
	DeviceJoinRoleID string = "9c094953-4995-41c8-84c8-3ebb9b32c93f"

	// Deprecated - Do Not Use.
	DeviceManagersRoleID string = "2b499bcd-da44-4968-8aec-78e1674fa64d"

	// Deprecated - Do Not Use.
	DeviceUsersRoleID string = "d405c6df-0af8-4e3b-95e4-4d06e542189e"

	// Can read basic directory information. Commonly used to grant directory read access to applications and guests.
	DirectoryReadersRoleID string = "88d8e3e3-8f55-4a1e-953a-9b9898b8876b"

	// Only used by Azure AD Connect service.
	DirectorySynchronizationAccountsRoleID string = "d29b2b05-8046-44ba-8758-1e26182fcf32"

	// Can read and write basic directory information. For granting access to applications, not intended for users.
	DirectoryWritersRoleID string = "9360feb5-f418-4baa-8175-e2a00bac4301"

	// Can manage domain names in cloud and on-premises.
	DomainNameAdministratorRoleID string = "8329153b-31d0-4727-b945-745eb3bc5f31"

	// Can manage all aspects of the Dynamics 365 product.
	Dynamics365AdministratorRoleID string = "44367163-eba1-44c3-98af-f5787879f96a"

	// Manage all aspects of Microsoft Edge.
	EdgeAdministratorRoleID string = "3f1acade-1e04-4fbc-9b69-f0302cd84aef"

	// Can manage all aspects of the Exchange product.
	ExchangeAdministratorRoleID string = "29232cdf-9323-42fd-ade2-1d097af3e4de"

	// Can create or update Exchange Online recipients within the Exchange Online organization.
	ExchangeRecipientAdministratorRoleID string = "31392ffb-586c-42d1-9346-e59415a2cc4e"

	// Can create and manage all aspects of user flows.
	ExternalIDUserFlowAdministratorRoleID string = "6e591065-9bad-43ed-90f3-e9424366d2f0"

	// Can create and manage the attribute schema available to all user flows.
	ExternalIDUserFlowAttributeAdministratorRoleID string = "0f971eea-41eb-4569-a71e-57bb8a3eff1e"

	// Can configure identity providers for use in direct federation.
	ExternalIdentityProviderAdministratorRoleID string = "be2f45a1-457d-42af-a067-6ec1fa63bc45"

	// Can manage all aspects of Azure AD and Microsoft services that use Azure AD identities.
	GlobalAdministratorRoleID string = "62e90394-69f5-4237-9190-012177145e10"

	// Can read everything that a Global Administrator can, but not update anything.
	GlobalReaderRoleID string = "f2ef992c-3afb-46b9-b7cf-a126ee74c451"

	// Members of this role can create/manage groups, create/manage groups settings like naming and expiration policies, and view groups activity and audit reports.
	GroupsAdministratorRoleID string = "fdd7a751-b60b-444a-984c-02652fe8fa1c"

	// Can invite guest users independent of the 'members can invite guests' setting.
	GuestInviterRoleID string = "95e79109-95c0-4d8e-aee3-d01accf2d47b"

	// Default role for guest users. Can read a limited set of directory information.
	GuestUserRoleID string = "10dae51f-b6af-4016-8d66-8c2a99b929b3"

	// Can reset passwords for non-administrators and Helpdesk Administrators.
	HelpdeskAdministratorRoleID string = "729827e3-9c14-49f7-bb1b-9608f156bbb8"

	// Can manage AD to Azure AD cloud provisioning, Azure AD Connect, and federation settings.
	HybridIdentityAdministratorRoleID string = "8ac3fc64-6eca-42ea-9e69-59f4c7b60eb2"

	// Manage access using Azure AD for identity governance scenarios.
	IdentityGovernanceAdministratorRoleID string = "45d8d3c5-c802-45c6-b32a-1d70b5e1e86e"

	// Has administrative access in the Microsoft 365 Insights app.
	InsightsAdministratorRoleID string = "eb1f4a8d-243a-41f0-9fbd-c7cdf6c5ef7c"

	// Access the analytical capabilities in Microsoft Viva Insights and run custom queries.
	InsightsAnalystRoleID string = "25df335f-86eb-4119-b717-0ff02de207e9"

	// Can view and share dashboards and insights via the M365 Insights app.
	InsightsBusinessLeaderRoleID string = "31e939ad-9672-4796-9c2e-873181342d2d"

	// Can manage all aspects of the Intune product.
	IntuneAdministratorRoleID string = "3a2c62db-5318-420d-8d74-23affee5d9d5"

	// Can manage settings for Microsoft Kaizala.
	KaizalaAdministratorRoleID string = "74ef975b-6605-40af-a5d2-b9539d836353"

	// Can configure knowledge, learning, and other intelligent features.
	KnowledgeAdministratorRoleID string = "b5a8dcf3-09d5-43a9-a639-8e29ef291470"

	// Has access to topic management dashboard and can manage content.
	KnowledgeManagerRoleID string = "744ec460-397e-42ad-a462-8b3f9747a02c"

	// Can manage product licenses on users and groups.
	LicenseAdministratorRoleID string = "4d6ac14f-3453-41d0-bef9-a3e0c569773a"

	// Create and manage all aspects of workflows and tasks associated with Lifecycle Workflows in Azure AD.
	LifecycleWorkflowsAdministratorRoleID string = "59d46f88-662b-457b-bceb-5c3809e5908f"

	// Can read security messages and updates in Office 365 Message Center only.
	MessageCenterPrivacyReaderRoleID string = "ac16e43d-7b2d-40e0-ac05-243ff356ab5b"

	// Can read messages and updates for their organization in Office 365 Message Center only.
	MessageCenterReaderRoleID string = "790c1fb9-7f7d-4f88-86a1-ef1f95c05c1b"

	// Can manage network locations and review enterprise network design insights for Microsoft 365 Software as a Service applications.
	NetworkAdministratorRoleID string = "d37c8bed-0711-4417-ba38-b4abe66ce4c2"

	// Can manage Office apps cloud services, including policy and settings management, and manage the ability to select, unselect and publish 'what's new' feature content to end-user's devices.
	OfficeAppsAdministratorRoleID string = "2b745bdf-0803-4d80-aa65-822c4493daac"

	// Do not use - not intended for general use.
	PartnerTier1SupportRoleID string = "4ba39ca4-527c-499a-b93d-d9b492c50246"

	// Do not use - not intended for general use.
	PartnerTier2SupportRoleID string = "e00e864a-17c5-4a4b-9c06-f5b95a8d5bd8"

	// Can reset passwords for non-administrators and Password Administrators.
	PasswordAdministratorRoleID string = "966707d0-3269-4727-9be2-8c3a10f19b9d"

	// Manage all aspects of Entra Permissions Management.
	PermissionsManagementAdministratorRoleID string = "af78dc32-cf4d-46f9-ba4e-4428526346b5"

	// Can manage all aspects of the Power BI product.
	PowerBIAdministratorRoleID string = "a9ea8996-122f-4c74-9520-8edcd192826c"

	// Can create and manage all aspects of Microsoft Dynamics 365, PowerApps and Microsoft Flow.
	PowerPlatformAdministratorRoleID string = "11648597-926c-4cf3-9c36-bcebb0ba8dcc"

	// Can manage all aspects of printers and printer connectors.
	PrinterAdministratorRoleID string = "644ef478-e28f-4e28-b9dc-3fdde9aa0b1f"

	// Can register and unregister printers and update printer status.
	PrinterTechnicianRoleID string = "e8cef6f1-e4bd-4ea8-bc07-4b8d950f4477"

	// Can access to view, set and reset authentication method information for any user (admin or non-admin).
	PrivilegedAuthenticationAdministratorRoleID string = "7be44c8a-adaf-4e2a-84d6-ab2649e08a13"

	// Can manage role assignments in Azure AD, and all aspects of Privileged Identity Management.
	PrivilegedRoleAdministratorRoleID string = "e8611ab8-c189-46e8-94e1-60213ab1f814"

	// Can read sign-in and audit reports.
	ReportsReaderRoleID string = "4a5d8f65-41da-4de4-8968-e035b65339cf"

	// Restricted role for guest users. Can read a limited set of directory information.
	RestrictedGuestUserRoleID string = "2af84b1e-32c8-42b7-82bc-daa82404023b"

	// Can create and manage all aspects of Microsoft Search settings.
	SearchAdministratorRoleID string = "0964bb5e-9bdb-4d7b-ac29-58e794862a40"

	// Can create and manage the editorial content such as bookmarks, Q and As, locations, floorplan.
	SearchEditorRoleID string = "8835291a-918c-4fd7-a9ce-faa49f0cf7d9"

	// Can read security information and reports, and manage configuration in Azure AD and Office 365.
	SecurityAdministratorRoleID string = "194ae4cb-b126-40b2-bd5b-6091b380977d"

	// Creates and manages security events.
	SecurityOperatorRoleID string = "5f2222b1-57c3-48ba-8ad5-d4759f1fde6f"

	// Can read security information and reports in Azure AD and Office 365.
	SecurityReaderRoleID string = "5d6b6bb7-de71-4623-b4af-96380a352509"

	// Can read service health information and manage support tickets.
	ServiceSupportAdministratorRoleID string = "f023fd81-a637-4b56-95fd-791ac0226033"

	// Can manage all aspects of the SharePoint service.
	SharePointAdministratorRoleID string = "f28a1f50-f6e7-4571-818b-6a12f2af6b6c"

	// Can manage all aspects of the Skype for Business product.
	SkypeforBusinessAdministratorRoleID string = "75941009-915a-4869-abe7-691bff18279e"

	// Can manage the Microsoft Teams service.
	TeamsAdministratorRoleID string = "69091246-20e8-4a56-aa4d-066075b2a7a8"

	// Can manage calling and meetings features within the Microsoft Teams service.
	TeamsCommunicationsAdministratorRoleID string = "baf37b3a-610e-45da-9e62-d9d1e5e8914b"

	// Can troubleshoot communications issues within Teams using advanced tools.
	TeamsCommunicationsSupportEngineerRoleID string = "f70938a0-fc10-4177-9e90-2178f8765737"

	// Can troubleshoot communications issues within Teams using basic tools.
	TeamsCommunicationsSupportSpecialistRoleID string = "fcf91098-03e3-41a9-b5ba-6f0ec8188a12"

	// Can perform management related tasks on Teams certified devices.
	TeamsDevicesAdministratorRoleID string = "3d762c5a-1b6c-493f-843e-55a3b42923d4"

	// Can see only tenant level aggregates in Microsoft 365 Usage Analytics and Productivity Score.
	UsageSummaryReportsReaderRoleID string = "75934031-6c7e-415a-99d7-48dbd49e875e"

	// Default role for member users. Can read all and write a limited set of directory information.
	UserRoleID string = "a0b1b346-4d3e-4e8b-98f8-753987be4970"

	// Can manage all aspects of users and groups, including resetting passwords for limited admins.
	UserAdministratorRoleID string = "fe930be7-5e62-47db-91af-98c3a49a38b1"

	// Manage and share Virtual Visits information and metrics from admin centers or the Virtual Visits app.
	VirtualVisitsAdministratorRoleID string = "e300d9e7-4a2b-4295-9eff-f1c78b36cc98"

	// Can provision and manage all aspects of Cloud PCs.
	Windows365AdministratorRoleID string = "11451d60-acb2-45eb-a7d6-43d0f0125c13"

	// Can create and manage all aspects of Windows Update deployments through the Windows Update for Business deployment service.
	WindowsUpdateDeploymentAdministratorRoleID string = "32696413-001a-46ae-978c-ce0f6b3620d2"

	// Deprecated - Do Not Use.
	WorkplaceDeviceJoinRoleID string = "c34f683f-4d5a-4403-affd-6615e00e3a7f"

	// Manage all aspects of the Yammer service.
	YammerAdministratorRoleID string = "810a2642-a034-447f-a5e8-41beaa378541"
)

// Azure ARM roles
// See https://docs.microsoft.com/en-us/azure/role-based-access-control/built-in-roles
const (
	// Can customize the developer portal, edit its content, and publish it.
	APIManagementDeveloperPortalContentEditorRoleID string = "c031e6a8-4391-4de0-8d69-4706a7ed3729"

	// Can manage service and the APIs
	APIManagementServiceContributorRoleID string = "312a565d-c81f-4fd8-895a-4e21e48d571c"

	// Can manage service but not the APIs
	APIManagementServiceOperatorRoleID string = "e022efe7-f5ba-4159-bbe4-b44f577e9b61"

	// Read-only access to service and APIs
	APIManagementServiceReaderRoleID string = "71522526-b88f-4d52-b57f-d31fc3546d0d"

	// Lets you grant Access Review System app permissions to discover and revoke access as needed by the access review process.
	AccessReviewOperatorServiceRoleID string = "76cc9ee4-d5d3-4a45-a930-26add3d73475"

	// acr delete
	AcrDeleteRoleID string = "c2f4ef07-c644-48eb-af81-4b1b4947fb11"

	// acr image signer
	AcrImageSignerRoleID string = "6cef56e8-d556-48e5-a04f-b8e64114680f"

	// acr pull
	AcrPullRoleID string = "7f951dda-4ed3-4680-a7ca-43fe172d538d"

	// acr push
	AcrPushRoleID string = "8311e382-0749-4cb8-b61a-304f252e45ec"

	// acr quarantine data reader
	AcrQuarantineReaderRoleID string = "cdda3590-29a3-44f6-95f2-9f980659eb04"

	// acr quarantine data writer
	AcrQuarantineWriterRoleID string = "c8d4ff99-41c3-41a8-9f60-21dfdad59608"

	// Provides contribute access to manage sensor related entities in AgFood Platform Service
	AgFoodPlatformSensorPartnerContributorRoleID string = "6b77f0a0-0d89-41cc-acd1-579c22c17a67"

	// Provides admin access to AgFood Platform Service
	AgFoodPlatformServiceAdminRoleID string = "f8da80de-1ff9-4747-ad80-a19b7f6079e3"

	// Provides contribute access to AgFood Platform Service
	AgFoodPlatformServiceContributorRoleID string = "8508508a-4469-4e45-963b-2518ee0bb728"

	// Provides read access to AgFood Platform Service
	AgFoodPlatformServiceReaderRoleID string = "7ec7ccdc-f61e-41fe-9aaf-980df0a44eba"

	// Basic user role for AnyBuild. This role allows listing of agent information and execution of remote build capabilities.
	AnyBuildBuilderRoleID string = "a2138dac-4907-4679-a376-736901ed8ad8"

	// Allows full access to App Configuration data.
	AppConfigurationDataOwnerRoleID string = "5ae67dd6-50cb-40e7-96ff-dc2bfa4b606b"

	// Allows read access to App Configuration data.
	AppConfigurationDataReaderRoleID string = "516239f1-63e1-4d78-a4de-a74fb236a071"

	// Contributor of the Application Group.
	ApplicationGroupContributorRoleID string = "ca6382a4-1721-4bcf-a114-ff0c70227b6b"

	// Can manage Application Insights components
	ApplicationInsightsComponentContributorRoleID string = "ae349356-3a1b-4a5e-921d-050484c6347e"

	// Gives user permission to use Application Insights Snapshot Debugger features
	ApplicationInsightsSnapshotDebuggerRoleID string = "08954f03-6346-4c2e-81c0-ec3a5cfae23b"

	// Can read write or delete the attestation provider instance
	AttestationContributorRoleID string = "bbf86eb8-f7b4-4cce-96e4-18cddf81d86e"

	// Can read the attestation provider properties
	AttestationReaderRoleID string = "fd1bd22b-8476-40bc-a0bc-69b95687b9f3"

	// Manage azure automation resources and other resources using azure automation.
	AutomationContributorRoleID string = "f353d9bd-d4a6-484e-a77a-8050b599b867"

	// Create and Manage Jobs using Automation Runbooks.
	AutomationJobOperatorRoleID string = "4fe576fe-1146-4730-92eb-48519fa6bf9f"

	// Automation Operators are able to start, stop, suspend, and resume jobs
	AutomationOperatorRoleID string = "d3881f73-407a-4167-8283-e981cbba0404"

	// Read Runbook properties - to be able to create Jobs of the runbook.
	AutomationRunbookOperatorRoleID string = "5fb5aef8-1081-4b8e-bb16-9d5d0385bab5"

	// Grants permissions to upload and manage new Autonomous Development Platform measurements.
	AutonomousDevelopmentPlatformDataContributorRoleID string = "b8b15564-4fa6-4a59-ab12-03e1d9594795"

	// Grants full access to Autonomous Development Platform data.
	AutonomousDevelopmentPlatformDataOwnerRoleID string = "27f8b550-c507-4db9-86f2-f4b8e816d59d"

	// Grants read access to Autonomous Development Platform data.
	AutonomousDevelopmentPlatformDataReaderRoleID string = "d63b75f7-47ea-4f27-92ac-e0d173aaf093"

	// Can create and manage an Avere vFXT cluster.
	AvereContributorRoleID string = "4f8fab4f-1852-4a58-a46a-8eaf358af14a"

	// Used by the Avere vFXT cluster to manage the cluster
	AvereOperatorRoleID string = "c025889f-8102-4ebf-b32c-fc0c6f0c6bd9"

	// List cluster user credentials action.
	AzureArcEnabledKubernetesClusterUserRoleID string = "00493d72-78f6-4148-b6c5-d3ce8e4799dd"

	// Lets you manage all resources under cluster/namespace, except update or delete resource quotas and namespaces.
	AzureArcKubernetesAdminRoleID string = "dffb1e0c-446f-4dde-a09f-99eb5cc68b96"

	// Lets you manage all resources in the cluster.
	AzureArcKubernetesClusterAdminRoleID string = "8393591c-06b9-48a2-a542-1bd6b377f6a2"

	// Lets you view all resources in cluster/namespace, except secrets.
	AzureArcKubernetesViewerRoleID string = "63f0a09d-1495-4db4-a681-037d84835eb4"

	// Lets you update everything in cluster/namespace, except (cluster)roles and (cluster)role bindings.
	AzureArcKubernetesWriterRoleID string = "5b999177-9696-4545-85c7-50de3797e5a1"

	// Arc ScVmm VM Administrator has permissions to perform all ScVmm actions.
	AzureArcScVmmAdministratorRoleID string = "a92dfd61-77f9-4aec-a531-19858b406c87"

	// Azure Arc ScVmm Private Cloud User has permissions to use the ScVmm resources to deploy VMs.
	AzureArcScVmmPrivateCloudUserRoleID string = "c0781e91-8102-4553-8951-97c6d4243cda"

	// Azure Arc ScVmm Private Clouds Onboarding role has permissions to provision all the required resources for onboard and deboard vmm server instances to Azure.
	AzureArcScVmmPrivateCloudsOnboardingRoleID string = "6aac74c4-6311-40d2-bbdd-7d01e7c6e3a9"

	// Arc ScVmm VM Contributor has permissions to perform all VM actions.
	AzureArcScVmmVMContributorRoleID string = "e582369a-e17b-42a5-b10c-874c387c530b"

	// Arc VMware VM Contributor has permissions to perform all connected VMwarevSphere actions.
	AzureArcVMwareAdministratorRoleID string = "ddc140ed-e463-4246-9145-7c664192013f"

	// Azure Arc VMware Private Cloud User has permissions to use the VMware cloud resources to deploy VMs.
	AzureArcVMwarePrivateCloudUserRoleID string = "ce551c02-7c42-47e0-9deb-e3b6fc3a9a83"

	// Azure Arc VMware Private Clouds Onboarding role has permissions to provision all the required resources for onboard and deboard vCenter instances to Azure.
	AzureArcVMwarePrivateCloudsOnboardingRoleID string = "67d33e57-3129-45e6-bb0b-7cc522f762fa"

	// Arc VMware VM Contributor has permissions to perform all VM actions.
	AzureArcVMwareVMContributorRoleID string = "b748a06d-6150-4f8a-aaa9-ce3940cd96cb"

	// Can onboard Azure Connected Machines.
	AzureConnectedMachineOnboardingRoleID string = "b64e21ea-ac4e-4cdf-9dc9-5b892992bee7"

	// Can read, write, delete and re-onboard Azure Connected Machines.
	AzureConnectedMachineResourceAdministratorRoleID string = "cd570a14-e51a-42ad-bac8-bafd67325302"

	// Microsoft.AzureArcData service role to access the resources of Microsoft.AzureArcData stored with RPSAAS.
	AzureConnectedSQLServerOnboardingRoleID string = "e8113dce-c529-4d33-91fa-e9b972617508"

	// Full access role for Digital Twins data-plane
	AzureDigitalTwinsDataOwnerRoleID string = "bcd981a7-7f74-457b-83e1-cceb9e632ffe"

	// Read-only role for Digital Twins data-plane properties
	AzureDigitalTwinsDataReaderRoleID string = "d57506d4-4c8d-48b1-8587-93c323f6a5a3"

	// Allows for full access to Azure Event Hubs resources.
	AzureEventHubsDataOwnerRoleID string = "f526a384-b230-433a-b45c-95f59c4a2dec"

	// Allows receive access to Azure Event Hubs resources.
	AzureEventHubsDataReceiverRoleID string = "a638d3c7-ab3a-418d-83e6-5f17a39d4fde"

	// Allows send access to Azure Event Hubs resources.
	AzureEventHubsDataSenderRoleID string = "2b629674-e913-4c01-ae53-ef4638d8f975"

	// List cluster admin credential action.
	AzureKubernetesServiceClusterAdminRoleID string = "0ab0b1a8-8aac-4efd-b8c2-3ee1fb270be8"

	// List cluster user credential action.
	AzureKubernetesServiceClusterUserRoleID string = "4abbcc35-e782-43d8-92c5-2d3f1bd2253f"

	// Grants access to read and write Azure Kubernetes Service clusters
	AzureKubernetesServiceContributorRoleID string = "ed7f3fbd-7b88-4dd4-9017-9adb7ce333f8"

	// Deploy the Azure Policy add-on on Azure Kubernetes Service clusters
	AzureKubernetesServicePolicyAddonDeploymentRoleID string = "18ed5180-3e48-46fd-8541-4ea054d57064"

	// Lets you manage all resources under cluster/namespace, except update or delete resource quotas and namespaces.
	AzureKubernetesServiceRBACAdminRoleID string = "3498e952-d568-435e-9b2c-8d77e338d7f7"

	// Lets you manage all resources in the cluster.
	AzureKubernetesServiceRBACClusterAdminRoleID string = "b1ff04bb-8a4e-4dc4-8eb5-8693973ce19b"

	// Allows read-only access to see most objects in a namespace. It does not allow viewing roles or role bindings. This role does not allow viewing Secrets, since reading the contents of Secrets enables access to ServiceAccount credentials in the namespace, which would allow API access as any ServiceAccount in the namespace (a form of privilege escalation). Applying this role at cluster scope will give access across all namespaces.
	AzureKubernetesServiceRBACReaderRoleID string = "7f6c6a51-bcf8-42ba-9220-52d62157d7db"

	// Allows read/write access to most objects in a namespace.This role does not allow viewing or modifying roles or role bindings. However, this role allows accessing Secrets and running Pods as any ServiceAccount in the namespace, so it can be used to gain the API access levels of any ServiceAccount in the namespace. Applying this role at cluster scope will give access across all namespaces.
	AzureKubernetesServiceRBACWriterRoleID string = "a7ffa36f-339b-4b5c-8bdf-e2c188b2c0eb"

	// Grants access all Azure Maps resource management.
	AzureMapsContributorRoleID string = "dba33070-676a-4fb0-87fa-064dc56ff7fb"

	// Grants access to read, write, and delete access to map related data from an Azure maps account.
	AzureMapsDataContributorRoleID string = "8f5e0ce6-4f7b-4dcf-bddf-e6f48634a204"

	// Grants access to read map related data from an Azure maps account.
	AzureMapsDataReaderRoleID string = "423170ca-a8f6-4b0f-8487-9e4eb8f49bfa"

	// Grants access to very limited set of data APIs for common visual web SDK scenarios. Specifically, render and search data APIs.
	AzureMapsSearchandRenderDataReaderRoleID string = "6be48352-4f82-47c9-ad5e-0acacefdb005"

	// Allows for listen access to Azure Relay resources.
	AzureRelayListenerRoleID string = "26e0b698-aa6d-4085-9386-aadae190014d"

	// Allows for full access to Azure Relay resources.
	AzureRelayOwnerRoleID string = "2787bf04-f1f5-4bfe-8383-c8a24483ee38"

	// Allows for send access to Azure Relay resources.
	AzureRelaySenderRoleID string = "26baccc8-eea7-41f1-98f4-1762cc7f685d"

	// Allows for full access to Azure Service Bus resources.
	AzureServiceBusDataOwnerRoleID string = "090c5cfd-751d-490a-894a-3ce6f1109419"

	// Allows for receive access to Azure Service Bus resources.
	AzureServiceBusDataReceiverRoleID string = "4f6d3b9b-027b-4f4c-9142-0e5a2a2247e0"

	// Allows for send access to Azure Service Bus resources.
	AzureServiceBusDataSenderRoleID string = "69a216fc-b8fb-44d8-bc22-1f3c2cd27a39"

	// Allow read, write and delete access to Azure Spring Cloud Config Server
	AzureSpringCloudConfigServerContributorRoleID string = "a06f5c24-21a7-4e1a-aa2b-f19eb6684f5b"

	// Allow read access to Azure Spring Cloud Config Server
	AzureSpringCloudConfigServerReaderRoleID string = "d04c6db6-4947-4782-9e91-30a88feb7be7"

	// Allow read access to Azure Spring Cloud Data
	AzureSpringCloudDataReaderRoleID string = "b5537268-8956-4941-a8f0-646150406f0c"

	// Allow read, write and delete access to Azure Spring Cloud Service Registry
	AzureSpringCloudServiceRegistryContributorRoleID string = "f5880b48-c26d-48be-b172-7927bfa1c8f1"

	// Allow read access to Azure Spring Cloud Service Registry
	AzureSpringCloudServiceRegistryReaderRoleID string = "cff1b556-2399-4e7e-856d-a8f754be7b65"

	// Lets you manage Azure Stack registrations.
	AzureStackRegistrationOwnerRoleID string = "6f12a6df-dd06-4f3e-bcb1-ce8be600526a"

	// Azure VM Managed identities restore Contributors are allowed to perform Azure VM Restores with managed identities both user and system
	AzureVMManagedidentitiesrestoreContributorRoleID string = "6ae96244-5829-4925-a7d3-5975537d91dd"

	// Can perform all actions within an Azure Machine Learning workspace, except for creating or deleting compute resources and modifying the workspace itself.
	AzureMLDataScientistRoleID string = "f6c7c914-8db3-469d-8ca1-694a8f32e121"

	// Lets you write metrics to AzureML workspace
	AzureMLMetricsWriterRoleID string = "635dd51f-9968-44d3-b7fb-6d9a6bd613ae"

	// Lets you manage backup service,but can't create vaults and give access to others
	BackupContributorRoleID string = "5e467623-bb1f-42f4-a55d-6e525e11384b"

	// Lets you manage backup services, except removal of backup, vault creation and giving access to others
	BackupOperatorRoleID string = "00c29273-979b-4161-815c-10b084fb9324"

	// Can view backup services, but can't make changes
	BackupReaderRoleID string = "a795c7a0-d4a2-40c1-ae25-d81f01202912"

	// Allows read access to billing data
	BillingReaderRoleID string = "fa23ad8b-c56e-40d8-ac0c-ce449e1d2c64"

	// Lets you manage BizTalk services, but not access to them.
	BizTalkContributorRoleID string = "5e3c6656-6cfa-4708-81fe-0de47ac73342"

	// Allows for access to Blockchain Member nodes
	BlockchainMemberNodeAccessRoleID string = "31a002a1-acaf-453e-8a5b-297c9ca1ea24"

	// Can manage blueprint definitions, but not assign them.
	BlueprintContributorRoleID string = "41077137-e803-4205-871c-5a86e6a753b4"

	// Can assign existing published blueprints, but cannot create new blueprints. NOTE: this only works if the assignment is done with a user-assigned managed identity.
	BlueprintOperatorRoleID string = "437d2ced-4a38-4302-8479-ed2bcb43d090"

	// Can manage CDN endpoints, but can’t grant access to other users.
	CDNEndpointContributorRoleID string = "426e0c7f-0c7e-4658-b36f-ff54d6c29b45"

	// Can view CDN endpoints, but can’t make changes.
	CDNEndpointReaderRoleID string = "871e35f6-b5c1-49cc-a043-bde969a0f2cd"

	// Can manage CDN profiles and their endpoints, but can’t grant access to other users.
	CDNProfileContributorRoleID string = "ec156ff8-a8d1-4d15-830c-5b80698ca432"

	// Can view CDN profiles and their endpoints, but can’t make changes.
	CDNProfileReaderRoleID string = "8f96442b-4075-438f-813d-ad51ab4019af"

	// Lets you manage everything under your HPC Workbench chamber.
	ChamberAdminRoleID string = "4e9b8407-af2e-495b-ae54-bb60a55b1b5a"

	// Lets you view everything under your HPC Workbench chamber, but not make any changes.
	ChamberUserRoleID string = "4447db05-44ed-4da3-ae60-6cbece780e32"

	// Lets you manage classic networks, but not access to them.
	ClassicNetworkContributorRoleID string = "b34d265f-36f7-4a0d-a4d4-e158ca92e90f"

	// Lets you manage classic storage accounts, but not access to them.
	ClassicStorageAccountContributorRoleID string = "86e8f5dc-a6e9-4c67-9d15-de283e8eac25"

	// Classic Storage Account Key Operators are allowed to list and regenerate keys on Classic Storage Accounts
	ClassicStorageAccountKeyOperatorServiceRoleID string = "985d6b00-f706-48f5-a6fe-d0ca12fb668d"

	// Lets you manage classic virtual machines, but not access to them, and not the virtual network or storage account they’re connected to.
	ClassicVirtualMachineContributorRoleID string = "d73bb868-a0df-4d4d-bd69-98a00b01fccb"

	// Lets you manage ClearDB MySQL databases, but not access to them.
	ClearDBMySQLDBContributorRoleID string = "9106cda0-8a86-4e81-b686-29a22c54effe"

	// Manage identity or business verification requests. This role is in preview and subject to change.
	CodeSigningIdentityVerifierRoleID string = "4339b7cf-9826-4e41-b4ed-c7f4505dac08"

	// Sign files with a certificate profile. This role is in preview and subject to change.
	CodeSigningCertificateProfileSignerRoleID string = "2837e146-70d7-4cfd-ad55-7efa6464f958"

	// Lets you create, read, update, delete and manage keys of Cognitive Services.
	CognitiveServicesContributorRoleID string = "25fbc0a9-bd7c-42a3-aa1a-3b75d497ee68"

	// Full access to the project, including the ability to view, create, edit, or delete projects.
	CognitiveServicesCustomVisionContributorRoleID string = "c1ff6cc2-c111-46fe-8896-e0ef812ad9f3"

	// Publish, unpublish or export models. Deployment can view the project but can’t update.
	CognitiveServicesCustomVisionDeploymentRoleID string = "5c4089e1-6d96-4d2f-b296-c1bc7137275f"

	// View, edit training images and create, add, remove, or delete the image tags. Labelers can view the project but can’t update anything other than training images and tags.
	CognitiveServicesCustomVisionLabelerRoleID string = "88424f51-ebe7-446f-bc41-7fa16989e96c"

	// Read-only actions in the project. Readers can’t create or update the project.
	CognitiveServicesCustomVisionReaderRoleID string = "93586559-c37d-4a6b-ba08-b9f0940c2d73"

	// View, edit projects and train the models, including the ability to publish, unpublish, export the models. Trainers can’t create or delete the project.
	CognitiveServicesCustomVisionTrainerRoleID string = "0a5ae4ab-0d65-4eeb-be61-29fc9b54394b"

	// Lets you read Cognitive Services data.
	CognitiveServicesDataReaderRoleID string = "b59867f0-fa02-499b-be73-45a86b5b3e1c"

	// Lets you perform detect, verify, identify, group, and find similar operations on Face API. This role does not allow create or delete operations, which makes it well suited for endpoints that only need inferencing capabilities, following 'least privilege' best practices.
	CognitiveServicesFaceRecognizerRoleID string = "9894cab4-e18a-44aa-828b-cb588cd6f2d7"

	// Provides access to create Immersive Reader sessions and call APIs
	CognitiveServicesImmersiveReaderUserRoleID string = "b2de6794-95db-4659-8781-7e080d3f2b9d"

	//  Has access to all Read, Test, Write, Deploy and Delete functions under LUIS
	CognitiveServicesLUISOwnerRoleID string = "f72c8140-2111-481c-87ff-72b910f6e3f8"

	// Has access to Read and Test functions under LUIS.
	CognitiveServicesLUISReaderRoleID string = "18e81cdc-4e98-4e29-a639-e7d10c5a6226"

	// Has access to all Read, Test, and Write functions under LUIS
	CognitiveServicesLUISWriterRoleID string = "6322a993-d5c9-4bed-b113-e49bbea25b27"

	// Has access to all Read, Test, Write, Deploy and Delete functions under Language portal
	CognitiveServicesLanguageOwnerRoleID string = "f07febfe-79bc-46b1-8b37-790e26e6e498"

	// Has access to Read and Test functions under Language portal
	CognitiveServicesLanguageReaderRoleID string = "7628b7b8-a8b2-4cdc-b46f-e9b35248918e"

	//  Has access to all Read, Test, and Write functions under Language Portal
	CognitiveServicesLanguageWriterRoleID string = "f2310ca1-dc64-4889-bb49-c8e0fa3d47a8"

	// Full access to the project, including the system level configuration.
	CognitiveServicesMetricsAdvisorAdministratorRoleID string = "cb43c632-a144-4ec5-977c-e80c4affc34a"

	// Access to the project.
	CognitiveServicesMetricsAdvisorUserRoleID string = "3b20f47b-3825-43cb-8114-4bd2201156a8"

	// Let’s you create, edit, import and export a KB. You cannot publish or delete a KB.
	CognitiveServicesQnAMakerEditorRoleID string = "f4cc2bf9-21be-47a1-bdf1-5c5804381025"

	// Let’s you read and test a KB only.
	CognitiveServicesQnAMakerReaderRoleID string = "466ccd10-b268-4a11-b098-b4849f024126"

	// Full access to Speech projects, including read, write and delete all entities, for real-time speech recognition and batch transcription tasks, real-time speech synthesis and long audio tasks, custom speech and custom voice.
	CognitiveServicesSpeechContributorRoleID string = "0e75ca1e-0464-4b4d-8b93-68208a576181"

	// Access to the real-time speech recognition and batch transcription APIs, real-time speech synthesis and long audio APIs, as well as to read the data/test/model/endpoint for custom models, but can’t create, delete or modify the data/test/model/endpoint for custom models.
	CognitiveServicesSpeechUserRoleID string = "f2dc8367-1007-4938-bd23-fe263f013447"

	// Lets you read and list keys of Cognitive Services.
	CognitiveServicesUserRoleID string = "a97b65f3-24c7-4388-baec-2e87135dc908"

	// Can manage data packages of a collaborative.
	CollaborativeDataContributorRoleID string = "daa9e50b-21df-454c-94a6-a8050adab352"

	// Can manage resources created by AICS at runtime
	CollaborativeRuntimeOperatorRoleID string = "7a6f0e70-c033-4fb1-828c-08514e5f4102"

	// This role allows user to share gallery to another subscription/tenant or share it to the public.
	ComputeGallerySharingAdminRoleID string = "1ef6a3be-d0ac-425d-8c01-acb62866290b"

	// Grants full access to manage all resources, but does not allow you to assign roles in Azure RBAC, manage assignments in Azure Blueprints, or share image galleries.
	ContributorRoleID string = "b24988ac-6180-42a0-ab88-20f7382dd24c"

	// Can read Azure Cosmos DB Accounts data
	CosmosDBAccountReaderRoleID string = "fbdf93bf-df7d-467e-a4d2-9458aa1360c8"

	// Lets you manage Azure Cosmos DB accounts, but not access data in them. Prevents access to account keys and connection strings.
	CosmosDBOperatorRoleID string = "230815da-be43-4aae-9cb4-875f7bd000aa"

	// Can submit restore request for a Cosmos DB database or a container for an account
	CosmosBackupOperatorRoleID string = "db7b14f2-5adf-42da-9f96-f2ee17bab5cb"

	// Can perform restore action for Cosmos DB database account with continuous backup mode
	CosmosRestoreOperatorRoleID string = "5432c526-bc82-444a-b7ba-57c5b0b5b34f"

	// Can view costs and manage cost configuration (e.g. budgets, exports)
	CostManagementContributorRoleID string = "434105ed-43f6-45c7-a02f-909b2ba83430"

	// Can view cost data and configuration (e.g. budgets, exports)
	CostManagementReaderRoleID string = "72fafb9e-0641-4937-9268-a91bfd8191a3"

	// Full access to DICOM data.
	DICOMDataOwnerRoleID string = "58a3b984-7adf-4c20-983a-32417c86fbc8"

	// Read and search DICOM data.
	DICOMDataReaderRoleID string = "e89c7a3c-2f64-4fa1-a847-3e4c9ba4283a"

	// Lets you manage DNS resolver resources.
	DNSResolverContributorRoleID string = "0f2ebee7-ffd4-4fc0-b3b7-664099fdad5d"

	// Lets you manage DNS zones and record sets in Azure DNS, but does not let you control who has access to them.
	DNSZoneContributorRoleID string = "befefa01-2a29-4197-83a8-272ff33ce314"

	// Lets you manage everything under Data Box Service except giving access to others.
	DataBoxContributorRoleID string = "add466c9-e687-43fc-8d98-dfcf8d720be5"

	// Lets you manage Data Box Service except creating order or editing order details and giving access to others.
	DataBoxReaderRoleID string = "028f4ed7-e2a9-465e-a8f4-9c0ffdfdc027"

	// Create and manage data factories, as well as child resources within them.
	DataFactoryContributorRoleID string = "673868aa-7521-48a0-acc6-0f60742d39f5"

	// Lets you submit, monitor, and manage your own jobs but not create or delete Data Lake Analytics accounts.
	DataLakeAnalyticsDeveloperRoleID string = "47b7735b-770e-4598-a7da-8b91488b4c88"

	// Provides permissions to upload data to empty managed disks, read, or export data of managed disks (not attached to running VMs) and snapshots using SAS URIs and Azure AD authentication.
	DataOperatorforManagedDisksRoleID string = "959f8984-c045-4866-89c7-12bf9737be2e"

	// Can purge analytics data
	DataPurgerRoleID string = "150f5e0c-0603-4f03-8c7f-cf70034c4e90"

	// Contributor of the Desktop Virtualization Application Group.
	DesktopVirtualizationApplicationGroupContributorRoleID string = "86240b0e-9422-4c43-887b-b61143f32ba8"

	// Reader of the Desktop Virtualization Application Group.
	DesktopVirtualizationApplicationGroupReaderRoleID string = "aebf23d0-b568-4e86-b8f9-fe83a2c6ab55"

	// Contributor of Desktop Virtualization.
	DesktopVirtualizationContributorRoleID string = "082f0a83-3be5-4ba1-904c-961cca79b387"

	// Contributor of the Desktop Virtualization Host Pool.
	DesktopVirtualizationHostPoolContributorRoleID string = "e307426c-f9b6-4e81-87de-d99efb3c32bc"

	// Reader of the Desktop Virtualization Host Pool.
	DesktopVirtualizationHostPoolReaderRoleID string = "ceadfde2-b300-400a-ab7b-6143895aa822"

	// This role is in preview and subject to change. Provide permission to the Azure Virtual Desktop Resource Provider to start virtual machines.
	DesktopVirtualizationPowerOnContributorRoleID string = "489581de-a3bd-480d-9518-53dea7416b33"

	// This role is in preview and subject to change. Provide permission to the Azure Virtual Desktop Resource Provider to start and stop virtual machines.
	DesktopVirtualizationPowerOnOffContributorRoleID string = "40c5ff49-9181-41f8-ae61-143b0e78555e"

	// Reader of Desktop Virtualization.
	DesktopVirtualizationReaderRoleID string = "49a72310-ab8d-41df-bbb0-79b649203868"

	// Operator of the Desktop Virtualization Session Host.
	DesktopVirtualizationSessionHostOperatorRoleID string = "2ad6aaab-ead9-4eaa-8ac5-da422f562408"

	// Allows user to use the applications in an application group.
	DesktopVirtualizationUserRoleID string = "1d18fff3-a72a-46b5-b4a9-0b38a3cd7e63"

	// Operator of the Desktop Virtualization Uesr Session.
	DesktopVirtualizationUserSessionOperatorRoleID string = "ea4bfff8-7fb4-485a-aadd-d4129a0ffaa6"

	// This role is in preview and subject to change. Provide permission to the Azure Virtual Desktop Resource Provider to create, delete, update, start, and stop virtual machines.
	DesktopVirtualizationVirtualMachineContributorRoleID string = "a959dbd1-f747-45e3-8ba6-dd80f235f97c"

	// Contributor of the Desktop Virtualization Workspace.
	DesktopVirtualizationWorkspaceContributorRoleID string = "21efdde3-836f-432b-bf3d-3e8e734d4b2b"

	// Reader of the Desktop Virtualization Workspace.
	DesktopVirtualizationWorkspaceReaderRoleID string = "0fa44ee9-7a7d-466b-9bb2-2bf446b1204d"

	// Provides access to create and manage dev boxes.
	DevCenterDevBoxUserRoleID string = "45d50f46-0b78-4001-a660-4198cbe8cd05"

	// Provides access to manage project resources.
	DevCenterProjectAdminRoleID string = "331c37c6-af14-46d9-b9f4-e1909e1b95a0"

	// Lets you connect, start, restart, and shutdown your virtual machines in your Azure DevTest Labs.
	DevTestLabsUserRoleID string = "76283e04-6283-4c54-8f91-bcf1374a3c64"

	// Allows for full access to Device Provisioning Service data-plane operations.
	DeviceProvisioningServiceDataContributorRoleID string = "dfce44e4-17b7-4bd1-a6d1-04996ec95633"

	// Allows for full read access to Device Provisioning Service data-plane properties.
	DeviceProvisioningServiceDataReaderRoleID string = "10745317-c249-44a1-a5ce-3a4353c0bbd8"

	// Gives you full access to management and content operations
	DeviceUpdateAdministratorRoleID string = "02ca0879-e8e4-47a5-a61e-5c618b76e64a"

	// Gives you full access to content operations
	DeviceUpdateContentAdministratorRoleID string = "0378884a-3af5-44ab-8323-f5b22f9f3c98"

	// Gives you read access to content operations, but does not allow making changes
	DeviceUpdateContentReaderRoleID string = "d1ee9a80-8b14-47f0-bdc2-f4a351625a7b"

	// Gives you full access to management operations
	DeviceUpdateDeploymentsAdministratorRoleID string = "e4237640-0e3d-4a46-8fda-70bc94856432"

	// Gives you read access to management operations, but does not allow making changes
	DeviceUpdateDeploymentsReaderRoleID string = "49e2f5d2-7741-4835-8efa-19e1fe35e47f"

	// Gives you read access to management and content operations, but does not allow making changes
	DeviceUpdateReaderRoleID string = "e9dba6fb-3d52-4cf0-bce3-f06ce71b9e0f"

	// Provides permission to backup vault to perform disk backup.
	DiskBackupReaderRoleID string = "3e5e47e6-65f7-47ef-90b5-e5dd4d455f24"

	// Used by the StoragePool Resource Provider to manage Disks added to a Disk Pool.
	DiskPoolOperatorRoleID string = "60fc6e62-5479-42d4-8bf4-67625fcc2840"

	// Provides permission to backup vault to perform disk restore.
	DiskRestoreOperatorRoleID string = "b50d9833-a0cb-478e-945f-707fcc997c13"

	// Provides permission to backup vault to manage disk snapshots.
	DiskSnapshotContributorRoleID string = "7efff54f-a5b4-42b5-a1c5-5411624893ce"

	// Lets you manage DocumentDB accounts, but not access to them.
	DocumentDBAccountContributorRoleID string = "5bd9cd88-fe45-4216-938b-f97437e15450"

	// Can manage Azure AD Domain Services and related network configurations
	DomainServicesContributorRoleID string = "eeaeda52-9324-47f6-8069-5d5bade478b2"

	// Can view Azure AD Domain Services and related network configurations
	DomainServicesReaderRoleID string = "361898ef-9ed1-48c2-849c-a832951106bb"

	// Lets you manage elastic san accounts
	ElasticSanOwnerRoleID string = "80dcbedb-47ef-405d-95bd-188a1b4ac406"

	// Read Azure Elastic SAN and all sub-resources
	ElasticSanReaderRoleID string = "af6a70f8-3c9f-4105-acf1-d719e9fca4ca"

	// Lets you manage a volume group in elastic san account
	ElasticSanVolumeGroupOwnerRoleID string = "a8281131-f312-4f34-8d98-ae12be9f0d23"

	// Lets you manage EventGrid operations.
	EventGridContributorRoleID string = "1e241071-0855-49ea-94dc-649edcd759de"

	// Allows send access to event grid events.
	EventGridDataSenderRoleID string = "d5a91429-5739-47e2-a06b-3470a27159e7"

	// Lets you manage EventGrid event subscription operations.
	EventGridEventSubscriptionContributorRoleID string = "428e0ff0-5e57-4d9c-a221-2c70d0e0a443"

	// Lets you read EventGrid event subscriptions.
	EventGridEventSubscriptionReaderRoleID string = "2414bbcf-6497-4faf-8c65-045460748405"

	// Experimentation Administrator
	ExperimentationAdministratorRoleID string = "7f646f1b-fa08-80eb-a33b-edd6ce5c915c"

	// Experimentation Contributor
	ExperimentationContributorRoleID string = "7f646f1b-fa08-80eb-a22b-edd6ce5c915c"

	// Allows for creation, writes and reads to the metric set via the metrics service APIs.
	ExperimentationMetricContributorRoleID string = "6188b7c9-7d01-4f99-a59f-c88b630326c0"

	// Experimentation Reader
	ExperimentationReaderRoleID string = "49632ef5-d9ac-41f4-b8e7-bbe587fa74a1"

	// Role allows user or principal full access to FHIR Data
	FHIRDataContributorRoleID string = "5a1fc7df-4bf1-4951-a576-89034ee01acd"

	// Role allows user or principal to convert data from legacy format to FHIR
	FHIRDataConverterRoleID string = "a1705bd2-3a8f-45a5-8683-466fcfd5cc24"

	// Role allows user or principal to read and export FHIR Data
	FHIRDataExporterRoleID string = "3db33094-8700-4567-8da5-1501d4e7e843"

	// Role allows user or principal to read and import FHIR Data
	FHIRDataImporterRoleID string = "4465e953-8ced-4406-a58e-0f6e3f3b530b"

	// Role allows user or principal to read FHIR Data
	FHIRDataReaderRoleID string = "4c8d0bbc-75d3-4935-991f-5f3c56d81508"

	// Role allows user or principal to read and write FHIR Data
	FHIRDataWriterRoleID string = "3f88fce4-5892-4214-ae73-ba5294559913"

	// Built-in Grafana admin role
	GrafanaAdminRoleID string = "22926164-76b3-42b3-bc55-97df8dab3e41"

	// Built-in Grafana Editor role
	GrafanaEditorRoleID string = "a79a5197-3a5c-4973-a920-486035ffd60f"

	// Built-in Grafana Viewer role
	GrafanaViewerRoleID string = "60921a7e-fef1-4a43-9b16-a26c52ad4769"

	// Create and manage all aspects of the Enterprise Graph - Ontology, Schema mapping, Conflation and Conversational AI and Ingestions
	GraphOwnerRoleID string = "b60367af-1334-4454-b71e-769d9a4f83d9"

	// Lets you read, write Guest Configuration Resource.
	GuestConfigurationResourceContributorRoleID string = "088ab73d-1256-47ae-bea9-9de8e7131f31"

	// Lets you read and modify HDInsight cluster configurations.
	HDInsightClusterOperatorRoleID string = "61ed4efc-fab3-44fd-b111-e24485cc132a"

	// Can Read, Create, Modify and Delete Domain Services related operations needed for HDInsight Enterprise Security Package
	HDInsightDomainServicesContributorRoleID string = "8d8d5a11-05d3-4bda-a417-a08778121c7c"

	// Allows users to edit and delete Hierarchy Settings
	HierarchySettingsAdministratorRoleID string = "350f8d15-c687-4448-8ae1-157740a3936d"

	// Can onboard new Hybrid servers to the Hybrid Resource Provider.
	HybridServerOnboardingRoleID string = "5d1e5ee4-7c68-4a71-ac8b-0739630a3dfb"

	// Can read, write, delete, and re-onboard Hybrid servers to the Hybrid Resource Provider.
	HybridServerResourceAdministratorRoleID string = "48b40c6e-82e0-4eb3-90d5-19e40f49b624"

	// Lets you manage integration service environments, but not access to them.
	IntegrationServiceEnvironmentContributorRoleID string = "a41e2c5b-bd99-4a07-88f4-9bf657a760b8"

	// Allows developers to create and update workflows, integration accounts and API connections in integration service environments.
	IntegrationServiceEnvironmentDeveloperRoleID string = "c7aa55d3-1abb-444a-a5ca-5e51e485d6ec"

	// Lets you manage Intelligent Systems accounts, but not access to them.
	IntelligentSystemsAccountContributorRoleID string = "03a6d094-3444-4b3d-88af-7477090a9e5e"

	// Allows for full access to IoT Hub data plane operations.
	IoTHubDataContributorRoleID string = "4fc6c259-987e-4a07-842e-c321cc9d413f"

	// Allows for full read access to IoT Hub data-plane properties
	IoTHubDataReaderRoleID string = "b447c946-2db7-41ec-983d-d8bf3b1c77e3"

	// Allows for full access to IoT Hub device registry.
	IoTHubRegistryContributorRoleID string = "4ea46cd5-c1b2-4a8e-910b-273211f9ce47"

	// Allows for read and write access to all IoT Hub device and module twins.
	IoTHubTwinContributorRoleID string = "494bdba2-168f-4f31-a0a1-191d2f7c028c"

	// Perform all data plane operations on a key vault and all objects in it, including certificates, keys, and secrets. Cannot manage key vault resources or manage role assignments. Only works for key vaults that use the 'Azure role-based access control' permission model.
	KeyVaultAdministratorRoleID string = "00482a5a-887f-4fb3-b363-3b7fe8e74483"

	// Perform any action on the certificates of a key vault, except manage permissions. Only works for key vaults that use the 'Azure role-based access control' permission model.
	KeyVaultCertificatesOfficerRoleID string = "a4417e6f-fecd-4de8-b567-7b0420556985"

	// Lets you manage key vaults, but not access to them.
	KeyVaultContributorRoleID string = "f25e0fa2-a7c8-4377-a976-54943a77a395"

	// Perform any action on the keys of a key vault, except manage permissions. Only works for key vaults that use the 'Azure role-based access control' permission model.
	KeyVaultCryptoOfficerRoleID string = "14b46e9e-c2b7-41b4-b07b-48a6ebf60603"

	// Read metadata of keys and perform wrap/unwrap operations. Only works for key vaults that use the 'Azure role-based access control' permission model.
	KeyVaultCryptoServiceEncryptionUserRoleID string = "e147488a-f6f5-4113-8e2d-b22465e65bf6"

	// Perform cryptographic operations using keys. Only works for key vaults that use the 'Azure role-based access control' permission model.
	KeyVaultCryptoUserRoleID string = "12338af0-0e69-4776-bea7-57ae8d297424"

	// Read metadata of key vaults and its certificates, keys, and secrets. Cannot read sensitive values such as secret contents or key material. Only works for key vaults that use the 'Azure role-based access control' permission model.
	KeyVaultReaderRoleID string = "21090545-7ca7-4776-b22c-e363652d74d2"

	// Perform any action on the secrets of a key vault, except manage permissions. Only works for key vaults that use the 'Azure role-based access control' permission model.
	KeyVaultSecretsOfficerRoleID string = "b86a8fe4-44ce-4948-aee5-eccb2c155cd7"

	// Read secret contents. Only works for key vaults that use the 'Azure role-based access control' permission model.
	KeyVaultSecretsUserRoleID string = "4633458b-17de-408a-b874-0445c86b69e6"

	// Knowledge Read permission to consume Enterprise Graph Knowledge using entity search and graph query
	KnowledgeConsumerRoleID string = "ee361c5d-f7b5-4119-b4b6-892157c8f64c"

	// Role definition to authorize any user/service to create connectedClusters resource
	KubernetesClusterAzureArcOnboardingRoleID string = "34e09817-6cbe-4d01-b1a2-e0eac5743d41"

	// Can create, update, get, list and delete Kubernetes Extensions, and get extension async operations
	KubernetesExtensionContributorRoleID string = "85cb6faf-e071-4c9b-8136-154b5a04f717"

	// The lab assistant role
	LabAssistantRoleID string = "ce40b423-cede-4313-a93f-9b28290b72e1"

	// The lab contributor role
	LabContributorRoleID string = "5daaa2af-1fe8-407c-9122-bba179798270"

	// Lets you create new labs under your Azure Lab Accounts.
	LabCreatorRoleID string = "b97fb8bc-a8b2-4522-a38b-dd33c7e65ead"

	// The lab operator role
	LabOperatorRoleID string = "a36e6959-b6be-4b12-8e9f-ef4b474d304d"

	// The lab services contributor role
	LabServicesContributorRoleID string = "f69b8690-cc87-41d6-b77a-a4bc3c0a966f"

	// The lab services reader role
	LabServicesReaderRoleID string = "2a5c394f-5eb7-4d4f-9c8e-e8eae39faebc"

	// View, create, update, delete and execute load tests. View and list load test resources but can not make any changes.
	LoadTestContributorRoleID string = "749a398d-560b-491b-bb21-08924219302e"

	// Execute all operations on load test resources and load tests
	LoadTestOwnerRoleID string = "45bb0b16-2f0c-4e78-afaa-a07599b003f6"

	// View and list all load tests and load test resources but can not make any changes
	LoadTestReaderRoleID string = "3ae3fb29-0000-4ccd-bf80-542e7b26e081"

	// Log Analytics Contributor can read all monitoring data and edit monitoring settings. Editing monitoring settings includes adding the VM extension to VMs; reading storage account keys to be able to configure collection of logs from Azure Storage; adding solutions; and configuring Azure diagnostics on all Azure resources.
	LogAnalyticsContributorRoleID string = "92aaf0da-9dab-42b6-94a3-d43ce8d16293"

	// Log Analytics Reader can view and search all monitoring data as well as and view monitoring settings, including viewing the configuration of Azure diagnostics on all Azure resources.
	LogAnalyticsReaderRoleID string = "73c42c96-874c-492b-b04d-ab87d138a893"

	// Lets you manage logic app, but not access to them.
	LogicAppContributorRoleID string = "87a39d53-fc1b-424a-814c-f7e04687dc9e"

	// Lets you read, enable and disable logic app.
	LogicAppOperatorRoleID string = "515c2055-d9d4-4321-b1b9-bd0c9a0f79fe"

	// Allows for creating managed application resources.
	ManagedApplicationContributorRoleID string = "641177b8-a67a-45b9-a033-47bc880bb21e"

	// Lets you read and perform actions on Managed Application resources
	ManagedApplicationOperatorRoleID string = "c7393b34-138c-406f-901b-d8cf2b17e6ae"

	// Lets you read resources in a managed app and request JIT access.
	ManagedApplicationsReaderRoleID string = "b9331d33-8a36-4f8c-b097-4f54124fdb44"

	// Lets you manage managed HSM pools, but not access to them.
	ManagedHSMcontributorRoleID string = "18500a29-7fe2-46b2-a342-b16a415e101d"

	// Create, Read, Update, and Delete User Assigned Identity
	ManagedIdentityContributorRoleID string = "e40ec5ca-96e0-45a2-b4ff-59039f2c2b59"

	// Read and Assign User Assigned Identity
	ManagedIdentityOperatorRoleID string = "f1a07417-d97a-45cb-824c-7a7467783830"

	// Managed Services Registration Assignment Delete Role allows the managing tenant users to delete the registration assignment assigned to their tenant.
	ManagedServicesRegistrationassignmentDeleteRoleID string = "91c1777a-f3dc-4fae-b103-61d183457e46"

	// Management Group Contributor Role
	ManagementGroupContributorRoleID string = "5d58bcaf-24a5-4b20-bdb6-eed9f69fbe4c"

	// Management Group Reader Role
	ManagementGroupReaderRoleID string = "ac63b705-f282-497d-ac71-919bf39d939d"

	// Marketplace Admin grants full access to manage Private Azure Marketplace, including read and take action for private marketplace notifications, but does not allow to assign Marketplace Admin role to others
	MarketplaceAdminRoleID string = "dd920d6d-f481-47f1-b461-f338c46b2d9f"

	// Create, read, modify, and delete Media Services accounts; read-only access to other Media Services resources.
	MediaServicesAccountAdministratorRoleID string = "054126f8-9a2b-4f1c-a9ad-eca461f08466"

	// Create, read, modify, and delete Live Events, Assets, Asset Filters, and Streaming Locators; read-only access to other Media Services resources.
	MediaServicesLiveEventsAdministratorRoleID string = "532bc159-b25e-42c0-969e-a1d439f60d77"

	// Create, read, modify, and delete Assets, Asset Filters, Streaming Locators, and Jobs; read-only access to other Media Services resources.
	MediaServicesMediaOperatorRoleID string = "e4395492-1534-4db2-bedf-88c14621589c"

	// Create, read, modify, and delete Account Filters, Streaming Policies, Content Key Policies, and Transforms; read-only access to other Media Services resources. Cannot create Jobs, Assets or Streaming resources.
	MediaServicesPolicyAdministratorRoleID string = "c4bba371-dacd-4a26-b320-7250bca963ae"

	// Create, read, modify, and delete Streaming Endpoints; read-only access to other Media Services resources.
	MediaServicesStreamingEndpointsAdministratorRoleID string = "99dba123-b5fe-44d5-874c-ced7199a5804"

	// Microsoft Sentinel Automation Contributor
	MicrosoftSentinelAutomationContributorRoleID string = "f4c81013-99ee-4d62-a7ee-b3f1f648599a"

	// Microsoft Sentinel Contributor
	MicrosoftSentinelContributorRoleID string = "ab8e14d6-4a74-4a29-9ba8-549422addade"

	// Microsoft Sentinel Reader
	MicrosoftSentinelReaderRoleID string = "8d289c81-5878-46d4-8554-54e1e3d8b5cb"

	// Microsoft Sentinel Responder
	MicrosoftSentinelResponderRoleID string = "3e150937-b8fe-4cfb-8069-0eaf05ecd056"

	// Microsoft.Kubernetes connected cluster role.
	MicrosoftKubernetesConnectedClusterRoleID string = "5548b2cf-c94c-4228-90ba-30851930a12f"

	// Can read and update Monitored Objects and associated Data Collection Rules.
	MonitoredObjectsContributorRoleID string = "56be40e2-4db1-4ccf-93c3-7e44c597135b"

	// Can read all monitoring data and update monitoring settings.
	MonitoringContributorRoleID string = "749f88d5-cbae-40b8-bcfc-e573ddc772fa"

	// Enables publishing metrics against Azure resources
	MonitoringMetricsPublisherRoleID string = "3913510d-42f4-4e42-8a64-420c390055eb"

	// Can read all monitoring data.
	MonitoringReaderRoleID string = "43d0d8ad-25c7-4714-9337-8ba259a9fe05"

	// Lets you manage networks, but not access to them.
	NetworkContributorRoleID string = "4d97b98b-1d4f-4787-a291-c67834d212e7"

	// Lets you manage New Relic Application Performance Management accounts and applications, but not access to them.
	NewRelicAPMAccountContributorRoleID string = "5d28c62d-5b37-4476-8438-e587778df237"

	// Provides user with ingestion capabilities for an object anchors account.
	ObjectAnchorsAccountOwnerRoleID string = "ca0835dd-bacc-42dd-8ed2-ed5e7230d15b"

	// Lets you read ingestion jobs for an object anchors account.
	ObjectAnchorsAccountReaderRoleID string = "4a167cdf-cb95-4554-9203-2347fe489bd9"

	// Provides user with ingestion capabilities for Azure Object Understanding.
	ObjectUnderstandingAccountOwnerRoleID string = "4dd61c23-6743-42fe-a388-d8bdd41cb745"

	// Lets you read ingestion jobs for an object understanding account.
	ObjectUnderstandingAccountReaderRoleID string = "d18777c0-1514-4662-8490-608db7d334b6"

	// Grants full access to manage all resources, including the ability to assign roles in Azure RBAC.
	OwnerRoleID string = "8e3af657-a8ff-443c-a75c-2fe8c4bcb635"

	// Provides contributor access to PlayFab resources
	PlayFabContributorRoleID string = "0c8b84dc-067c-4039-9615-fa1a4b77c726"

	// Provides read access to PlayFab resources
	PlayFabReaderRoleID string = "a9a19cc5-31f4-447c-901f-56c0bb18fcaf"

	// Allows read access to resource policies and write access to resource component policy events.
	PolicyInsightsDataWriterRoleID string = "66bb4e9e-b016-4a94-8249-4c0511c2be84"

	// The user has access to perform administrative actions on all PowerApps resources within the tenant.
	PowerAppsAdministratorRoleID string = "53be45b2-ad40-43ab-bc1f-2c962ac99ded"

	// PowerAppsReadersWithReshare can use the resource and re-share it with other users, but cannot edit the resource or re-share it with edit permissions.
	PowerAppsReaderWithReshareRoleID string = "6877c72c-edd3-4048-9b4b-cf8e514477b0"

	// Lets you manage private DNS zone resources, but not the virtual networks they are linked to.
	PrivateDNSZoneContributorRoleID string = "b12aa53e-6015-4669-85d0-8515ebb3ae7f"

	// The Microsoft.ProjectBabylon data curator can create, read, modify and delete catalog data objects and establish relationships between objects. This role is in preview and subject to change.
	ProjectBabylonDataCuratorRoleID string = "9ef4ef9c-a049-46b0-82ab-dd8ac094c889"

	// The Microsoft.ProjectBabylon data reader can read catalog data objects. This role is in preview and subject to change.
	ProjectBabylonDataReaderRoleID string = "c8d896ba-346d-4f50-bc1d-7d1c84130446"

	// The Microsoft.ProjectBabylon data source administrator can manage data sources and data scans. This role is in preview and subject to change.
	ProjectBabylonDataSourceAdministratorRoleID string = "05b7651b-dc44-475e-b74d-df3db49fae0f"

	// Deprecated role.
	Purview1DeprecatedRoleID string = "8a3c2885-9b38-4fd2-9d99-91af537c1347"

	// Deprecated role.
	Purview2DeprecatedRoleID string = "200bba9e-f0c8-430f-892b-6f0794863803"

	// Deprecated role.
	Purview3DeprecatedRoleID string = "ff100721-1b9d-43d8-af52-42b69c1272db"

	// Read and create quota requests, get quota request status, and create support tickets.
	QuotaRequestOperatorRoleID string = "0e5f05e5-9ab9-446b-b98d-1e2157c94125"

	// View all resources, but does not allow you to make any changes.
	ReaderRoleID string = "acdd72a7-3385-48ef-bd42-f606fba81ae7"

	// Lets you view everything but will not let you delete or create a storage account or contained resource. It will also allow read/write access to all data contained in a storage account via access to storage account keys.
	ReaderandDataAccessRoleID string = "c12c1c16-33a1-487b-954d-41c89c60f349"

	// Lets you manage Redis caches, but not access to them.
	RedisCacheContributorRoleID string = "e0f68234-74aa-48ed-b826-c38b57376e17"

	// Provides user with conversion, manage session, rendering and diagnostics capabilities for Azure Remote Rendering
	RemoteRenderingAdministratorRoleID string = "3df8b902-2a6f-47c7-8cc5-360e9b272a7e"

	// Provides user with manage session, rendering and diagnostics capabilities for Azure Remote Rendering.
	RemoteRenderingClientRoleID string = "d39065c4-c120-43c9-ab0a-63eed9795f0a"

	// Lets you purchase reservations
	ReservationPurchaserRoleID string = "f7b75c60-3036-4b75-91c3-6b41c27c1689"

	// Lets one read and manage all the reservations in a tenant
	ReservationsAdministratorRoleID string = "a8889054-8d42-49c9-bc1c-52486c10e7cd"

	// Lets one read all the reservations in a tenant
	ReservationsReaderRoleID string = "582fc458-8989-419f-a480-75249bc5db7e"

	// Users with rights to create/modify resource policy, create support ticket and read resources/hierarchy.
	ResourcePolicyContributorRoleID string = "36243c78-bf99-498c-9df9-86d9f8d28608"

	// Lets you manage SQL databases, but not access to them. Also, you can't manage their security-related policies or their parent SQL servers.
	SQLDBContributorRoleID string = "9b7fa17d-e63e-47b0-bb0a-15c516ac86ec"

	// Lets you manage SQL Managed Instances and required network configuration, but can’t give access to others.
	SQLManagedInstanceContributorRoleID string = "4939a1f6-9ae0-4e48-a1e0-f2cbe897382d"

	// Lets you manage the security-related policies of SQL servers and databases, but not access to them.
	SQLSecurityManagerRoleID string = "056cd41c-7e88-42e1-933e-88ba6a50c9c3"

	// Lets you manage SQL servers and databases, but not access to them, and not their security -related policies.
	SQLServerContributorRoleID string = "6d8ee4ec-f05a-4a1d-8b00-a9b17e38b437"

	// Provides access to manage maintenance configurations with maintenance scope InGuestPatch and corresponding configuration assignments
	ScheduledPatchingContributorRoleID string = "cd08ab90-6b14-449c-ad9a-8f8e549482c6"

	// Lets you manage Scheduler job collections, but not access to them.
	SchedulerJobCollectionsContributorRoleID string = "188a0f2f-5c9e-469b-ae67-2aa5ce574b94"

	// Read, write, and delete Schema Registry groups and schemas.
	SchemaRegistryContributorRoleID string = "5dffeca3-4936-4216-b2bc-10343a5abb25"

	// Read and list Schema Registry groups and schemas.
	SchemaRegistryReaderRoleID string = "2c56ea50-c6b3-40a6-83c0-9d98858bc7d2"

	// Grants full access to Azure Cognitive Search index data.
	SearchIndexDataContributorRoleID string = "8ebe5a00-799e-43f5-93ac-243d3dce84a7"

	// Grants read access to Azure Cognitive Search index data.
	SearchIndexDataReaderRoleID string = "1407120a-92aa-4202-b7e9-c0e197c71c8f"

	// Lets you manage Search services, but not access to them.
	SearchServiceContributorRoleID string = "7ca78c08-252a-4471-8644-bb5ff32d4ba0"

	// Security Admin Role
	SecurityAdminRoleID string = "fb1c8493-542b-48eb-b624-b4c8fea62acd"

	// Lets you push assessments to Security Center
	SecurityAssessmentContributorRoleID string = "612c2aa1-cb24-443b-ac28-3ab7272de6f5"

	// Allowed to publish and modify platforms, workflows and toolsets to Security Detonation Chamber
	SecurityDetonationChamberPublisherRoleID string = "352470b3-6a9c-4686-b503-35deb827e500"

	// Allowed to query submission info and files from Security Detonation Chamber
	SecurityDetonationChamberReaderRoleID string = "28241645-39f8-410b-ad48-87863e2951d5"

	// Allowed to create and manage submissions to Security Detonation Chamber
	SecurityDetonationChamberSubmissionManagerRoleID string = "a37b566d-3efa-4beb-a2f2-698963fa42ce"

	// Allowed to create submissions to Security Detonation Chamber
	SecurityDetonationChamberSubmitterRoleID string = "0b555d9b-b4a7-4f43-b330-627f0e5be8f0"

	// This is a legacy role. Please use Security Administrator instead
	SecurityManagerLegacyRoleID string = "e3d13bf0-dd5a-482e-ba6b-9b8433878d10"

	// Security Reader Role
	SecurityReaderARMRoleID string = "39bc4728-0917-49c7-9d2c-d95423bc2eb4"

	// Services Hub Operator allows you to perform all read, write, and deletion operations related to Services Hub Connectors.
	ServicesHubOperatorRoleID string = "82200a5b-e217-47a5-b665-6d8765ee745b"

	// Read SignalR Service Access Keys
	SignalRAccessKeyReaderRoleID string = "04165923-9d83-45d5-8227-78b77b0a687e"

	// Lets your app server access SignalR Service with AAD auth options.
	SignalRAppServerRoleID string = "420fcaa2-552c-430f-98ca-3264be4806c7"

	// Full access to Azure SignalR Service REST APIs
	SignalRRESTAPIOwnerRoleID string = "fd53cd77-2268-407a-8f46-7e7863d0f521"

	// Read-only access to Azure SignalR Service REST APIs
	SignalRRESTAPIReaderRoleID string = "ddde6b66-c0df-4114-a159-3618637b3035"

	// Full access to Azure SignalR Service REST APIs
	SignalRServiceOwnerRoleID string = "7e4f1700-ea5a-4f59-8f37-079cfe29dce3"

	// Create, Read, Update, and Delete SignalR service resources
	SignalRWebPubSubContributorRoleID string = "8cf5e20a-e4b2-4e9d-b3a1-5ceb692c2761"

	// Lets you manage Site Recovery service except vault creation and role assignment
	SiteRecoveryContributorRoleID string = "6670b86e-a3f7-4917-ac9b-5d6ab1be4567"

	// Lets you failover and failback but not perform other Site Recovery management operations
	SiteRecoveryOperatorRoleID string = "494ae006-db33-4328-bf46-533a6560a3ca"

	// Lets you view Site Recovery status but not perform other management operations
	SiteRecoveryReaderRoleID string = "dbaa88c4-0c30-4179-9fb3-46319faa6149"

	// Lets you manage spatial anchors in your account, but not delete them
	SpatialAnchorsAccountContributorRoleID string = "8bbe83f1-e2a6-4df7-8cb4-4e04d4e5c827"

	// Lets you manage spatial anchors in your account, including deleting them
	SpatialAnchorsAccountOwnerRoleID string = "70bbe301-9835-447d-afdd-19eb3167307c"

	// Lets you locate and read properties of spatial anchors in your account
	SpatialAnchorsAccountReaderRoleID string = "5d51204f-eb77-4b1c-b86a-2ec626c49413"

	// Lets you perform backup and restore operations using Azure Backup on the storage account.
	StorageAccountBackupContributorRoleID string = "e5e2a7ff-d759-4cd2-bb51-3152d37e2eb1"

	// Lets you manage storage accounts, including accessing storage account keys which provide full access to storage account data.
	StorageAccountContributorRoleID string = "17d1049b-9a84-46fb-8f53-869881c3d3ab"

	// Storage Account Key Operators are allowed to list and regenerate keys on Storage Accounts
	StorageAccountKeyOperatorServiceRoleID string = "81a9662b-bebf-436f-a333-f67b29880f12"

	// Allows for read, write and delete access to Azure Storage blob containers and data
	StorageBlobDataContributorRoleID string = "ba92f5b4-2d11-453d-a403-e96b0029c9fe"

	// Allows for full access to Azure Storage blob containers and data, including assigning POSIX access control.
	StorageBlobDataOwnerRoleID string = "b7e6dc6d-f1e8-4753-8033-0f276bb0955b"

	// Allows for read access to Azure Storage blob containers and data
	StorageBlobDataReaderRoleID string = "2a2b9908-6ea1-4ae2-8e65-a410df84e7d1"

	// Allows for generation of a user delegation key which can be used to sign SAS tokens
	StorageBlobDelegatorRoleID string = "db58b8e5-c6ad-4a2a-8342-4190687cbf4a"

	// Allows for read, write, and delete access in Azure Storage file shares over SMB
	StorageFileDataSMBShareContributorRoleID string = "0c867c2a-1d8c-454a-a3db-ab2ea1bdc8bb"

	// Allows for read, write, delete and modify NTFS permission access in Azure Storage file shares over SMB
	StorageFileDataSMBShareElevatedContributorRoleID string = "a7264617-510b-434b-a828-9731dc254ea7"

	// Allows for read access to Azure File Share over SMB
	StorageFileDataSMBShareReaderRoleID string = "aba4ae5f-2193-4029-9191-0cb91df5e314"

	// Allows for read, write, and delete access to Azure Storage queues and queue messages
	StorageQueueDataContributorRoleID string = "974c5e8b-45b9-4653-ba55-5f855dd0fb88"

	// Allows for peek, receive, and delete access to Azure Storage queue messages
	StorageQueueDataMessageProcessorRoleID string = "8a0f0c08-91a1-4084-bc3d-661d67233fed"

	// Allows for sending of Azure Storage queue messages
	StorageQueueDataMessageSenderRoleID string = "c6a89b2d-59bc-44d0-9896-0f6e12d7b80a"

	// Allows for read access to Azure Storage queues and queue messages
	StorageQueueDataReaderRoleID string = "19e7f393-937e-4f77-808e-94535e297925"

	// Allows for read, write and delete access to Azure Storage tables and entities
	StorageTableDataContributorRoleID string = "0a9a7e1f-b9d0-4cc4-a60d-0319b160aaa3"

	// Allows for read access to Azure Storage tables and entities
	StorageTableDataReaderRoleID string = "76199698-9eea-4c19-bc75-cec21354c6b6"

	// Lets you perform query testing without creating a stream analytics job first
	StreamAnalyticsQueryTesterRoleID string = "1ec5b3c1-b17e-4e25-8312-2acb3c3c5abf"

	// Lets you create and manage Support requests
	SupportRequestContributorRoleID string = "cfd33db0-3dd1-45e3-aa9d-cdbdf3b6f24e"

	// Lets you manage tags on entities, without providing access to the entities themselves.
	TagContributorRoleID string = "4a9ae827-6dc8-4573-8ac7-8239d42aa03f"

	// Let you view and download packages and test results.
	TestBaseReaderRoleID string = "15e0f5a1-3450-4248-8e25-e2afe88a9e85"

	// Lets you manage Traffic Manager profiles, but does not let you control who has access to them.
	TrafficManagerContributorRoleID string = "a4b10055-b0c7-44c2-b00f-c7b5b3550cf7"

	// Lets you manage user access to Azure resources.
	UserAccessAdminRoleID string = "18d7d88d-d35e-4fb5-a5c3-7773c20a72d9"

	// Role that provides access to disk snapshot for security analysis.
	VMScannerOperatorRoleID string = "d24ecba3-c1f4-40fa-a7bb-4588a071e8fd"

	// Has access to view and search through all video's insights and transcription in the Video Indexer portal. No access to model customization, embedding of widget, downloading videos, or sharing the account.
	VideoIndexerRestrictedViewerRoleID string = "a2c4a527-7dc0-4ee3-897b-403ade70fafb"

	// View Virtual Machines in the portal and login as administrator
	VirtualMachineAdministratorLoginRoleID string = "1c0163c0-47e6-4577-8991-ea5c82e286e4"

	// Deprecated. Use VirtualMachineAdministratorLoginRoleID instead.
	AdminLoginRoleID string = "1c0163c0-47e6-4577-8991-ea5c82e286e4"

	// Lets you manage virtual machines, but not access to them, and not the virtual network or storage account they're connected to.
	VirtualMachineContributorRoleID string = "9980e02c-c2be-4d73-94e8-173b1dc7cf3c"

	// View Virtual Machines in the portal and login as a local user configured on the arc server
	VirtualMachineLocalUserLoginRoleID string = "602da2ba-a5c2-41da-b01d-5360126ab525"

	// View Virtual Machines in the portal and login as a regular user.
	VirtualMachineUserLoginRoleID string = "fb879df8-f326-4884-b1cf-06f3ad86be52"

	// Lets you manage the web plans for websites, but not access to them.
	WebPlanContributorRoleID string = "2cc479cb-7b4d-49a8-b449-8c00fd0f0a4b"

	// Full access to Azure Web PubSub Service REST APIs
	WebPubSubServiceOwnerRoleID string = "12cf5a90-567b-43ae-8102-96cf46c7d9b4"

	// Read-only access to Azure Web PubSub Service REST APIs
	WebPubSubServiceReaderRoleID string = "bfb1c7d2-fb1a-466b-b2ba-aee63b92deaf"

	// Lets you manage websites (not web plans), but not access to them.
	WebsiteContributorRoleID string = "de139f84-1756-47ae-9be6-808fbbe84772"

	// Let's you manage the OS of your resource via Windows Admin Center as an administrator.
	WindowsAdminCenterAdministratorLoginRoleID string = "a6333a3e-0164-44c3-b281-7a577aff287f"

	// Can save shared workbooks.
	WorkbookContributorRoleID string = "e8ddcd69-c73f-4f9f-9844-4100522f16ad"

	// Can read workbooks.
	WorkbookReaderRoleID string = "b279062a-9be3-42a0-92ae-8b3cf002ec4d"

	// WorkloadBuilder Migration Agent Role.
	WorkloadBuilderMigrationAgentRoleID string = "d17ce0a2-0697-43bc-aac5-9113337ab61c"
)
