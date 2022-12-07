package day03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPuzzleB(t *testing.T){
	p := PuzzleB{}
	res := p.Run()
	assert.Equal(t, 2703, res)
}