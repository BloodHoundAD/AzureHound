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

// Represents an application. Any application that outsources authentication to Azure Active Directory (Azure AD) must
// be registered in a directory. Application registration involves telling Azure AD about your application, including
// the URL where it's located, the URL to send replies after authentication, the URI to identify your application, and
// more.
// For more detail see https://docs.microsoft.com/en-us/graph/api/resources/application?view=graph-rest-1.0
type Application struct {
	DirectoryObject

	// Defines custom behavior that a consuming service can use to call an app in specific contexts.
	// For example, applications that can render file streams may set the addIns property for its "FileHandler"
	// functionality. This will let services like Office 365 call the application in the context of a document the user
	// is working on.
	AddIns []AddIn `json:"addIns,omitempty"`

	// Specifies settings for an application that implements a web API.
	Api ApiApplication `json:"api,omitempty"`

	// The unique identifier for the application that is assigned to an application by Azure AD. Not nullable. Read-only.
	AppId string `json:"appId,omitempty"`

	// Unique identifier of the applicationTemplate.
	ApplicationTemplateId string `json:"applicationTemplateId,omitempty"`

	// The collection of roles assigned to the application.
	// With app role assignments, these roles can be assigned to users, groups, or service principals associated with
	// other applications. Not nullable.
	AppRoles []AppRole `json:"appRoles,omitempty"`

	// The date and time the application was registered.
	// The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time.
	// For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	//
	// Supports $filter (eq, ne, NOT, ge, le, in) and $orderBy.
	CreatedDateTime string `json:"createdDateTime,omitempty"`

	// The date and time the application was deleted. The DateTimeOffset type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	// Read-only.
	DeletedDateTime string `json:"deletedDateTime,omitempty"`

	// An optional description of the application.
	// Supports $filter (eq, ne, NOT, ge, le, startsWith) and $search.
	Description string `json:"description,omitempty"`

	// Specifies whether Microsoft has disabled the registered application.
	// Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement
	// (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services
	// Agreement).
	DisabledByMicrosoftStatus string `json:"disabledByMicrosoftStatus,omitempty"`

	// The display name for the application.
	// Supports $filter (eq, ne, NOT, ge, le, in, startsWith), $search, and $orderBy.
	DisplayName string `json:"displayName,omitempty"`

	// Configures the groups claim issued in a user or OAuth 2.0 access token that the application expects.
	// To set this attribute, use one of the following valid string values:
	// - None
	// - SecurityGroup (for security groups and Azure AD roles)
	// - All (this gets all of the security groups, distribution groups, and Azure AD directory roles that the signed-in user is a member of)
	GroupMembershipClaims string `json:"groupMembershipClaims,omitempty"`

	// The URIs that identify the application within its Azure AD tenant, or within a verified custom domain if the
	// application is multi-tenant. For more information see Application Objects and Service Principal Objects.
	// The any operator is required for filter expressions on multi-valued properties.
	// Not nullable.
	// Supports $filter (eq, ne, ge, le, startsWith).
	IdentifierUris []string `json:"identifierUris,omitempty"`

	// Basic profile information of the application such as app's marketing, support, terms of service and privacy
	// statement URLs. The terms of service and privacy statement are surfaced to users through the user consent
	// experience. For more info, see How to: Add Terms of service and privacy statement for registered Azure AD apps.
	// Supports $filter (eq, ne, NOT, ge, le).
	Info InformationalUrl `json:"info,omitempty"`

	// Specifies whether this application supports device authentication without a user.
	// The default is false.
	IsDeviceOnlyAuthSupported bool `json:"isDeviceOnlyAuthSupported,omitempty"`

	// Specifies the fallback application type as public client, such as an installed application running on a mobile
	// device.
	// The default value is false which means the fallback application type is confidential client such as a web app.
	// There are certain scenarios where Azure AD cannot determine the client application type. For example, the ROPC
	// flow where it is configured without specifying a redirect URI. In those cases Azure AD interprets the application
	// type based on the value of this property.
	IsFallbackPublicClient bool `json:"isFallbackPublicClient,omitempty"`

	// The collection of key credentials associated with the application.
	// Not nullable.
	// Supports $filter (eq, NOT, ge, le).
	KeyCredentials []KeyCredential `json:"keyCredentials,omitempty"`

	// The main logo for the application. Not nullable.
	// Base64Url encoded.
	Logo string `json:"logo,omitempty"`

	// Notes relevant for the management of the application.
	Notes string `json:"notes,omitempty"`

	// Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests.
	// The default is false, which specifies that only GET requests are allowed.
	OAuth2RequiredPostResponse bool `json:"oauth2RequiredPostResponse,omitempty"`

	// Application developers can configure optional claims in their Azure AD applications to specify the claims that
	// are sent to their application by the Microsoft security token service.
	// For more information, see How to: Provide optional claims to your app.
	OptionalClaims OptionalClaims `json:"optionalClaims,omitempty"`

	// Specifies parental control settings for an application.
	ParentalControlSettings ParentalControlSettings `json:"parentalControlSettings,omitempty"`

	// The collection of password credentials associated with the application. Not nullable.
	PasswordCredentials []PasswordCredential `json:"passwordCredentials,omitempty"`

	// Specifies settings for installed clients such as desktop or mobile devices.
	PublicClient PublicClientApplication `json:"publicClient,omitempty"`

	// The verified publisher domain for the application. Read-only.
	// For more information, see How to: Configure an application's publisher domain.
	// Supports $filter (eq, ne, ge, le, startsWith).
	PublisherDomain string `json:"publisherDomain,omitempty"`

	// Specifies the resources that the application needs to access.
	// This property also specifies the set of delegated permissions and application roles that it needs for each of
	// those resources. This configuration of access to the required resources drives the consent experience. No more
	// than 50 resource services (APIs) can be configured. Beginning mid-October 2021, the total number of required
	// permissions must not exceed 400. Not nullable.
	RequiredResourceAccess []RequiredResourceAccess `json:"requiredResourceAccess,omitempty"`

	// Specifies the Microsoft accounts that are supported for the current application.
	// The possible values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount (default), and
	// PersonalMicrosoftAccount. See more in the table below.
	SignInAudience string `json:"signInAudience,omitempty"`

	// Specifies settings for a single-page application, including sign out URLs and redirect URIs for authorization
	// codes and access tokens.
	SPA SPAApplication `json:"spa,omitempty"`

	// Custom strings that can be used to categorize and identify the application. Not nullable.
	Tags []string `json:"tags,omitempty"`

	// Specifies the keyId of a public key from the keyCredentials collection.
	// When configured, Azure AD encrypts all the tokens it emits by using the key this property points to. The
	// application code that receives the encrypted token must use the matching private key to decrypt the token before
	// it can be used for the signed-in user.
	TokenEncryptionKeyId string `json:"tokenEncryptionKeyId,omitempty"`

	// Specifies the verified publisher of the application.
	VerifiedPublisher VerifiedPublisher `json:"verifiedPublisher,omitempty"`

	// Specifies settings for a web application.
	Web WebApplication `json:"web,omitempty"`
}
