package main

import (
	"fmt"
	"strings"

	"github.com/kynrai/advent_of_code/libs/read"
	"github.com/kynrai/advent_of_code/libs/slice"
	"github.com/kynrai/advent_of_code/libs/slices"
	. "github.com/kynrai/advent_of_code/libs/utils"
)

func main() {
	one("input.txt")
	two("input.txt")
}

func one(path string) {
	input := read.InputAsStr(path)
	total := 0
	for _, v := range input {
		p := strings.Split(v, "x")
		l, w, h := Int(p[0]), Int(p[1]), Int(p[2])
		a := l * w
		b := w * h
		c := h * l
		total += slice.From(a, b, c).Min() + 2*a + 2*b + 2*c
	}
	fmt.Println(total)
}

func two(path string) {
	input := read.InputAsStr(path)
	total := 0
	for _, v := range input {
		p := strings.Split(v, "x")
		l, w, h := Int(p[0]), Int(p[1]), Int(p[2])
		s := slice.From(l, w, h)
		total += (2 * s.MinN(2).Sum()) + slices.Product(s.Slice())
	}
	fmt.Println(total)
}
