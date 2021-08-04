// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package logger

import (
	"fmt"
	"io"
	"strings"
)

var _ Logger = (*Nop)(nil)

type Standard struct {
	out io.Writer
	err io.Writer
}

func NewStandardLogger(out io.Writer, err io.Writer) Logger {
	return &Standard{
		out: out,
		err: err,
	}
}

func (l *Standard) Log(level LevelType, message string, options ...Option) {
	out := l.err

	if level > LevelWarning {
		out = l.out
	}

	ctx := &Context{}

	for _, opt := range options {
		opt(ctx)
	}

	fmt.Fprint(out, "["+strings.ToUpper(level.String())+"]: ")

	if len(ctx.Formats) > 0 {
		fmt.Fprintf(out, message, ctx.Formats...)
	} else {
		fmt.Fprint(out, message)
	}

	if len(ctx.Labels) > 0 {
		labels := []string{}

		fmt.Fprint(out, " ")

		for key, val := range ctx.Labels {
			labels = append(labels, fmt.Sprintf("%s=%s", key, val))
		}

		fmt.Fprint(out, strings.Join(labels, ", "))
	}

	fmt.Fprint(out, "\n")
}

func (l *Standard) Panic(message string, options ...Option) {
	l.Log(LevelPanic, message, options...)
}

func (l *Standard) Fatal(message string, options ...Option) {
	l.Log(LevelFatal, message, options...)
}

func (l *Standard) Error(message string, options ...Option) {
	l.Log(LevelError, message, options...)
}

func (l *Standard) Warning(message string, options ...Option) {
	l.Log(LevelWarning, message, options...)
}

func (l *Standard) Info(message string, options ...Option) {
	l.Log(LevelInfo, message, options...)
}

func (l *Standard) Debug(message string, options ...Option) {
	l.Log(LevelDebug, message, options...)
}

func (l *Standard) Trace(message string, options ...Option) {
	l.Log(LevelTrace, message, options...)
}
