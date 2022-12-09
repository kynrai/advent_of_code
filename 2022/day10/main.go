package main

import (
	"fmt"

	"github.com/kynrai/advent_of_code/libs/read"
)

func main() {
	one("input.txt")
	two("input.txt")
}

func one(path string) {
	input := read.InputAsStr(path)
	fmt.Println(input)
}

func two(path string) {
	input := read.InputAsStr(path)
	fmt.Println(input)
}
