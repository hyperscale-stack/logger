// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package logger

// LevelType type.
type LevelType int

//nolint:exhaustive
func (l LevelType) String() string {
	switch l {
	case LevelPanic:
		return "panic"
	case LevelFatal:
		return "fatal"
	case LevelError:
		return "error"
	case LevelWarning:
		return "warning"
	case LevelInfo:
		return "info"
	case LevelDebug:
		return "debug"
	default:
		return "trace"
	}
}

// LevelType consts.
const (
	LevelPanic LevelType = iota
	LevelFatal
	LevelError
	LevelWarning
	LevelInfo
	LevelDebug
	LevelTrace
)
