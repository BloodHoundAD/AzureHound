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

// Specifies the operating system settings for the virtual machine. Some of the settings cannot be changed once VM is
// provisioned.
type OSProfile struct {
	// Specifies the password of the administrator account.
	// Minimum-length (Windows): 8 characters
	// Minimum-length (Linux): 6 characters
	// Max-length (Windows): 123 characters
	// Max-length (Linux): 72 characters
	// Complexity requirements: 3 out of 4 conditions below need to be fulfilled
	// Has lower characters
	// Has upper characters
	// Has a digit
	// Has a special character (Regex match [\W_])
	// Disallowed values: "abc@123", "P@$$w0rd", "P@ssw0rd", "P@ssword123", "Pa$$word", "pass@word1", "Password!", "Password1", "Password22", "iloveyou!"
	// For resetting the password, see How to reset the Remote Desktop service or its login password in a Windows VM
	// For resetting root password, see Manage users, SSH, and check or repair disks on Azure Linux VMs using the VMAccess Extension
	AdminPassword string `json:"adminPassword,omitempty,omitempty"`

	// Specifies the name of the administrator account.
	// This property cannot be updated after the VM is created.
	// Windows-only restriction: Cannot end in "."
	// Disallowed values: "administrator", "admin", "user", "user1", "test", "user2", "test1", "user3", "admin1", "1", "123", "a", "actuser", "adm", "admin2", "aspnet", "backup", "console", "david", "guest", "john", "owner", "root", "server", "sql", "support", "support_388945a0", "sys", "test2", "test3", "user4", "user5".
	// Minimum-length (Linux): 1 character
	// Max-length (Linux): 64 characters
	// Max-length (Windows): 20 characters.
	AdminUsername string `json:"adminUsername,omitempty,omitempty"`

	// Specifies whether extension operations should be allowed on the virtual machine.
	// This may only be set to False when no extensions are present on the virtual machine.
	AllowExtensionOperations bool `json:"allowExtensionOperations,omitempty,omitempty"`

	// Specifies the host OS name of the virtual machine.
	// This name cannot be updated after the VM is created.
	// Max-length (Windows): 15 characters
	// Max-length (Linux): 64 characters.
	// For naming conventions and restrictions see Azure infrastructure services implementation guidelines.
	ComputerName string `json:"computerName,omitempty,omitempty"`

	// Specifies a base-64 encoded string of custom data. The base-64 encoded string is decoded to a binary array that is saved as a file on the Virtual Machine. The maximum length of the binary array is 65535 bytes.
	// Note: Do not pass any secrets or passwords in customData property
	// This property cannot be updated after the VM is created.
	// customData is passed to the VM to be saved as a file, for more information see Custom Data on Azure VMs
	// For using cloud-init for your Linux VM, see Using cloud-init to customize a Linux VM during creation
	CustomData string `json:"customData,omitempty,omitempty"`

	// Specifies the Linux operating system settings on the virtual machine.
	// For a list of supported Linux distributions, see Linux on Azure-Endorsed Distributions.
	LinuxConfiguration LinuxConfiguration `json:"linuxConfiguration,omitempty,omitempty"`

	// Specifies whether the guest provision signal is required to infer provision success of the virtual machine.
	// Note: This property is for private testing only, and all customers must not set the property to false.
	RequireGuestProvisionSignal bool `json:"requireGuestProvisionSignal,omitempty,omitempty"`

	// Specifies set of certificates that should be installed onto the virtual machine. To install certificates on a
	// virtual machine it is recommended to use the Azure Key Vault virtual machine extension for Linux or the
	// Azure Key Vault virtual machine extension for Windows.
	Secrets []VaultSecretGroup `json:"secrets,omitempty,omitempty"`

	// Specifies Windows operating system settings on the virtual machine.
	WindowsConfiguration WindowsConfiguration `json:"windowsConfiguration,omitempty,omitempty"`
}
