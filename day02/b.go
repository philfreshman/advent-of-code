package day01

import (
	"bufio"
	"os"
	"strconv"
)

var leaders = [3]int{}

func (p PuzzleB) String() string {
	return "01b"
}

type PuzzleB struct{}

func (p PuzzleB) Run() int {
	myDir, _ := os.Getwd()
	fileData, _ := os.Open(myDir + "/input.txt")
	getHighestCal(fileData)
	_ = fileData.Close()

	sum := 0
	for _, val := range leaders {
		sum += val
	}
	return sum
}

func getHighestCal(f *os.File) {
	input := bufio.NewScanner(f)
	//sum := 0
	tempSum := 0

	for input.Scan() {
		value := input.Text()
		number, _ := strconv.Atoi(value)

		if value == "" {
			handleLeaders(tempSum)
			tempSum = 0
		} else {
			tempSum += number
		}
	}

}

func handleLeaders(n int) {
	if n > leaders[0] {
		leaders[2] = leaders[1]
		leaders[1] = leaders[0]
		leaders[0] = n
	} else if n > leaders[1] && n < leaders[0] {
		leaders[2] = leaders[1]
		leaders[1] = n
	} else if n > leaders[2] && n < leaders[1] {
		leaders[2] = n
	}

}
