package main

import (
	"fmt"
	"strings"

	"github.com/kynrai/advent_of_code/libs/read"
	. "github.com/kynrai/advent_of_code/libs/utils"
	"golang.org/x/exp/slices"
)

func main() {
	one("input.txt")
	two("input.txt")
}

func buildReg(path string) []int {
	input := read.InputAsStr(path)
	register := []int{1, 1}
	currentRegVal := 1
	cycles := 0
	for _, v := range input {
		if strings.HasPrefix(v, "addx") {
			val := Int(strings.Split(v, " ")[1])
			cycles++
			register = append(register, currentRegVal)
			cycles++
			currentRegVal += val
			register = append(register, currentRegVal)
		} else {
			cycles++
			register = append(register, currentRegVal)
		}
	}
	return register
}

func one(path string) {
	register := buildReg(path)
	total := 0
	for i := 20; i <= 220; i += 40 {
		total += i * register[i]
	}
	fmt.Println(total)
}

func two(path string) {
	register := buildReg(path)
	crt := ""
	for cycle := 1; cycle < len(register); cycle++ {
		if len(crt) == 40 {
			fmt.Println(crt)
			crt = ""
		}
		v := register[cycle]
		spritPos := []int{v - 1, v, v + 1}
		if slices.Contains(spritPos, len(crt)) {
			crt += "#"
		} else {
			crt += "."
		}
	}
}
