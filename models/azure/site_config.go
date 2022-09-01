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

type SiteConfig struct {
	NumberOfWorkers int `json:"numberOfWorkers,omitempty"`
	// DefaultDocuments string                       :
	NetFrameworkVersion string `json:"netFrameworkVersion,omitemtpy"`
	PhpVersion          string `json:"phpVersion,omitemtpy"`
	PythonVersion       string `json:"pythonVersion,omitemtpy"`
	NodeVersion         string `json:"nodeVersion,omitemtpy"`
	PowerShellVersion   string `json:"powerShellVersion,omitemtpy"`
	LinuxFxVersion      string `json:"linuxFxVersion,omitemtpy"`
	WindowsFxVersion    string `json:"windowsFxVersion,omitemtpy"`
	// RequestTracingEnabled                  string `json:",omitemtpy"`
	// RemoteDebuggingEnabled                 string `json:",omitemtpy"`
	// RemoteDebuggingVersion                 string `json:",omitemtpy"`
	// HttpLoggingEnabled                     string `json:",omitemtpy"`
	// AzureMonitorLogCategories              string `json:",omitemtpy"`
	AcrUseManagedIdentityCreds bool `json:"acrUseManagedIdentityCreds,omitemtpy"`
	AcrUserManagedIdentityID   bool `json:"acrUserManagedIdentityID,omitemtpy"`
	// LogsDirectorySizeLimit                 string `json:",omitemtpy"`
	// DetailedErrorLoggingEnabled            string `json:",omitemtpy"`
	PublishingUsername string `json:"publishingUsername,omitemtpy"`
	PublishingPassword string `json:"publishingPassword,omitemtpy"`
	// AppSettings                            string `json:",omitemtpy"`
	// Metadata                               string `json:",omitemtpy"`
	ConnectionStrings string `json:"connectionStrings,omitemtpy"`
	// MachineKey                             string `json:",omitemtpy"`
	// HandlerMappings                        string `json:",omitemtpy"`
	// DocumentRoot                           string `json:",omitemtpy"`
	// ScmType                                string `json:",omitemtpy"`
	// Use32BitWorkerProcess                  string `json:",omitemtpy"`
	// WebSocketsEnabled                      string `json:",omitemtpy"`
	// AlwaysOn                               string `json:",omitemtpy"`
	JavaVersion string `json:"javaVersion,omitemtpy"`
	// JavaContainer                          string `json:",omitemtpy"`
	// JavaContainerVersion                   string `json:",omitemtpy"`
	AppCommandLine string `json:"appCommandLine,omitemtpy"`
	// ManagedPipelineMode                    string `json:",omitemtpy"`
	// VirtualApplications                    string `json:",omitemtpy"`
	// WinAuthAdminState                      string `json:",omitemtpy"`
	// WinAuthTenantState                     string `json:",omitemtpy"`
	// CustomAppPoolIdentityAdminState        string `json:",omitemtpy"`
	// CustomAppPoolIdentityTenantState       string `json:",omitemtpy"`
	RuntimeADUser         string `json:"runtimeADUser,omitemtpy"`
	RuntimeADUserPassword string `json:"runtimeADUserPassword,omitemtpy"`
	// LoadBalancing                          string `json:",omitemtpy"`
	// RoutingRules                           string `json:",omitemtpy"`
	// Experiments                            string `json:",omitemtpy"`
	// Limits                                 string `json:",omitemtpy"`
	// AutoHealEnabled                        string `json:",omitemtpy"`
	// AutoHealRules                          string `json:",omitemtpy"`
	// TracingOptions                         string `json:",omitemtpy"`
	// VnetName                               string `json:",omitemtpy"`
	// VnetRouteAllEnabled                    string `json:",omitemtpy"`
	// VnetPrivatePortsCount                  string `json:",omitemtpy"`
	// PublicNetworkAccess                    string `json:",omitemtpy"`
	// Cors                                   string `json:",omitemtpy"`
	// Push                                   string `json:",omitemtpy"`
	// ApiDefinition                          string `json:",omitemtpy"`
	// ApiManagementConfig                    string `json:",omitemtpy"`
	// AutoSwapSlotName                       string `json:",omitemtpy"`
	// LocalMySqlEnabled                      string `json:",omitemtpy"`
	// ManagedServiceIdentityId               string `json:",omitemtpy"`
	// XManagedServiceIdentityId              string `json:",omitemtpy"`
	// KeyVaultReferenceIdentity              string `json:",omitemtpy"`
	// IpSecurityRestrictions                 string `json:",omitemtpy"`
	// IpSecurityRestrictionsDefaultAction    string `json:",omitemtpy"`
	// ScmIpSecurityRestrictions              string `json:",omitemtpy"`
	// ScmIpSecurityRestrictionsDefaultAction string `json:",omitemtpy"`
	// ScmIpSecurityRestrictionsUseMain       string `json:",omitemtpy"`
	// Http20Enabled                          string `json:",omitemtpy"`
	// MinTlsVersion                          string `json:",omitemtpy"`
	// MinTlsCipherSuite                      string `json:",omitemtpy"`
	// SupportedTlsCipherSuites               string `json:",omitemtpy"`
	// ScmMinTlsVersion                       string `json:",omitemtpy"`
	// FtpsState                              string `json:",omitemtpy"`
	// PreWarmedInstanceCount                 string `json:",omitemtpy"`
	// FunctionAppScaleLimit                  string `json:",omitemtpy"`
	// ElasticWebAppScaleLimit                string `json:",omitemtpy"`
	// HealthCheckPath                        string `json:",omitemtpy"`
	// FileChangeAuditEnabled                 string `json:",omitemtpy"`
	// FunctionsRuntimeScaleMonitoringEnabled string `json:",omitemtpy"`
	// WebsiteTimeZone                        string `json:",omitemtpy"`
	// MinimumElasticInstanceCount            string `json:",omitemtpy"`
	// AzureStorageAccounts                   string `json:",omitemtpy"`
	// Http20ProxyFlag                        string `json:",omitemtpy"`
	// SitePort                               string `json:",omitemtpy"`
	// AntivirusScanEnabled                   string `json:",omitemtpy"`
	// StorageType                            string `json:",omitemtpy"`
}
