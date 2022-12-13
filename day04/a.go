package day04

import (
	_ "embed"
	"github.com/philfreshman/advent-of-code-2022/helpers"
	"strings"
)

//go:embed input.txt
var input string

type PuzzleA struct{}

func (p PuzzleA) String() string {
	return "04a"
}

func (p PuzzleA) Run() any {
	result := 0
	arrOfSections := strings.Split(input, "\n")

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

		if p.DoesContainAll(slicePair) {
			result++
		}
	}

	return result
}

// DoesContainAll checks if for ex. slice {2,3,4} contains all numbers from {3,4}
func (p PuzzleA) DoesContainAll(pair SlicePair) bool {
	for _, val := range pair.elf2 {
		if !helpers.Contains(pair.elf1, val) {
			return false
		}
	}
	return true
}
