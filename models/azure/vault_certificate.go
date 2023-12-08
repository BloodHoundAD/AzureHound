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

// Describes a single certificate reference in a Key Vault, and where the certificate should reside on the VM.
type VaultCertificate struct {
	// For Windows VMs, specifies the certificate store on the Virtual Machine to which the certificate should be added.
	// The specified certificate store is implicitly in the LocalMachine account.
	// For Linux VMs, the certificate file is placed under the /var/lib/waagent directory, with the file name
	// <UppercaseThumbprint>.crt for the X509 certificate file and <UppercaseThumbprint>.prv for private key.
	// Both of these files are .pem formatted.
	CertificateStore string `json:"certificateStore,omitempty"`

	// This is the URL of a certificate that has been uploaded to Key Vault as a secret.
	// For adding a secret to the Key Vault, see Add a key or secret to the key vault. In this case, your certificate
	// needs to be It is the Base64 encoding of the following JSON Object which is encoded in UTF-8:
	//
	// ```json
	// {
	//   "data":"",
	//   "dataType":"pfx",
	//   "password":""
	// }
	// ```
	//
	// To install certificates on a virtual machine it is recommended to use the Azure Key Vault virtual machine
	// extension for Linux or the Azure Key Vault virtual machine extension for Windows.
	CertificateUrl string `json:"certificateUrl,omitempty"`
}
