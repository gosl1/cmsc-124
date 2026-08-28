package main

import (
	"fmt"
	"os"
)

func fail(exitCode int, format string, arguments ...any) {
	fmt.Fprintf(os.Stderr, "lab0: "+format+"\n", arguments...)
	os.Exit(exitCode)
}

func main() {
	if len(os.Args) < 2 {
		fail(65, "expected one source-file path")
	}

	path := os.Args[1]
	source, err := os.ReadFile(path)
	if err != nil {
		fail(65, "cannot read '%s': %v", path, err)
	}

	if _, err := os.Stdout.Write(source); err != nil {
		fail(70, "cannot write output: %v", err)
	}
}