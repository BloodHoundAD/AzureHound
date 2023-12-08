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

// Specifies additional XML formatted information that can be included in the Unattend.xml file, which is used by
// Windows Setup. Contents are defined by setting name, component name, and the pass in which the content is applied.
type AdditionalUnattendContent struct {
	// The component name. Currently, the only allowable value is Microsoft-Windows-Shell-Setup.
	ComponentName string `json:"componentName,omitempty"`

	// Specifies the XML formatted content that is added to the unattend.xml file for the specified path and component.
	// The XML must be less than 4KB and must include the root element for the setting or feature that is being inserted.
	Content string `json:"content,omitempty"`

	// The pass name. Currently, the only allowable value is OobeSystem.
	PassName string `json:"passName,omitempty"`

	// Specifies the name of the setting to which the content applies.
	// Possible values are: FirstLogonCommands and AutoLogon.
	SettingName string `json:"settingName,omitempty"`
}
