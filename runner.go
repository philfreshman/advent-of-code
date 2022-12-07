package main

import "fmt"

type Runner interface {
	fmt.Stringer
	Run() int
}
