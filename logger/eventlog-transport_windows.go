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
	"syscall"

	"github.com/rs/zerolog"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/svc/eventlog"
)

func NewEventLogWriter(name string) (zerolog.LevelWriter, error) {
	if elog, err := eventlog.Open(name); err != nil {
		return nil, err
	} else {
		return EventLogLevelWriter{elog}, nil
	}
}

type EventLogLevelWriter struct {
	eventLog *eventlog.Log
}

func (s EventLogLevelWriter) Write(msg []byte) (int, error) {
	return s.WriteLevel(zerolog.InfoLevel, msg)
}

func (s EventLogLevelWriter) WriteLevel(level zerolog.Level, msg []byte) (n int, err error) {
	var eventType uint16
	switch level {
	case zerolog.Disabled:
		return 0, nil
	case zerolog.ErrorLevel:
		eventType = windows.EVENTLOG_ERROR_TYPE
	default:
		eventType = windows.EVENTLOG_INFORMATION_TYPE
	}

	sysString := []*uint16{syscall.StringToUTF16Ptr(string(msg))}
	if err := windows.ReportEvent(s.eventLog.Handle, eventType, 0, 1, 0, 1, 0, &sysString[0], nil); err != nil {
		return 0, err
	} else {
		return len(msg), nil
	}
}

func (s EventLogLevelWriter) Close() error {
	return s.eventLog.Close()
}
