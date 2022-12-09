package main

import (
	"fmt"
	"strings"

	"github.com/kynrai/advent_of_code/libs/read"
	. "github.com/kynrai/advent_of_code/libs/util"
)

func main() {
	one("input.txt")
	two("input.txt")
}

type coord struct {
	X, Y int
}

func moveTail(hx, hy, tx, ty int) (int, int) {
	dx := hx - tx
	dy := hy - ty

	moveX := 1
	if dx < 0 {
		moveX = -1
	} else if dx == 0 {
		moveX = 0
	}

	moveY := 1
	if dy < 0 {
		moveY = -1
	} else if dy == 0 {
		moveY = 0
	}

	if Abs(dx) > 1 {
		return tx + moveX, ty + moveY
	}
	if Abs(dy) > 1 {
		return tx + moveX, ty + moveY
	}
	return tx, ty
}

func one(path string) {
	input := read.InputAsStr(path)
	visited := make(map[coord]struct{})
	visited[coord{0, 0}] = struct{}{}
	hx, hy := 0, 0
	tx, ty := 0, 0

	for _, moveStr := range input {
		move := strings.Split(moveStr, " ")
		dir := move[0]
		dist := Int(move[1])

		for i := 0; i < dist; i++ {
			switch dir {
			case "U":
				hy++
			case "D":
				hy--
			case "L":
				hx--
			case "R":
				hx++
			}
			tx, ty = moveTail(hx, hy, tx, ty)
			visited[coord{tx, ty}] = struct{}{}
		}
	}
	println(len(visited))
}

func two(path string) {
	input := read.InputAsStr(path)
	knots := make([]coord, 10)
	visited := make(map[coord]struct{})
	visited[coord{0, 0}] = struct{}{}
	for _, moveStr := range input {
		move := strings.Split(moveStr, " ")
		dir := move[0]
		dist := Int(move[1])

		for i := 0; i < dist; i++ {
			switch dir {
			case "R":
				knots[0].X++
			case "L":
				knots[0].X--
			case "U":
				knots[0].Y++
			case "D":
				knots[0].Y--
			}

			for i := 1; i < 10; i++ {
				x, y := moveTail(knots[i-1].X, knots[i-1].Y, knots[i].X, knots[i].Y)
				knots[i] = coord{x, y}
			}
			visited[knots[9]] = struct{}{}
		}
	}

	fmt.Println(len(visited))
}
