// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithLabels(t *testing.T) {
	labels := map[string]interface{}{
		"foo": "bar",
	}

	ctx := &Context{}
	o := WithLabels(labels)

	o(ctx)

	assert.Equal(t, labels, ctx.Labels)
}

func TestWithFormat(t *testing.T) {
	formats := []interface{}{"foo", "bar"}

	ctx := &Context{}
	o := WithFormat("foo", "bar")

	o(ctx)

	assert.Equal(t, formats, ctx.Formats)
}
