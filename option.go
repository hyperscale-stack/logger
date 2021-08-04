// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package logger

type Option func(*Context)

// WithLabels add labels to logger.Context.
func WithLabels(labels map[string]interface{}) Option {
	return func(c *Context) {
		c.Labels = labels
	}
}

// WithFormat add formats args to logger.Context.
func WithFormat(formats ...interface{}) Option {
	return func(c *Context) {
		c.Formats = formats
	}
}
