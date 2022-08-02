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

package logger

import (
	"io"
	"os"

	"github.com/bloodhoundad/azurehound/config"
	"github.com/bloodhoundad/azurehound/constants"
	logger "github.com/bloodhoundad/azurehound/logger/internal"
	"github.com/go-logr/logr"
	"github.com/rs/zerolog"
	"golang.org/x/sys/windows/svc"
)

var eventLogWriter zerolog.LevelWriter

func setupLogger() (*logr.Logger, error) {
	var (
		logr    logr.Logger
		options = logger.Options{
			Level:      config.VerbosityLevel.Value().(int),
			Structured: config.JsonLogs.Value().(bool),
			Colors:     false,
			Writers:    []io.Writer{os.Stderr},
		}
	)

	// services should emit messages to the windows event log
	if eventLogWriter := getEventLogLevelWriter(); eventLogWriter != nil {
		options.Writers = append(options.Writers, eventLogWriter)
	}

	// XXX: This is gross, however, reading in the config file when starting the process as a windows service before
	// initializing the eventLogWriter causes the program to panic. It doesn't make sense as to why it does that but
	// this call will have to remain here until we can figure out what's going on.
	config.LoadValues(nil, config.Options())

	// emit logs to file if configured
	if fileLogWriter := getFileLogLevelWriter(); fileLogWriter != nil {
		options.Writers = append(options.Writers, fileLogWriter)
	}

	logr = logger.NewLogger(options)
	return &logr, nil
}

func getEventLogLevelWriter() zerolog.LevelWriter {
	if eventLogWriter != nil {
		return eventLogWriter
	} else if isWindowsService, err := svc.IsWindowsService(); !isWindowsService || err != nil {
		return nil
	} else if eventLogWriter, err := NewEventLogWriter(constants.Name); err != nil {
		return nil
	} else {
		return eventLogWriter
	}
}
