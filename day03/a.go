package day03

import (
	_ "embed"
	"strings"
	"unicode"
)

//go:embed input.txt
var	inputA string

type PuzzleA struct{}

func (p PuzzleA) String() string {
	return "03a"
}

func (p PuzzleA) Run() int{
	items := strings.Split(inputA, "\n")
	sum := 0

	for _, val := range items{
		l := len(val)
		firstHalf := val[:l/2]

		for i := l/2; i < l; i++ {
			char := string(val[i])
			if strings.Contains(firstHalf, char){
				if unicode.IsUpper([]rune(char)[0]) {
					upIdx := strings.Index(upper, char)
					sum += upIdx + 27
					break
				} else {
					lowIdx := strings.Index(lower, char)
					sum += lowIdx + 1
					break
				}
			}
		}
	}

	return sum

}

