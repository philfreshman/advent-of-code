package day06

import (
	_ "embed"
)

//go:embed input.txt
var input string

type Que []string

// Append adds one element at the end of the que
func (q *Que) Append(s string) {
	*q = append(*q, s)
}

// PopFirst returns the first element in the que
// and shifts all elements by one index to the left
func (q *Que) PopFirst() string {
	temp := (*q)[0]

	q.shiftLeft()

	return temp
}

// shiftLeft shifts all elements in the que by 1
// index to the left and removes the last element
func (q *Que) shiftLeft() {
	for i := 0; i < len(*q)-1; i++ {
		(*q)[i] = (*q)[i+1]
	}
	(*q).trim()
}

// trim removes the last element of the Que
func (q *Que) trim() {
	*q = (*q)[:len(*q)-1]
}

func (q *Que) CheckIfAllDifferent() bool {
	return CheckRecursive(*q, 0)
}

func CheckRecursive(que Que, idx int) bool {
	if len(que) == idx+1 {
		return true
	}

	x := string(que[idx])
	for i := idx + 1; i < len(que); i++ {
		y := string(que[i])

		if x == y {
			return false
		}
	}
	return CheckRecursive(que, idx+1)
}
