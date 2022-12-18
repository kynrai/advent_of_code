package main

import (
	"fmt"

	"github.com/kynrai/advent_of_code/libs/queue"
	"github.com/kynrai/advent_of_code/libs/read"
	. "github.com/kynrai/advent_of_code/libs/utils"
)

func main() {
	one("input.txt")
	two("input.txt")
}

const size = 100 // big enough to hold all cubes???? maybe

var moves = [6][3]int{
	{1, 0, 0},
	{0, 1, 0},
	{0, 0, 1},
	{-1, 0, 0},
	{0, -1, 0},
	{0, 0, -1},
}

func droplets(path string) map[[3]int]struct{} {
	input := read.InputAsStr(path)
	drops := make(map[[3]int]struct{}, len(input))
	for _, v := range input {
		var x, y, z int
		Ints(v, &x, &y, &z)
		drops[[3]int{x, y, z}] = struct{}{}
	}
	return drops
}

func exposedSides(drops map[[3]int]struct{}) int {
	exposedSides := 0
	for drop := range drops {
		exposed := 6
		for _, m := range moves {
			x, y, z := drop[0]+m[0], drop[1]+m[1], drop[2]+m[2]
			_, ok := drops[[3]int{x, y, z}]
			if ok {
				exposed--
			}
		}
		exposedSides += exposed
	}
	return exposedSides
}

func seenBlocks(drops map[[3]int]struct{}) map[[3]int]struct{} {
	seen := make(map[[3]int]struct{})

	queue := queue.Queue[[3]int]{}
	queue.Push([3]int{-1, -1, -1})

	for {
		cube, hasItem := queue.Pop()
		if !hasItem {
			break
		}

		for _, v := range moves {
			n := [3]int{v[0] + cube[0], v[1] + cube[1], v[2] + cube[2]}
			_, d := drops[n]
			_, s := seen[n]
			if n[0] >= -1 && n[0] < size && n[1] >= -1 && n[1] < size && n[2] >= -1 && n[2] < size &&
				!d && !s {
				queue.Push(n)
				seen[n] = struct{}{}
			}
		}
	}
	return seen
}

func one(path string) {
	drops := droplets(path)
	fmt.Println(exposedSides(drops))
}

func two(path string) {
	drops := droplets(path)
	seen := seenBlocks(drops)
	total := exposedSides(seen) - ((size + 1) * (size + 1) * 6)
	fmt.Println(total)
}
