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
	"io"
	"os"
	"path"
	"time"

	"github.com/go-logr/logr"
	"github.com/rs/zerolog"
)

const (
	BaseCallDepth int = 2

	// ErrorLevel limits logs to ERROR messages only
	ErrorLevel int = -1
	// MaxInfoLevel allows ERROR, INFO, DEBUG and TRACE messages
	MaxInfoLevel int = 2
	// MedInfoLevel allows ERROR, INFO, and DEBUG messages
	MedInfoLevel int = 1
	// MinInfoLevel limits logs to ERROR and INFO messages
	MinInfoLevel int = 0
)

type logSink struct {
	logger    *zerolog.Logger
	name      string
	callDepth int
}

// Options allows the user to set various options for the logr.Logger implementation
type Options struct {
	// Structured enables structured logging
	Structured bool
	// Colors enables colors for unstructured logging only
	Colors bool
	// Writers defines which transports the logger should write to
	Writers []io.Writer
	// Level defines the logging verbosity; defaults to MinInfoLevel
	Level int
}

// NewLogger returns a new logr.Logger instance
func NewLogger(options Options) logr.Logger {
	if len(options.Writers) == 0 {
		options.Writers = append(options.Writers, os.Stderr)
	}

	if !options.Structured {
		for i, writer := range options.Writers {
			options.Writers[i] = zerolog.ConsoleWriter{Out: writer, NoColor: !options.Colors, TimeFormat: time.RFC3339}
		}
	}

	writer := zerolog.MultiLevelWriter(options.Writers...)
	logger := zerolog.New(writer).With().Timestamp().Logger()

	if options.Level < MinInfoLevel {
		logger = logger.Level(zerolog.ErrorLevel)
	} else {
		lvl := calcLevel(options.Level)
		logger = logger.Level(lvl)
	}

	return logr.New(&logSink{
		logger:    &logger,
		name:      "",
		callDepth: BaseCallDepth,
	})
}

// Enabled tests whether this logr.LogSink is enabled at the specified V-level.
// For example, commandline flags might be used to set the logging
// verbosity and disable some info logs.
func (s logSink) Enabled(level int) bool {
	lvl := calcLevel(level)
	if logEvent := s.logger.WithLevel(lvl); logEvent == nil {
		return false
	} else {
		return logEvent.Enabled()
	}
}

// Error logs an error, with the given message and key/value pairs as
// context. See logr.Logger.Error for more details.
func (s logSink) Error(err error, msg string, keysAndValues ...interface{}) {
	logEvent := s.logger.Error().Err(err)
	s.log(logEvent, msg, keysAndValues)
}

// Info logs a non-error message with the given key/value pairs as context.
// The level argument is provided for optional logging.  This method will
// only be called when Enabled(level) is true. See logr.Logger.Info for more
// details.
func (s logSink) Info(level int, msg string, keysAndValues ...interface{}) {
	lvl := calcLevel(level)
	logEvent := s.logger.WithLevel(lvl)
	s.log(logEvent, msg, keysAndValues)
}

// Init receives optional information about the logr library for logr.LogSink
// implementations that need it.
func (s *logSink) Init(info logr.RuntimeInfo) {
	s.callDepth = info.CallDepth + BaseCallDepth
}

// WithName returns a new logr.LogSink with the specified name appended.  See
// logr.Logger.WithName for more details.
func (s logSink) WithName(name string) logr.LogSink {
	s.name = path.Join(s.name, name)
	return &s
}

// WithValues returns a new logr.LogSink with additional key/value pairs. See
// logr.Logger.WithValues for more details.
func (s logSink) WithValues(keysAndValues ...interface{}) logr.LogSink {
	logger := s.logger.With().Fields(keysAndValues).Logger()
	s.logger = &logger
	return &s
}

// WithCallDepth returns a logr.LogSink that will offset the call
// stack by the specified number of frames when logging call
// site information.
//
// If depth is 0, the logr.LogSink should skip exactly the number
// of call frames defined in logr.RuntimeInfo.CallDepth when Info
// or Error are called, i.e. the attribution should be to the
// direct caller of logr.Logger.Info or logr.Logger.Error.
//
// If depth is 1 the attribution should skip 1 call frame, and so on.
// Successive calls to this are additive.
func (s logSink) WithCallDepth(depth int) logr.LogSink {
	s.callDepth += depth
	return &s
}

func (s logSink) log(e *zerolog.Event, msg string, keysAndValues []interface{}) {
	if e != nil {
		if s.name != "" {
			e.Str("name", s.name)
		}

		e.Fields(keysAndValues)
		e.CallerSkipFrame(s.callDepth)
		e.Msg(msg)
	}
}

func calcLevel(level int) zerolog.Level {
	lvl := level
	if level < MinInfoLevel {
		lvl = MinInfoLevel
	} else if level > MaxInfoLevel {
		lvl = MaxInfoLevel
	}
	return zerolog.InfoLevel - zerolog.Level(lvl)
}
