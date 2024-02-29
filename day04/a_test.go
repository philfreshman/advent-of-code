package day04

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPuzzleA(t *testing.T) {
	p := PuzzleA{}
	res := p.Run()
	assert.Equal(t, 494, res)
}
