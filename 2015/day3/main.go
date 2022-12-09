package main

import (
	"fmt"

	"github.com/kynrai/advent_of_code/libs/read"
)

func main() {
	one("input.txt")
	two("input.txt")
}

type c struct {
	X, Y int
}

func one(path string) {
	input := read.InputAsStr(path)[0]
	pos := c{0, 0}
	been := make(map[c]struct{})
	been[pos] = struct{}{}
	for _, v := range input {
		switch v {
		case '>':
			pos.X += 1
		case '<':
			pos.X -= 1
		case '^':
			pos.Y += 1
		case 'v':
			pos.Y -= 1
		}
		been[pos] = struct{}{}
	}
	fmt.Println(len(been))
}

func two(path string) {
	input := read.InputAsStr(path)[0]
	santa := c{0, 0}
	robo := c{0, 0}
	been := make(map[c]struct{})
	been[santa] = struct{}{}
	santaTurn := true
	for _, v := range input {
		if santaTurn {
			switch v {
			case '>':
				santa.X += 1
			case '<':
				santa.X -= 1
			case '^':
				santa.Y += 1
			case 'v':
				santa.Y -= 1
			}
			been[santa] = struct{}{}
			santaTurn = !santaTurn
		} else {
			switch v {
			case '>':
				robo.X += 1
			case '<':
				robo.X -= 1
			case '^':
				robo.Y += 1
			case 'v':
				robo.Y -= 1
			}
			been[robo] = struct{}{}
			santaTurn = !santaTurn
		}

	}
	fmt.Println(len(been))
}
