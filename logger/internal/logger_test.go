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
	"bytes"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/go-logr/logr"
)

const (
	LogErrorTemplate            string = `{"level":"error","error":"I'm an error","time":"%s","message":"Something happened"}%s`
	LogInfoTemplate             string = `{"level":"info","foo":"bar","name":"fakeName","baz":42,"buzz":true,"time":"%s","message":"teapot"}%s`
	LogInfoUnstructuredTemplate string = "%s INF teapot baz=42 buzz=true foo=bar name=fakeName%s"
)

func TestError(t *testing.T) {
	writer := &bytes.Buffer{}
	options := Options{
		Structured: true,
		Writers:    []io.Writer{writer},
	}
	logger := NewLogger(options)
	logError(logger)()

	got := writer.String()
	want := fmt.Sprintf(LogErrorTemplate, now(), "\n")

	if got != want {
		t.Errorf("got: %v\nwant: %v", got, want)
	}
}

func TestInfo(t *testing.T) {
	writer := &bytes.Buffer{}
	options := Options{
		Structured: true,
		Writers:    []io.Writer{writer},
	}
	logger := NewLogger(options)
	logInfo(logger)()

	got := writer.String()
	want := fmt.Sprintf(LogInfoTemplate, now(), "\n")

	if got != want {
		t.Errorf("got: %v\nwant: %v", got, want)
	}
}

func TestInfoErrorLevel(t *testing.T) {
	writer := &bytes.Buffer{}
	options := Options{
		Structured: true,
		Writers:    []io.Writer{writer},
		Level:      ErrorLevel,
	}
	logger := NewLogger(options)
	logInfo(logger)()

	got := writer.String()
	want := ""

	if got != want {
		t.Errorf("got: %v\nwant: %v", got, want)
	}
}

func TestInfoUnstructured(t *testing.T) {
	writer := &bytes.Buffer{}
	options := Options{
		Writers: []io.Writer{writer},
	}
	logger := NewLogger(options)
	logInfo(logger)()

	got := writer.String()
	want := fmt.Sprintf(LogInfoUnstructuredTemplate, now(), "\n")

	if got != want {
		t.Errorf("got: %v\nwant: %v", got, want)
	}
}

func TestEnabled(t *testing.T) {
	errorsOnlyLogger := NewLogger(Options{Level: ErrorLevel})
	infoLogger := NewLogger(Options{Level: MinInfoLevel})

	if errorsOnlyLogger.GetSink().Enabled(MinInfoLevel) != false {
		t.Errorf("got: %v\nwant: %v", infoLogger.Enabled(), true)
	}

	if infoLogger.GetSink().Enabled(MinInfoLevel-1) != true {
		t.Errorf("got: %v\nwant: %v", infoLogger.Enabled(), true)
	}

	if infoLogger.GetSink().Enabled(MaxInfoLevel+1) != false {
		t.Errorf("got: %v\nwant: %v", infoLogger.Enabled(), false)
	}
}

func logInfo(logger logr.Logger) func() {
	return func() {
		logger.WithName("fakeName").WithValues("foo", "bar").Info("teapot", "baz", 42, "buzz", true)
	}
}

func logError(logger logr.Logger) func() {
	return func() {
		logger.Error(fmt.Errorf("I'm an error"), "Something happened")
	}
}

func now() string {
	return time.Now().Format(time.RFC3339)
}
