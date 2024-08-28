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

package query

import (
	"strconv"
	"strings"
)

const (
	ApiVersion                 string = "api-version"
	Count                      string = "$count"
	Expand                     string = "$expand"
	Filter                     string = "$filter"
	Format                     string = "$format"
	IncludeDeleted             string = "$include"
	IncludeAllTenantCategories string = "$includeAllTenantCategories"
	MaxPageSize                string = "$maxpagesize"
	OrderBy                    string = "$orderby"
	Recurse                    string = "$recurse"
	Search                     string = "$search"
	Select                     string = "$select"
	Skip                       string = "$skip"
	SkipToken                  string = "$skipToken"
	StatusOnly                 string = "StatusOnly"
	TenantId                   string = "tenantId"
	Top                        string = "$top"
)

type Params interface {
	AsMap() map[string]string
	NeedsEventualConsistencyHeaderFlag() bool
}

type RMParams struct {
	ApiVersion                 string
	Expand                     string
	Filter                     string
	IncludeDeleted             string
	IncludeAllTenantCategories bool
	MaxPageSize                string
	Recurse                    bool
	SkipToken                  string
	StatusOnly                 bool
	TenantId                   string // For cross-tenant request
	Top                        int32
}

func (s RMParams) NeedsEventualConsistencyHeaderFlag() bool {
	return false
}

func (s RMParams) AsMap() map[string]string {
	params := make(map[string]string)

	if s.ApiVersion != "" {
		params[ApiVersion] = s.ApiVersion
	}

	if s.Expand != "" {
		params[Expand] = s.Expand
	}

	if s.Filter != "" {
		params[Filter] = s.Filter
	}

	if s.IncludeAllTenantCategories {
		params[IncludeAllTenantCategories] = "true"
	}

	if s.Recurse {
		params[Recurse] = "true"
	}

	if s.SkipToken != "" {
		params[SkipToken] = s.SkipToken
	}

	if s.StatusOnly {
		params[StatusOnly] = "true"
	}

	if s.TenantId != "" {
		params[TenantId] = s.TenantId
	}
	if s.Top > 0 {
		params[Top] = strconv.FormatInt(int64(s.Top), 10)
	}

	return params
}

type GraphParams struct {
	Count     bool
	Expand    string
	Format    string
	Filter    string
	OrderBy   string
	Search    string
	Select    []string
	Skip      int
	Top       int32
	SkipToken string
}

func (s GraphParams) NeedsEventualConsistencyHeaderFlag() bool {
	return s.Count || s.Search != "" || s.OrderBy != "" || (s.Filter != "" && s.OrderBy != "") || strings.Contains(s.Filter, "endsWith")
}

func (s GraphParams) AsMap() map[string]string {
	params := make(map[string]string)

	if s.Count {
		params[Count] = "true"
	}

	if s.Expand != "" {
		params[Expand] = s.Expand
	}

	if s.Format != "" {
		params[Format] = s.Format
	}

	if s.Filter != "" {
		params[Filter] = s.Filter
	}

	if s.OrderBy != "" {
		params[OrderBy] = s.OrderBy
	}

	if s.Search != "" {
		params[Search] = s.Search
	}

	if len(s.Select) > 0 {
		params[Select] = strings.Join(s.Select, ",")
	}

	if s.Skip > 0 {
		params[Skip] = strconv.Itoa(s.Skip)
	}

	if s.SkipToken != "" {
		params[SkipToken] = s.SkipToken
	}

	if s.Top > 0 {
		params[Top] = strconv.FormatInt(int64(s.Top), 10)
	}

	return params
}
