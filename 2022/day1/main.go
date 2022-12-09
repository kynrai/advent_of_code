package main

import (
	"fmt"

	"github.com/kynrai/advent_of_code/libs/read"
	"github.com/kynrai/advent_of_code/libs/sliceconv"
	"github.com/kynrai/advent_of_code/libs/slices"
)

func main() {
	one("input.txt")
	two("input.txt")
}

func vals(path string) []int {
	vals := []int{}
	for _, batch := range read.InputAsBatches(path) {
		vals = append(vals, slices.Sum(sliceconv.StrToInt(batch)))
	}
	return vals
}

func one(path string) {
	fmt.Println(slices.Max(vals(path)))
}

func two(path string) {
	fmt.Println(slices.Sum(slices.MaxN(vals(path), 3)))
}
