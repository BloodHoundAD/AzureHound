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

type Definition struct {
	Schema string `json:"$schema,omitempty"`
	// Certain actions can be nested, have different elements based on the name(key) of given action - Condition is an example
	// Actions        map[string]Action       `json:"actions,omitempty"`
	Actions        map[string]interface{}  `json:"actions,omitempty"`
	ContentVersion string                  `json:"contentVersion,omitempty"`
	Outputs        map[string]Output       `json:"outputs,omitempty"`
	Parameters     map[string]Parameter    `json:"parameters,omitempty"`
	StaticResults  map[string]StaticResult `json:"staticResults,omitempty"`
	Triggers       map[string]Trigger      `json:"triggers,omitempty"`
}

type Action struct {
	Type string `json:"type"`
	// Kind is missing in the MSDN, but returned and present in examples and during testing
	Kind                 string                 `json:"kind,omitempty"`
	Inputs               map[string]interface{} `json:"inputs,omitempty"`
	RunAfter             interface{}            `json:"runAfter,omitempty"`
	RuntimeConfiguration interface{}            `json:"runtimeConfiguration,omitempty"`
	OperationOptions     string                 `json:"operationOptions,omitempty"`
}

type Output struct {
	Type string `json:"type,omitempty"`
	// Type of this is based on above Type
	Value interface{} `json:"value,omitempty"`
}

type Parameter struct {
	Type          string        `json:"type,omitempty"`
	DefaultValue  interface{}   `json:"defaultValue,omitempty"`
	AllowedValues []interface{} `json:"allowedValues,omitempty"`
	Metadata      Metadata      `json:"metadata,omitempty"`
}

type Metadata struct {
	Description interface{} `json:"description,omitempty"`
}

type StaticResult struct {
	Outputs ResultOutput `json:"outputs,omitempty"`
	Status  string       `json:"status,omitempty"`
}

type ResultOutput struct {
	Headers    map[string]string `json:"headers,omitempty"`
	StatusCode string            `json:"statusCode,omitempty"`
}

type Trigger struct {
	Type string `json:"type,omitempty"`
	// Kind is missing in the MSDN, but returned and present in examples and during testing
	Kind string `json:"kind,omitempty"`
	// Inputs is a custom element based on the type of trigger
	Inputs     interface{} `json:"inputs,omitempty"`
	Recurrence Recurrence  `json:"recurrence,omitempty"`
	Conditions []Condition `json:"conditions,omitempty"`
	// Runtime configuration is a custom element based on the type of trigger
	RuntimeConfiguration interface{} `json:"runtimeConfiguration,omitempty"`
	SplitOn              string      `json:"splitOn,omitempty"`
	OperationOptions     string      `json:"operationOptions,omitempty"`
}

type Recurrence struct {
	Frequency string `json:"frequency,omitempty"`
	Interval  int    `json:"interval,omitempty"`
}

type Condition struct {
	Expression string `json:"expression,omitempty"`
}
