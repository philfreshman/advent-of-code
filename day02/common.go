package day02

const (
	lose = 0
	draw = 3
	win  = 6

	rock     = 1
	paper    = 2
	scissors = 3
)

var shapes = map[string]int{
	"A": rock,
	"X": rock,
	"B": paper,
	"Y": paper,
	"C": scissors,
	"Z": scissors,
}

func calcWin(shape int) int {
	return shape + win
}

func calcLose(shape int) int {
	return shape + lose
}

func calcDraw(shape int) int {
	return shape + draw
}
