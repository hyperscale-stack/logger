// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelString(t *testing.T) {
	for _, item := range []struct {
		level    LevelType
		expected string
	}{
		{
			level:    LevelPanic,
			expected: "panic",
		},
		{
			level:    LevelFatal,
			expected: "fatal",
		},
		{
			level:    LevelError,
			expected: "error",
		},
		{
			level:    LevelWarning,
			expected: "warning",
		},
		{
			level:    LevelInfo,
			expected: "info",
		},
		{
			level:    LevelDebug,
			expected: "debug",
		},
		{
			level:    LevelTrace,
			expected: "trace",
		},
	} {
		assert.Equal(t, item.expected, item.level.String())
	}
}
