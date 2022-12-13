package day04

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type StringPair struct {
	elf1 string
	elf2 string
}

type SlicePair struct {
	elf1 []int
	elf2 []int
}

// SliceFromSection turns range (ex. "2-4") into []int{2,3,4}
func SliceFromSection(section string) []int {
	sectionRange := strings.Split(section, "-")
	start, _ := strconv.Atoi(sectionRange[0])
	end, _ := strconv.Atoi(sectionRange[1])
	sectionArray := []int{start}
	last := sectionArray[len(sectionArray)-1]

	for last < end {
		sectionArray = append(sectionArray, last+1)
		last = sectionArray[len(sectionArray)-1]
	}

	return sectionArray
}
