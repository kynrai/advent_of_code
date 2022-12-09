package main

import (
	"fmt"

	"github.com/kynrai/advent_of_code/libs/read"
)

func main() {
	one("input.txt")
	two("input.txt")
}

type Forest struct {
	rows  int
	cols  int
	input [][]int
	grid  [][]bool
}

func NewForest(ints [][]int) *Forest {
	grid := make([][]bool, len(ints))
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]bool, len(ints[0]))
	}
	return &Forest{
		rows:  len(ints),
		cols:  len(ints[0]),
		input: ints,
		grid:  grid,
	}
}

func visible(s []int) []bool {
	visible := make([]bool, len(s))
	visible[0] = true
	visible[len(visible)-1] = true
	edge1 := s[0]
	edge2 := s[len(s)-1]
	for i := 1; i < len(s)-1; i++ {
		if s[i] > edge1 {
			visible[i] = true
			edge1 = s[i]
		}

	}
	for i := len(s) - 2; i > 0; i-- {
		if s[i] > edge2 {
			visible[i] = true
			edge2 = s[i]
		}
	}
	return visible
}

func (f *Forest) Visible() int {
	for i, row := range f.input {
		f.grid[i] = visible(row)
	}

	for c := 0; c < f.cols; c++ {
		col := []int{}
		for r := range f.input {
			col = append(col, f.input[r][c])
		}

		for r, v := range visible(col) {
			if !f.grid[r][c] {
				f.grid[r][c] = v
			}
		}
	}
	total := 0
	for r := 0; r < f.rows; r++ {
		for c := 0; c < f.cols; c++ {
			if f.grid[c][r] {
				total++
			}
		}
	}
	return total
}

func (f *Forest) ScenicScore() int {
	maxScore := 0
	for r := 1; r < f.rows-1; r++ {
		for c := 1; c < f.cols-1; c++ {
			up, down, left, right := 0, 0, 0, 0
			for i := r - 1; i >= 0; i -= 1 {
				up++
				if f.input[i][c] >= f.input[r][c] {
					break
				}
			}
			for i := r + 1; i < f.cols; i++ {
				down++
				if f.input[i][c] >= f.input[r][c] {
					break
				}
			}
			for i := c - 1; i >= 0; i -= 1 {
				left++
				if f.input[r][i] >= f.input[r][c] {
					break
				}
			}
			for i := c + 1; i < f.rows; i++ {
				right++
				if f.input[r][i] >= f.input[r][c] {
					break
				}
			}
			if left*up*right*down > maxScore {
				maxScore = left * up * right * down
			}
		}
	}

	return maxScore
}

func one(path string) {
	input := read.Grid(path)
	f := NewForest(input)
	fmt.Println(f.Visible())
}

func two(path string) {
	input := read.Grid(path)
	f := NewForest(input)
	fmt.Println(f.ScenicScore())
}
