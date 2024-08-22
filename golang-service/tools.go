//go:build tools

// These are packages and tools needed for building the project.

// This file will never be built (excluded due to its build tags),
// but `go mod tidy` will see the packages imported here as dependencies
// and not remove them from `go.mod`.

// See https://github.com/golang/go/issues/25922#issuecomment-413898264

package main

import (
	_ "github.com/go-delve/delve/cmd/dlv"
	_ "github.com/air-verse/air"
)
