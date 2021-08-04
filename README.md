Hyperscale logger [![Last release](https://img.shields.io/github/release/hyperscale-stack/logger.svg)](https://github.com/hyperscale-stack/logger/releases/latest) [![Documentation](https://godoc.org/github.com/hyperscale-stack/logger?status.svg)](https://godoc.org/github.com/hyperscale-stack/logger)
====================

[![Go Report Card](https://goreportcard.com/badge/github.com/hyperscale-stack/logger)](https://goreportcard.com/report/github.com/hyperscale-stack/logger)

| Branch  | Status | Coverage |
|---------|--------|----------|
| master  | [![Build Status](https://github.com/hyperscale-stack/logger/workflows/Go/badge.svg?branch=master)](https://github.com/hyperscale-stack/logger/actions?query=workflow%3AGo) | [![Coveralls](https://img.shields.io/coveralls/hyperscale-stack/logger/master.svg)](https://coveralls.io/github/hyperscale-stack/logger?branch=master) |

The Hyperscale logger is an standard interface for abstract logger provider

## Example

```go
package main

import (
    "fmt"
    "github.com/hyperscale-stack/logger"
)

func main() {
    l := NewStandardLogger(&outBuf, &errBuf)

	l.Panic("my panic message without option")
	l.Fatal("my fatal message without option")
	l.Error("my error message without option")
	l.Warning("my warning message without option")
	l.Info("my info message without option")
	l.Debug("my debug message without option")
	l.Trace("my trace message without option")

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


    l.Panic("my panic message with option format: %s => %d", WithFormat("foo", 1))
	l.Fatal("my fatal message with option format: %s => %d", WithFormat("foo", 1))
	l.Error("my error message with option format: %s => %d", WithFormat("foo", 1))
	l.Warning("my warning message with option format: %s => %d", WithFormat("foo", 1))
	l.Info("my info message with option format: %s => %d", WithFormat("foo", 1))
	l.Debug("my debug message with option format: %s => %d", WithFormat("foo", 1))
	l.Trace("my trace message with option format: %s => %d", WithFormat("foo", 1))
}

```

## License

Hyperscale logger is licensed under [the MIT license](LICENSE.md).
