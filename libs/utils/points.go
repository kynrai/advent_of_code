package utils

import (
	"fmt"

	"github.com/kynrai/advent_of_code/libs/slice"
)

type Point struct {
	X, Y int
}

func Manhattan(p1, p2 Point) int {
	return Abs(p1.Y-p2.Y) + Abs(p1.X-p2.X)
}

func Draw(grid map[Point]string) {
	xs := []int{}
	ys := []int{}
	for point := range grid {
		xs = append(xs, point.X)
		ys = append(ys, point.Y)
	}
	xss := slice.From(xs...)
	minX, maxX := xss.Min(), xss.Max()
	yss := slice.From(ys...)
	minY, maxY := yss.Min(), yss.Max()

	for y := minY; y <= maxY; y++ {
		row := []string{}
		for x := minX; x <= maxX; x++ {
			if v, ok := grid[Point{X: x, Y: y}]; ok {
				row = append(row, v)
			} else {
				row = append(row, ".")
			}
		}
		fmt.Println(row)
	}
}
