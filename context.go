// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package logger

type Context struct {
	Labels  map[string]interface{}
	Formats []interface{}
}
