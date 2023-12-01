package day02

import (
	"bufio"
	"os"
	"strings"
)

type PuzzleB struct{}

func (b PuzzleB) String() string {
	return "02b"
}

func (b PuzzleB) Run() any {
	myDir, _ := os.Getwd()
	fileData, err := os.Open(myDir + "/input2.txt.txt")
	if err != nil {
		return 0
	}
	result := b.playGame(fileData)
	return result
}

func (b PuzzleB) playGame(f *os.File) int {
	input := bufio.NewScanner(f)
	sum := 0

	for input.Scan() {
		x := input.Text()
		opponent := strings.Split(x, " ")[0]
		player := strings.Split(x, " ")[1]
		res := b.fight(shapes[opponent], shapes[player])
		sum += res
	}

	return sum
}

func (b PuzzleB) fight(playerA int, playerX int) int {
	switch playerX {
	case 1:
		return b.RockHandler(playerA)
	case 2:
		return b.PaperHandler(playerA)
	default:
		return b.ScissorHandler(playerA)
	}
}

func (b PuzzleB) ScissorHandler(playerA int) int {
	switch playerA {
	case 1:
		return calcWin(paper)
	case 2:
		return calcWin(scissors)
	default:
		return calcWin(rock)
	}
}

func (b PuzzleB) PaperHandler(playerA int) int {
	switch playerA {
	case 1:
		return calcDraw(rock)
	case 2:
		return calcDraw(paper)
	default:
		return calcDraw(scissors)
	}
}

func (b PuzzleB) RockHandler(playerA int) int {
	switch playerA {
	case 1:
		return calcLose(scissors)
	case 2:
		return calcLose(rock)
	default:
		return calcLose(paper)
	}
}
