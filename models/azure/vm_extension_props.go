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

type VMExtensionProperties struct {
	// Indicates whether the extension should use a newer minor version if one is available at deployment time.
	// Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property
	// set to true.
	AutoUpgradeMinorVersion bool `json:"autoUpgradeMinorVersion,omitempty"`

	// Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of
	// the extension available.
	EnabledAutomaticUpgrade bool `json:"enabledAutomaticUpgrade,omitempty"`

	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag string `json:"forceUpdateTag,omitempty"`

	// The virtual machine extension instance view.
	InstanceView VirtualMachineExtensionInstanceView `json:"instanceView,omitempty"`

	// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at
	// all.
	ProtectedSettings map[string]string `json:"protectedSettings,omitempty"`

	// The provisioning state, which only appears in the response.
	ProvisioningState string `json:"provisioningState,omitempty"`

	// The name of the extension handler publisher.
	Publisher string `json:"publisher,omitempty"`

	// Json formatted public settings for the extension.
	Settings map[string]string `json:"settings,omitempty"`

	// Indicates whether failures stemming from the extension will be suppressed (Operational failures such as not
	// connecting to the VM will not be suppressed regardless of this value).
	// The default is false.
	SuppressFailures bool `json:"suppressFailures,omitempty"`

	// Specifies the type of the extension; an example is "CustomScriptExtension".
	Type string `json:"type,omitempty"`

	// Specifies the version of the script handler.
	TypeHandlerVersion string `json:"typeHandlerVersion,omitempty"`
}
