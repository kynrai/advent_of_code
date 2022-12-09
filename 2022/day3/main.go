package main

import (
	"fmt"
	"unicode"

	"github.com/kynrai/advent_of_code/libs/read"
	"github.com/kynrai/advent_of_code/libs/slices"
)

func main() {
	one("input.txt")
	two("input.txt")
}

func pri(c rune) int {
	if unicode.IsUpper(c) {
		return int(c) - 38
	}
	return int(c) - 96
}

func one(path string) {
	vals := read.InputAsRunes(path)
	total := 0
	for _, val := range vals {
		first := val[:len(val)/2]
		last := val[len(val)/2:]
		total += pri(slices.Intersect(first, last)[0])
	}
	fmt.Println(total)
}

func two(path string) {
	groups := read.InputAsChunks(path, 3)
	total := 0
	for _, group := range groups {
		total += pri(slices.Intersect([]rune(group[0]), []rune(group[1]), []rune(group[2]))[0])
	}
	fmt.Println(total)
}
