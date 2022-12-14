package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kynrai/advent_of_code/libs/read"
	"github.com/kynrai/advent_of_code/libs/sliceconv"
	. "github.com/kynrai/advent_of_code/libs/utils"
)

func main() {
	one("input.txt")
	two("input.txt")
}

type Monkey struct {
	Items        []int
	Op           func(int) int
	Test         func(int) bool
	Div          int
	True         int
	False        int
	InspectCount int
}

func parseOp(s string) func(int) int {
	parts := strings.Split(strings.Split(strings.TrimSpace(s), ": ")[1], " ")
	op, val := parts[3], parts[4]

	switch op {
	case "+":
		if val == "old" {
			return func(i int) int {
				return i + i
			}
		} else {
			return func(i int) int {
				return i + Int(val)
			}
		}
	case "*":
		if val == "old" {
			return func(i int) int {
				return i * i
			}
		} else {
			return func(i int) int {
				return i * Int(val)
			}
		}
	}
	return nil
}

func parseTest(s string) func(int) bool {
	n := Int(strings.Split(strings.TrimSpace(s), " ")[3])
	return func(i int) bool {
		return i%n == 0
	}
}

func parseMonkeys(path string) []*Monkey {
	input := read.InputAsBatches(path)
	monkeys := []*Monkey{}
	for _, v := range input {
		m := &Monkey{}
		m.Items = sliceconv.StrToInt(strings.Split(strings.Split(v[1], ": ")[1], ", "))
		m.Op = parseOp(v[2])
		m.Test = parseTest(v[3])
		m.Div = Int(strings.Split(strings.TrimSpace(v[3]), " ")[3])
		m.True = Int(strings.Split(strings.TrimSpace(v[4]), " ")[5])
		m.False = Int(strings.Split(strings.TrimSpace(v[5]), " ")[5])
		monkeys = append(monkeys, m)
	}
	return monkeys
}

func one(path string) {
	monkeys := parseMonkeys(path)

	for round := 0; round < 20; round++ {
		for _, m := range monkeys {
			for _, item := range m.Items {
				newWorry := m.Op(item) / 3
				if m.Test(newWorry) {
					monkeys[m.True].Items = append(monkeys[m.True].Items, newWorry)
				} else {
					monkeys[m.False].Items = append(monkeys[m.False].Items, newWorry)
				}
				m.InspectCount++
			}
			m.Items = []int{}
		}
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].InspectCount > monkeys[j].InspectCount
	})

	fmt.Println(monkeys[0].InspectCount * monkeys[1].InspectCount)
}

func two(path string) {
	monkeys := parseMonkeys(path)
	lcm := 1
	for _, m := range monkeys {
		lcm *= m.Div
	}

	for round := 0; round < 10000; round++ {
		for _, m := range monkeys {
			for _, item := range m.Items {
				newWorry := m.Op(item) % lcm
				if m.Test(newWorry) {
					monkeys[m.True].Items = append(monkeys[m.True].Items, newWorry)
				} else {
					monkeys[m.False].Items = append(monkeys[m.False].Items, newWorry)
				}
				m.InspectCount++
			}
			m.Items = []int{}
		}
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].InspectCount > monkeys[j].InspectCount
	})

	fmt.Println(monkeys[0].InspectCount * monkeys[1].InspectCount)
}
