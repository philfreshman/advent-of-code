package day06

import (
	_ "embed"
)

type PuzzleB struct{}

func (p PuzzleB) String() string {
	return "06b"
}

func (p PuzzleB) Run() any {
	var que Que
	var allDiff bool
	var resultIdx int
	for i := 0; i < len(input); i++ {
		if i <= 13 {
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
