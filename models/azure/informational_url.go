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

// Basic profile information of the application.
// For more detail see https://docs.microsoft.com/en-us/graph/api/resources/informationalurl?view=graph-rest-1.0
type InformationalUrl struct {
	// CDN URL to the application's logo, Read-only.
	LogoUrl string `json:"logoUrl,omitempty"`

	// Link to the application's marketing page. For example, https://www.contoso.com/app/marketing
	MarketingUrl string `json:"marketingUrl,omitempty"`

	// Link to the application's privacy statement. For example, https://www.contoso.com/app/privacy
	PrivacyStatementUrl string `json:"privacyStatementUrl,omitempty"`

	// Link to the application's support page. For example, https://www.contoso.com/app/support
	SupportUrl string `json:"supportUrl,omitempty"`

	// Link to the application's terms of service statement. For example, https://www.contoso.com/app/termsofservice
	TermsOfServiceUrl string `json:"termsOfServiceUrl,omitempty"`
}
