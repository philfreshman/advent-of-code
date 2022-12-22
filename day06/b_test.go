package day06

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPuzzleB(t *testing.T) {
	p := PuzzleB{}
	res := p.Run()
	assert.Equal(t, 3298, res)
}
