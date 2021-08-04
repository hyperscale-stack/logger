// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package logger

var _ Logger = (*Nop)(nil)

type Nop struct{}

func (Nop) Log(level LevelType, message string, options ...Option) {}
func (Nop) Panic(message string, options ...Option)                {}
func (Nop) Fatal(message string, options ...Option)                {}
func (Nop) Error(message string, options ...Option)                {}
func (Nop) Warning(message string, options ...Option)              {}
func (Nop) Info(message string, options ...Option)                 {}
func (Nop) Debug(message string, options ...Option)                {}
func (Nop) Trace(message string, options ...Option)                {}
