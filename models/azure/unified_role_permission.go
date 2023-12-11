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

// Represents a collection of allowed resource actions and the conditions that must be met for the action to be allowed.
// Resource actions are tasks that can be performed on a resource. For example, an application resource may support
// create, update, delete, and reset password actions.
type UnifiedRolePermission struct {
	// Set of tasks that can be performed on a resource.
	// Required.
	//
	// The following is the schema for resource actions:
	// <Namespace>/<Entity>/<PropertySet>/<Action>
	//
	// For example: microsoft.directory/applications/credentials/update.
	//
	// * Namespace - The services that exposes the task. For example, all tasks in Azure Active Directory use the namespace microsoft.directory.
	// * Entity - The logical features or components exposed by the service in Microsoft Graph. For example, applications, servicePrincipals, or groups.
	// * PropertySet - The specific properties or aspects of the entity for which access is being granted. For example, microsoft.directory/applications/authentication/read grants the ability to read the reply URL, logout URL, and implicit flow property on the application object in Azure AD. The following are reserved names for common property sets:
	//		* allProperties - Designates all properties of the entity, including privileged properties. Examples include microsoft.directory/applications/allProperties/read and microsoft.directory/applications/allProperties/update.
	//		* basic - Designates common read properties but excludes privileged ones. For example, microsoft.directory/applications/basic/update includes the ability to update standard properties like display name.
	//		* standard - Designates common update properties but excludes privileged ones. For example, microsoft.directory/applications/standard/read.
	// * Actions - The operations being granted. In most circumstances, permissions should be expressed in terms of CRUD or allTasks. Actions include:
	//		* Create - The ability to create a new instance of the entity.
	//		* Read - The ability to read a given property set (including allProperties).
	//		* Update - The ability to update a given property set (including allProperties).
	//		* Delete - The ability to delete a given entity.
	//		* AllTasks - Represents all CRUD operations (create, read, update, and delete).
	AllowedResourceActions []string `json:"allowedResourceActions,omitempty"`

	// Optional constraints that must be met for the permission to be effective.
	//
	// Conditions define constraints that must be met. For example, a requirement that the principal be an owner of the
	// target resource. The following are the supported conditions:
	//
	// Self: "@Subject.objectId == @Resource.objectId"
	// Owner: "@Subject.objectId Any_of @Resource.owners"
	//
	// The following is an example of a role permission with a condition that the principal be the owner of the target
	// resource:
	//
	//	"rolePermissions": [
	// 	        {
	// 	            "allowedResourceActions": [
	// 	                "microsoft.directory/applications/basic/update",
	// 	                "microsoft.directory/applications/credentials/update"
	// 	            ],
	// 	            "condition":  "@Subject.objectId Any_of @Resource.owners"
	// 	        }
	// 	]
	Condition string `json:"condition,omitempty"`

	// Set of tasks tat may not be performed on a resource. Not yet supported.
	// See AllowedResourceActions for more information.
	ExcludedResourceActions []string `json:"excludedResourceActions,omitempty"`
}
