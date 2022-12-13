package day05

import (
	_ "embed"
)

//go:embed input.txt
var _ string

type PuzzleB struct{}

func (p PuzzleB) String() string {
	return "05b"
}

func (p PuzzleB) Run() any {

	return 0
}
