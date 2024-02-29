package day02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPuzzleB(t *testing.T) {
	b := PuzzleB{}
	res := b.Run()
	assert.Equal(t, 12382, res)
}
