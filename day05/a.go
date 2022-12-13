package day05

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Stack []string

type PuzzleA struct{}

func (p PuzzleA) String() string {
	return "05a"
}

type Instruction struct {
	height int
	srcIdx int
	dstInx int
}

// IsEmpty - checks if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

// Pop removes and returns the top element of the stack. Return false if stack is empty.
func (s *Stack) Pop() string {
	//if s.IsEmpty() {
	//	return "", false
	//}
	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the stack by slicing it off.
	return element
}

func (p PuzzleA) Run() any {
	inputArr := strings.Split(input, "\n")
	maxStackHeight := p.GetMaxStackHeight(&inputArr)

	stacks := p.InitStacks(&inputArr, maxStackHeight)
	instructions := p.InitInstructions(&inputArr, maxStackHeight)
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

func (p PuzzleA) InitInstructions(inputArr *[]string, startIdx int) []Instruction {
	var instr []Instruction

	for i := startIdx + 2; i < len(*inputArr)-1; i++ {
		row := strings.Split((*inputArr)[i], " ")
		h, _ := strconv.Atoi(row[1])
		s, _ := strconv.Atoi(row[3])
		d, _ := strconv.Atoi(row[5])
		i := Instruction{
			height: h - 1,
			srcIdx: s - 1,
			dstInx: d - 1,
		}
		instr = append(instr, i)
	}

	return instr
}

// InitStacks reads values from input in a O(N^2) approach.
func (p PuzzleA) InitStacks(inputArr *[]string, height int) []Stack {
	var stackList []Stack

	width := len((*inputArr)[7]) // matrix width

	for w := 0; w <= width; w++ { // x axis (width)
		var stack Stack
		for h := height - 1; h >= 0; h-- { // y axis (height)
			maxRowLength := len((*inputArr)[h])
			if w >= maxRowLength { // escape for idx out of range
				break
			}
			val := int32((*inputArr)[h][w])
			if strings.ContainsRune(abc, val) { // rune is faster
				stack.Push(string(val))
			}
		}
		if stack != nil {
			stackList = append(stackList, stack)
		}
	}

	return stackList
}

func (p PuzzleA) GetMaxStackHeight(input *[]string) int {
	height := 0
	for (*input)[height] != "" {
		height++
	}
	return height - 1
}

// map input to a matrix
// map matrix to stacks
