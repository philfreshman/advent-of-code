package day05

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPuzzleB(t *testing.T) {
	p := PuzzleB{}
	res := p.Run()
	assert.Equal(t, "TZLTLWRNF", res)
}
