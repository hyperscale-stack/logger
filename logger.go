// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package logger

// Logger describes a logger instance.
type Logger interface {
	// Log with an arbitrary level.
	Log(level LevelType, message string, options ...Option)
	Panic(message string, options ...Option)
	Fatal(message string, options ...Option)
	Error(message string, options ...Option)
	Warning(message string, options ...Option)
	Info(message string, options ...Option)
	Debug(message string, options ...Option)
	Trace(message string, options ...Option)
}

// LoggerAware describes a logger-aware instance.
type LoggerAware interface {
	SetLogger(logger Logger)
}
