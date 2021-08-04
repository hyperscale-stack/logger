// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package logger

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStandardWithoutOption(t *testing.T) {
	var outBuf bytes.Buffer
	var errBuf bytes.Buffer
	l := NewStandardLogger(&outBuf, &errBuf)

	l.Panic("my panic message without option")
	l.Fatal("my fatal message without option")
	l.Error("my error message without option")
	l.Warning("my warning message without option")
	l.Info("my info message without option")
	l.Debug("my debug message without option")
	l.Trace("my trace message without option")

	assert.Equal(t, "[INFO]: my info message without option\n[DEBUG]: my debug message without option\n[TRACE]: my trace message without option\n", outBuf.String())
	assert.Equal(t, "[PANIC]: my panic message without option\n[FATAL]: my fatal message without option\n[ERROR]: my error message without option\n[WARNING]: my warning message without option\n", errBuf.String())
}

func TestStandardWithLabels(t *testing.T) {
	var outBuf bytes.Buffer
	var errBuf bytes.Buffer
	l := NewStandardLogger(&outBuf, &errBuf)

	labels := map[string]interface{}{
		"foo": "bar",
	}

	l.Panic("my panic message with option labels", WithLabels(labels))
	l.Fatal("my fatal message with option labels", WithLabels(labels))
	l.Error("my error message with option labels", WithLabels(labels))
	l.Warning("my warning message with option labels", WithLabels(labels))
	l.Info("my info message with option labels", WithLabels(labels))
	l.Debug("my debug message with option labels", WithLabels(labels))
	l.Trace("my trace message with option labels", WithLabels(labels))

	assert.Equal(t, "[INFO]: my info message with option labels foo=bar\n[DEBUG]: my debug message with option labels foo=bar\n[TRACE]: my trace message with option labels foo=bar\n", outBuf.String())
	assert.Equal(t, "[PANIC]: my panic message with option labels foo=bar\n[FATAL]: my fatal message with option labels foo=bar\n[ERROR]: my error message with option labels foo=bar\n[WARNING]: my warning message with option labels foo=bar\n", errBuf.String())
}

func TestStandardWithFormat(t *testing.T) {
	var outBuf bytes.Buffer
	var errBuf bytes.Buffer
	l := NewStandardLogger(&outBuf, &errBuf)

	l.Panic("my panic message with option format: %s => %d", WithFormat("foo", 1))
	l.Fatal("my fatal message with option format: %s => %d", WithFormat("foo", 1))
	l.Error("my error message with option format: %s => %d", WithFormat("foo", 1))
	l.Warning("my warning message with option format: %s => %d", WithFormat("foo", 1))
	l.Info("my info message with option format: %s => %d", WithFormat("foo", 1))
	l.Debug("my debug message with option format: %s => %d", WithFormat("foo", 1))
	l.Trace("my trace message with option format: %s => %d", WithFormat("foo", 1))

	assert.Equal(t, "[INFO]: my info message with option format: foo => 1\n[DEBUG]: my debug message with option format: foo => 1\n[TRACE]: my trace message with option format: foo => 1\n", outBuf.String())
	assert.Equal(t, "[PANIC]: my panic message with option format: foo => 1\n[FATAL]: my fatal message with option format: foo => 1\n[ERROR]: my error message with option format: foo => 1\n[WARNING]: my warning message with option format: foo => 1\n", errBuf.String())
}
