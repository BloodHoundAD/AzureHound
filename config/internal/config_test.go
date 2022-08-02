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

package internal

import (
	"testing"

	"github.com/spf13/cobra"
)

var (
	fooConfig = Config{
		Name:      "foo",
		Shorthand: "f",
		Usage:     "configure foo",
		Default:   "foo",
	}

	barConfig = Config{
		Name:      "bar",
		Shorthand: "b",
		Usage:     "configure bar",
		Default:   1,
	}

	bazConfig = Config{
		Name:       "baz",
		Usage:      "configure baz",
		Persistent: true,
		Required:   true,
		Default:    false,
	}

	cmd = cobra.Command{
		Use: "test",
		Run: func(cmd *cobra.Command, args []string) {},
	}
)

func init() {
	Init(&cmd, []Config{fooConfig, barConfig, bazConfig})
}

func TestFooConfig(t *testing.T) {
	cmd.Execute()

	if actual := fooConfig.Value(); actual != "foo" {
		t.Errorf("got %s, want %s\n", actual, "foo")
	}

	fooConfig.Set("bar")

	if actual := fooConfig.Value(); actual != "bar" {
		t.Errorf("got %s, want %s\n", actual, "bar")
	}
}

func TestBarConfig(t *testing.T) {
	cmd.Execute()

	if actual := barConfig.Value(); actual != barConfig.Default {
		t.Errorf("got %v, want %v\n", actual, barConfig.Default)
	}

	barConfig.Set(2)

	if actual := barConfig.Value(); actual != 2 {
		t.Errorf("got %v, want %v\n", actual, 2)
	}
}

func TestBazConfig(t *testing.T) {
	cmd.Execute()

	if actual := bazConfig.Value(); actual != bazConfig.Default {
		t.Errorf("got %v, want %v\n", actual, bazConfig.Default)
	}

	bazConfig.Set(true)

	if actual := bazConfig.Value(); actual != true {
		t.Errorf("got %v, want %v\n", actual, true)
	}
}
