package main

import (
	"fmt"
	"os"
	"scrub/pkg/commands"
)

var (
	// Populated by goreleaser during build
	version = "master"
	commit  = "?"
	date    = ""
)

func main() {
	e := commands.NewExecutor(version, commit, date)
	if err := e.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "failed with error %v\n", err)
		os.Exit(1)
	}
}
