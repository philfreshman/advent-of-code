package day06

import (
	_ "embed"
)

type PuzzleA struct{}

func (p PuzzleA) String() string {
	return "06a"
}

func (p PuzzleA) Run() any {
	var que Que
	var allDiff bool
	var resultIdx int

	for i := 0; i < len(input); i++ {
		if i <= 3 {
			que.Append(string(input[i]))
		} else {
			allDiff = que.CheckIfAllDifferent()
			if allDiff {
				resultIdx = i
				break
			}
			que.Append(string(input[i]))
			que.PopFirst()
		}
	}

	return resultIdx
}
