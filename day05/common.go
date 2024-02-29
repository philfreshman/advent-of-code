package day05

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var abc = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Instruction struct {
	height int
	srcIdx int
	dstInx int
}

type Stack []string

// IsEmpty - checks if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push adds one item at the top of the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

// Append add multiple items at the top of the stack
func (s *Stack) Append(stack []string) {
	for i := len(stack) - 1; i >= 0; i-- {
		(*s).Push(stack[i])
	}
}

// Pop removes and returns the top element of the stack. Return false if stack is empty.
func (s *Stack) Pop() string {
	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the stack by slicing it off.
	return element
}

// PopLast removes and returns the top n elements of the stack.
func (s *Stack) PopLast(n int) []string {
	var result []string
	lastIdx := len(*s) - 1
	for i := lastIdx; i > lastIdx-n; i-- {
		x := (*s)[i]
		result = append(result, x)
		*s = (*s)[:i]
	}
	return result
}

func GetMaxStackHeight(input *[]string) int {
	height := 0
	for (*input)[height] != "" {
		height++
	}
	return height - 1
}

// InitStacks reads values from input in a O(N^2) approach.
func InitStacks(inputArr *[]string, height int) []Stack {
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

func InitInstructions(inputArr *[]string, startIdx int) []Instruction {
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
