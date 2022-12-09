package main

import (
	"fmt"
	"strings"

	"github.com/kynrai/advent_of_code/libs/read"
)

func main() {
	one("input.txt")
	two("input.txt")
}

func one(path string) {
	vals := read.InputAsStr(path)
	score := 0
	for _, v := range vals {
		p := strings.Split(v, " ")
		switch {
		case p[0] == "A" && p[1] == "X":
			score += 4
		case p[0] == "A" && p[1] == "Y":
			score += 8
		case p[0] == "A" && p[1] == "Z":
			score += 3
		case p[0] == "B" && p[1] == "X":
			score += 1
		case p[0] == "B" && p[1] == "Y":
			score += 5
		case p[0] == "B" && p[1] == "Z":
			score += 9
		case p[0] == "C" && p[1] == "X":
			score += 7
		case p[0] == "C" && p[1] == "Y":
			score += 2
		case p[0] == "C" && p[1] == "Z":
			score += 6
		}
	}
	fmt.Println(score)
}

func two(path string) {
	vals := read.InputAsStr(path)
	score := 0
	for _, v := range vals {
		p := strings.Split(v, " ")
		switch {
		case p[0] == "A" && p[1] == "X":
			score += 3
		case p[0] == "A" && p[1] == "Y":
			score += 4
		case p[0] == "A" && p[1] == "Z":
			score += 8
		case p[0] == "B" && p[1] == "X":
			score += 1
		case p[0] == "B" && p[1] == "Y":
			score += 5
		case p[0] == "B" && p[1] == "Z":
			score += 9
		case p[0] == "C" && p[1] == "X":
			score += 2
		case p[0] == "C" && p[1] == "Y":
			score += 6
		case p[0] == "C" && p[1] == "Z":
			score += 7
		}
	}
	fmt.Println(score)
}
