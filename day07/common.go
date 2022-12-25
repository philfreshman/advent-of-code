package day07

import (
	_ "embed"
)

//go:embed input.txt
var input string

func isFile(t interface{}) bool {
	switch t.(type) {
	case File:
		return true
	default:
		return false
	}
}
