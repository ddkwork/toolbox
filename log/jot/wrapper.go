// Copyright ©2016-2023 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package jot

import (
	"io"

	"github.com/ddkwork/toolbox/log/logadapter"
)

// Logger wraps the various jot function calls into a struct that can be passed around, typically for the sake of
// satisfying one or more logging interfaces.
//
// Deprecated: Use slog instead. August 28, 2023
type Logger struct{}

// SetWriter sets the io.Writer to use when writing log messages. Default is os.Stderr.
//
// Deprecated: Use slog instead. August 28, 2023
func (lgr *Logger) SetWriter(w io.Writer) {
	SetWriter(w)
}

// SetMinimumLevel sets the minimum log level that will be output. Default is DEBUG.
//
// Deprecated: Use slog instead. August 28, 2023
func (lgr *Logger) SetMinimumLevel(level Level) {
	SetMinimumLevel(level)
}

// Debug logs a debug message. Arguments are handled in the manner of fmt.Print.
//
// Deprecated: Use slog instead. August 28, 2023
func (lgr *Logger) Debug(v ...any) {
	Debug(v...)
}

// Debugf logs a debug message. Arguments are handled in the manner of fmt.Printf.
//
// Deprecated: Use slog instead. August 28, 2023
func (lgr *Logger) Debugf(format string, v ...any) {
	Debugf(format, v...)
}

// Info logs an informational message. Arguments are handled in the manner of fmt.Print.
//
// Deprecated: Use slog instead. August 28, 2023
func (lgr *Logger) Info(v ...any) {
	Info(v...)
}

// Infof logs an informational message. Arguments are handled in the manner of fmt.Printf.
//
// Deprecated: Use slog instead. August 28, 2023
func (lgr *Logger) Infof(format string, v ...any) {
	Infof(format, v...)
}

// Warn logs a warning message. Arguments are handled in the manner of fmt.Print.
//
// Deprecated: Use slog instead. August 28, 2023
func (lgr *Logger) Warn(v ...any) {
	Warn(v...)
}

// Warnf logs a warning message. Arguments are handled in the manner of fmt.Printf.
//
// Deprecated: Use slog instead. August 28, 2023
func (lgr *Logger) Warnf(format string, v ...any) {
	Warnf(format, v...)
}

// Error logs an error message. Arguments are handled in the manner of fmt.Print.
//
// Deprecated: Use slog instead. August 28, 2023
func (lgr *Logger) Error(v ...any) {
	Error(v...)
}

// Errorf logs an error message. Arguments are handled in the manner of fmt.Printf.
//
// Deprecated: Use slog instead. August 28, 2023
func (lgr *Logger) Errorf(format string, v ...any) {
	Errorf(format, v...)
}

// Fatal logs a fatal error message. Arguments other than the status are handled in the manner of fmt.Print.
//
// Deprecated: Use slog instead. August 28, 2023
func (lgr *Logger) Fatal(status int, v ...any) {
	Fatal(status, v...)
}

// Fatalf logs a fatal error message. Arguments other than the status are handled in the manner of fmt.Printf.
//
// Deprecated: Use slog instead. August 28, 2023
func (lgr *Logger) Fatalf(status int, format string, v ...any) {
	Fatalf(status, format, v...)
}

// Time starts timing an event and logs an informational message. Arguments are handled in the manner of fmt.Print.
//
// Deprecated: Use slog instead. August 28, 2023
func (lgr *Logger) Time(v ...any) logadapter.Timing {
	return Time(v...)
}

// Timef starts timing an event and logs an informational message. Arguments are handled in the manner of fmt.Printf.
//
// Deprecated: Use slog instead. August 28, 2023
func (lgr *Logger) Timef(format string, v ...any) logadapter.Timing {
	return Timef(format, v...)
}

// Flush waits for all current log entries to be written before returning.
//
// Deprecated: Use slog instead. August 28, 2023
func (lgr *Logger) Flush() {
	Flush()
}

// Writer logs the data as an error after casting it to a string.
//
// Deprecated: Use slog instead. August 28, 2023
func (lgr *Logger) Write(data []byte) (int, error) {
	Error(string(data))
	return len(data), nil
}
