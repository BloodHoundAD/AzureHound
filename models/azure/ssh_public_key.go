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

// Contains information about SSH certificate public key and the path on the Linux VM where the public key is placed.
type SshPublicKey struct {
	// SSH public key certificate used to authenticate with the VM through ssh.
	// The key needs to be at least 2048-bit and in ssh-rsa format.
	// For creating ssh keys, see [Create SSH keys on Linux and Mac for Linux VMs in Azure](https://docs.microsoft.com/azure/virtual-machines/linux/create-ssh-keys-detailed).
	KeyData string `json:"keyData,omitempty"`

	// Specifies the full path on the created VM where ssh public key is stored.
	// If the file already exists, the specified key is appended to the file. Example: /home/user/.ssh/authorized_keys
	Path string `json:"path,omitempty"`
}
