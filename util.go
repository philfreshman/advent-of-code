package main

import "io/ioutil"

func ReadInputFile() string {
	data, _ := ioutil.ReadFile("input2.txt.txt")
	return string(data)
}
