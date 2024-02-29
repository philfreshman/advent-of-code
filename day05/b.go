package day05

import (
	_ "embed"
	"strings"
)

type PuzzleB struct{}

func (p PuzzleB) String() string {
	return "05b"
}

func (p PuzzleB) Run() any {
	inputArr := strings.Split(input, "\n")
	maxStackHeight := GetMaxStackHeight(&inputArr)

	stacks := InitStacks(&inputArr, maxStackHeight)
	instructions := InitInstructions(&inputArr, maxStackHeight)
	p.ReadInstructions(&stacks, &instructions)

	result := ""
	for _, stack := range stacks {
		result = result + stack[len(stack)-1]
	}

	return result
}

func (p PuzzleB) ReadInstructions(stacks *[]Stack, instr *[]Instruction) {
	for _, val := range *instr {
		times := val.height + 1
		srcIdx := val.srcIdx
		dstIdx := val.dstInx

		rest := (*stacks)[srcIdx].PopLast(times)
		(*stacks)[dstIdx].Append(rest)
		times--
	}
}
