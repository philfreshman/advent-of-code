package day01

import (
	"bufio"
	_ "embed"
	"os"
	"strconv"
)

type PuzzleA struct{}

func (p PuzzleA) String() string {
	return "01a"
}

func (p PuzzleA) Run() int {
	myDir, _ := os.Getwd()
	fileData, _ := os.Open(myDir + "/input.txt")
	result := getCalories(fileData)
	_ = fileData.Close()
	return result
}

// returns the highest sum of calories
func getCalories(f *os.File) int {
	input := bufio.NewScanner(f)
	sum := 0
	tempSum := 0

	for input.Scan() {
		value := input.Text()
		number, _ := strconv.Atoi(value)
		if value == "" && tempSum > sum {
			sum = tempSum
			tempSum = 0
		} else if value == "" {
			tempSum = 0
		} else {
			tempSum += number
		}
	}

	return sum
}
