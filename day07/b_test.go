package day07

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPuzzleB(t *testing.T) {
	p := PuzzleB{}
	res := p.Run()
	assert.Equal(t, 3866390, res)
}
