package day05

import (
	_ "embed"
	"strings"
)

type PuzzleA struct{}

func (p PuzzleA) String() string {
	return "05a"
}

func (p PuzzleA) Run() any {
	inputArr := strings.Split(input, "\n")
	maxStackHeight := GetMaxStackHeight(&inputArr)

	stacks := InitStacks(&inputArr, maxStackHeight)
	instructions := InitInstructions(&inputArr, maxStackHeight)
	p.ReadInstructions(&stacks, &instructions)

	result := ""
	for _, stack := range stacks {
		result = result + stack[len(stack)-1]
	}

	return 2
}

func (p PuzzleA) ReadInstructions(stacks *[]Stack, instr *[]Instruction) {
	for _, val := range *instr {
		times := val.height
		srcIdx := val.srcIdx
		dstIdx := val.dstInx

		for times >= 0 {
			(*stacks)[dstIdx].Push((*stacks)[srcIdx].Pop())
			times--
		}
	}
}
