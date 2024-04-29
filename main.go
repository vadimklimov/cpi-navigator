package main

import (
	"cmp"

	"github.com/vadimklimov/cpi-navigator/cmd"
	"github.com/vadimklimov/cpi-navigator/internal"
)

// Set using ldflags during build.
var version string

func main() {
	internal.AppVersion = cmp.Or(version, "unknown")

	cmd.Execute()
}
