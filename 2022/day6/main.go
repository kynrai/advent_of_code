package main

import (
	"fmt"

	"github.com/kynrai/advent_of_code/libs/read"
)

func main() {
	one("input.txt")
	two("input.txt")
}

func check(s string, n int) bool {
	chars := make(map[rune]struct{})
	for _, v := range s {
		chars[v] = struct{}{}
	}
	return len(chars) == n
}

func one(path string) {
	input := read.InputAsStr(path)[0]
	for i := 4; i < len(input); i++ {
		if check(input[i-4:i], 4) {
			fmt.Println(i)
			break
		}
	}
}

func two(path string) {
	input := read.InputAsStr(path)[0]
	for i := 14; i < len(input); i++ {
		if check(input[i-14:i], 14) {
			fmt.Println(i)
			break
		}
	}
}
