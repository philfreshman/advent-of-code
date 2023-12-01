package day02

import (
	"bufio"
	"os"
	"strings"
)

type PuzzleA struct{}

func (p PuzzleA) String() string {
	return "02a"
}

func (p PuzzleA) Run() any {
	myDir, _ := os.Getwd()
	fileData, err := os.Open(myDir + "/input2.txt.txt")
	if err != nil {
		return 0
	}
	result := playGame(fileData)
	return result
}

func playGame(f *os.File) int {
	input := bufio.NewScanner(f)
	sum := 0

	for input.Scan() {
		x := input.Text()
		opponent := strings.Split(x, " ")[0]
		player := strings.Split(x, " ")[1]
		res := fight(shapes[opponent], shapes[player])
		sum += res
	}

	return sum
}

func fight(playerA int, playerX int) int {
	switch playerX {
	case 1:
		return RockHandler(playerA)
	case 2:
		return PaperHandler(playerA)
	default:
		return ScissorHandler(playerA)
	}
}

func ScissorHandler(playerA int) int {
	switch playerA {
	case 1:
		return calcLose(scissors)
	case 2:
		return calcWin(scissors)
	default:
		return calcDraw(scissors)
	}
}

func PaperHandler(playerA int) int {
	switch playerA {
	case 1:
		return calcWin(paper)
	case 2:
		return calcDraw(paper)
	default:
		return calcLose(paper)
	}
}

func RockHandler(playerA int) int {
	switch playerA {
	case 1:
		return calcDraw(rock)
	case 2:
		return calcLose(rock)
	default:
		return calcWin(rock)
	}
}
