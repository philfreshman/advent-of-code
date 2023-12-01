package day03

import (
	_ "embed"
	"strings"
	"unicode"
)

//go:embed input.txt
var inputB string

type PuzzleB struct{}

func (p PuzzleB) String() string {
	return "03b"
}

type Group struct {
	elf1 string
	elf2 string
	elf3 string
}

func (p PuzzleB) Run() any {
	items := strings.Split(inputB, "\n")

	var g Group
	var duplicate string
	var sum = 0

	for i := 0; i < len(items); i += 3 {
		g = Group{
			items[i], items[i+1], items[i+2],
		}

		for _, val := range g.elf1 {
			test := string(val)
			sec := strings.Contains(g.elf2, test)
			thi := strings.Contains(g.elf3, test)
			if sec && thi {
				duplicate = string(val)
				break
			}
		}

		if unicode.IsUpper([]rune(duplicate)[0]) {
			upIdx := strings.Index(upper, duplicate)
			sum += upIdx + 27
		} else {
			lowIdx := strings.Index(lower, duplicate)
			sum += lowIdx + 1
		}
	}

	return sum
}
