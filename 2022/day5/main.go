package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kynrai/advent_of_code/libs/read"
	"github.com/kynrai/advent_of_code/libs/stack"
)

func main() {
	one("input.txt")
	two("input.txt")
}

var stacks = []stack.Stack[string]{}

func move(n, from, to int) {
	for i := 0; i < n; i++ {
		s, _ := stacks[from].Pop()
		stacks[to].Push(s)
	}
}

func readTop() string {
	res := []string{}
	for _, v := range stacks {
		if s, ok := v.Pop(); ok {
			res = append(res, s)
		}
	}
	return strings.Join(res, "")
}

func one(path string) {
	stacks = []stack.Stack[string]{
		{},
		{"D", "L", "V", "T", "M", "H", "F"},
		{"H", "Q", "G", "J", "C", "T", "N", "P"},
		{"R", "S", "D", "M", "P", "H"},
		{"L", "B", "V", "F"},
		{"N", "H", "G", "L", "Q"},
		{"W", "B", "D", "G", "R", "M", "P"},
		{"G", "M", "N", "R", "C", "H", "L", "Q"},
		{"C", "L", "W"},
		{"R", "D", "L", "Q", "J", "Z", "M", "T"},
	}
	input := read.InputAsStr(path)
	for _, v := range input {
		ins := strings.Split(v, " ")
		n, _ := strconv.Atoi(ins[1])
		from, _ := strconv.Atoi(ins[3])
		to, _ := strconv.Atoi(ins[5])
		move(n, from, to)
	}
	fmt.Println(readTop())
}

func moveN(n, from, to int) {
	toMove := stack.Stack[string]{}
	for i := 0; i < n; i++ {
		s, _ := stacks[from].Pop()
		toMove.Push(s)
	}
	for i := 0; i < n; i++ {
		s, _ := toMove.Pop()
		stacks[to].Push(s)
	}
}

func two(path string) {
	stacks = []stack.Stack[string]{
		{},
		{"D", "L", "V", "T", "M", "H", "F"},
		{"H", "Q", "G", "J", "C", "T", "N", "P"},
		{"R", "S", "D", "M", "P", "H"},
		{"L", "B", "V", "F"},
		{"N", "H", "G", "L", "Q"},
		{"W", "B", "D", "G", "R", "M", "P"},
		{"G", "M", "N", "R", "C", "H", "L", "Q"},
		{"C", "L", "W"},
		{"R", "D", "L", "Q", "J", "Z", "M", "T"},
	}
	input := read.InputAsStr(path)
	for _, v := range input {
		ins := strings.Split(v, " ")
		n, _ := strconv.Atoi(ins[1])
		from, _ := strconv.Atoi(ins[3])
		to, _ := strconv.Atoi(ins[5])
		moveN(n, from, to)
	}
	fmt.Println(readTop())
}
