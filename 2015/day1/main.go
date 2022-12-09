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
	input := read.InputAsStr(path)[0]
	floor := 0
	for _, v := range input {
		switch v {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	fmt.Println(floor)
}

func two(path string) {
	input := read.InputAsStr(path)[0]
	floor := 0
	for k, v := range input {
		switch v {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor == -1 {
			fmt.Println(k + 1)
			break
		}
	}
}
