package day04

import (
	_ "embed"
	"github.com/philfreshman/advent-of-code-2022/helpers"
	"strings"
)

//go:embed input.txt
var inputB string

type PuzzleB struct{}

func (p PuzzleB) String() string {
	return "04b"
}

func (p PuzzleB) Run() int {
	result := 0
	arrOfSections := strings.Split(inputB, "\n")

	for _, val := range arrOfSections {
		stringPair := StringPair{
			strings.Split(val, ",")[0],
			strings.Split(val, ",")[1],
		}

		slicePair := SlicePair{
			SliceFromSection(stringPair.elf1),
			SliceFromSection(stringPair.elf2),
		}

		// rearrange positions (bigger range first)
		if len(slicePair.elf1) < len(slicePair.elf2) {
			slicePair.elf1, slicePair.elf2 = slicePair.elf2, slicePair.elf1
		}

		if p.DoesContainAny(slicePair) {
			result++
		}
	}

	return result
}

// DoesContainAny checks if any of number of one slice are contained in the other
func (p PuzzleB) DoesContainAny(pair SlicePair) bool {
	for _, val := range pair.elf2 {
		if helpers.Contains(pair.elf1, val) {
			return true
		}
	}
	return false
}
