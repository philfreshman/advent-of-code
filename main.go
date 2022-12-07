package main

import (
	"os"
)

// go run . aoc-2022 --benchmark --puzzle 1

func main() {
	cmd, err := newRootCmd(os.Args[1:], os.Stdout)
	if err != nil {
		os.Exit(1)
	}

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
