package main

import (
	"fmt"
	"strings"

	"github.com/kynrai/advent_of_code/libs/read"
	"github.com/kynrai/advent_of_code/libs/utils"
)

func main() {
	one("input.txt")
	two("input.txt")
}

type Monkey struct {
	Name      string
	Answer    int
	HasAnswer bool
	Task      func(int, int) int
	Requires  []string
}

var (
	add = func(a, b int) int { return a + b }
	sub = func(a, b int) int { return a - b }
	mul = func(a, b int) int { return a * b }
	div = func(a, b int) int { return a / b }
)

func getTask(s string) func(int, int) int {
	switch s {
	case "+":
		return add
	case "-":
		return sub
	case "*":
		return mul
	case "/":
		return div
	}
	return nil
}

func one(path string) {
	input := read.InputAsStr(path)
	monkeys := map[string]*Monkey{}
	for _, in := range input {
		parts := strings.Split(in, ": ")
		name := parts[0]
		// first add all monkeys to the map, with the answer if they have one
		m := &Monkey{Name: name}
		if len(in) != 17 {
			m.Answer = utils.Int(parts[1])
			m.HasAnswer = true
		} else {
			// parents and function if they do not
			reqs := strings.Split(parts[1], " ")
			m.Requires = []string{reqs[0], reqs[2]}
			m.Task = getTask(reqs[1])
		}
		monkeys[name] = m
	}

	// go through all monkeys forever until root has an answer
	for !monkeys["root"].HasAnswer {
		for _, v := range monkeys {
			if !v.HasAnswer {
				m1 := monkeys[v.Requires[0]]
				m2 := monkeys[v.Requires[1]]
				if m1.HasAnswer && m2.HasAnswer {
					v.Answer = v.Task(m1.Answer, m2.Answer)
					v.HasAnswer = true
				}
			}
		}
	}

	fmt.Println(monkeys["root"].Answer)
}

func two(path string) {

}
