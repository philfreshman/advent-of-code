package day03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPuzzleA(t *testing.T) {
	p := PuzzleA{}
	result := p.Run()
	assert.Equal(t, 7795, result)
}
