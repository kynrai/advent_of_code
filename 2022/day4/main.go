package main

import (
	"fmt"
	"strings"

	"github.com/kynrai/advent_of_code/libs/read"
	. "github.com/kynrai/advent_of_code/libs/utils"
)

func main() {
	one("input.txt")
	two("input.txt")
}

func limits(s string) (int, int) {
	r := strings.Split(s, "-")
	return Int(r[0]), Int(r[1])
}

func one(path string) {
	input := read.InputAsSplitLines(path, ",")
	total := 0
	for _, v := range input {
		as, ae := limits(v[0])
		bs, be := limits(v[1])
		if (as >= bs && ae <= be) || (bs >= as && be <= ae) {
			total += 1
		}
	}
	fmt.Println(total)
}

func two(path string) {
	input := read.InputAsSplitLines(path, ",")
	total := 0
	for _, v := range input {
		as, ae := limits(v[0])
		bs, be := limits(v[1])
		if (ae >= bs) && (as <= be) {
			total += 1
		}
	}
	fmt.Println(total)
}
