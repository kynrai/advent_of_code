package main

import (
	"fmt"

	"github.com/kynrai/advent_of_code/libs/queue"
	"github.com/kynrai/advent_of_code/libs/read"
)

func main() {
	// one("input.txt")
	two("input.txt")
}

type node struct {
	x, y, steps int
	val         rune
}

type point struct {
	x, y int
}

// func canMoveUp(x,y int) [][]
func one(path string) {
	grid := read.InputAsRunes(path)
	searched := make(map[point]struct{})
	rows := len(grid)
	cols := len(grid[0])
	sq := &queue.Queue[node]{} //search queue
	sq.Push(node{
		x:     20,
		y:     0,
		steps: 0,
		val:   'a'}) // test 0,0,0 prod 20,0,0
	grid[20][112] = 'z'
	for {
		// if sq.IsEmpty() {
		// 	break
		// }
		// current pos and steps
		cur, ok := sq.Pop()
		if !ok {
			break
		}

		// if cur.val == 'z' {
		// 	fmt.Println(cur.steps)
		// 	break
		// }
		x, y, steps := cur.x, cur.y, cur.steps
		// if x == 2 && y == 5 {
		if x == 20 && y == 112 {
			fmt.Println(x, y, steps, string(cur.val))
		}
		// find steps possible from current x, y
		if x+1 < rows {
			if val := grid[x+1][y]; cur.val+1 >= val {
				n := node{
					x:     x + 1,
					y:     y,
					steps: steps + 1,
					val:   val}
				p := point{n.x, n.y}
				if _, ok := searched[p]; !ok {
					searched[p] = struct{}{}
					sq.Push(n)
				}
			}
		}
		if x-1 >= 0 {
			if val := grid[x-1][y]; cur.val+1 >= val {
				n := node{
					x:     x - 1,
					y:     y,
					steps: steps + 1,
					val:   val}
				p := point{n.x, n.y}
				if _, ok := searched[p]; !ok {
					searched[p] = struct{}{}
					sq.Push(n)
				}
			}
		}
		if y+1 < cols {
			if val := grid[x][y+1]; cur.val+1 >= val {
				n := node{
					x:     x,
					y:     y + 1,
					steps: steps + 1,
					val:   val}
				p := point{n.x, n.y}
				if _, ok := searched[p]; !ok {
					searched[p] = struct{}{}
					sq.Push(n)
				}
			}
		}
		if y-1 >= 0 {
			if val := grid[x][y-1]; cur.val+1 >= val {
				n := node{
					x:     x,
					y:     y - 1,
					steps: steps + 1,
					val:   val}
				p := point{n.x, n.y}
				if _, ok := searched[p]; !ok {
					searched[p] = struct{}{}
					sq.Push(n)
				}
			}
		}
	}
}

func two(path string) {
	grid := read.InputAsRunes(path)
	searched := make(map[point]struct{})
	rows := len(grid)
	cols := len(grid[0])
	sq := &queue.Queue[node]{} //search queue
	sq.Push(node{
		x:     20,
		y:     112,
		steps: 0,
		val:   'z'}) // test 2,5,0 prod 20,122,0
	for {
		// current pos and steps
		cur, ok := sq.Pop()
		if !ok {
			break
		}

		x, y, steps := cur.x, cur.y, cur.steps
		fmt.Println(x, y, steps, string(cur.val))
		// if x == 2 && y == 5 {
		if cur.val == 'a' {
			fmt.Println(x, y, steps, string(cur.val))
			break
		}
		// find steps possible from current x, y
		if x+1 < rows {
			val := grid[x+1][y]
			if cur.val-1 <= val {
				n := node{
					x:     x + 1,
					y:     y,
					steps: steps + 1,
					val:   val}
				p := point{n.x, n.y}
				if _, ok := searched[p]; !ok {
					searched[p] = struct{}{}
					sq.Push(n)
				}
			}
		}
		if x-1 >= 0 {
			val := grid[x-1][y]
			if cur.val-1 <= val {
				n := node{
					x:     x - 1,
					y:     y,
					steps: steps + 1,
					val:   val}
				p := point{n.x, n.y}
				if _, ok := searched[p]; !ok {
					searched[p] = struct{}{}
					sq.Push(n)
				}
			}
		}
		if y+1 < cols {
			val := grid[x][y+1]
			if cur.val-1 <= val {
				n := node{
					x:     x,
					y:     y + 1,
					steps: steps + 1,
					val:   val}
				p := point{n.x, n.y}
				if _, ok := searched[p]; !ok {
					searched[p] = struct{}{}
					sq.Push(n)
				}
			}
		}
		if y-1 >= 0 {
			val := grid[x][y-1]
			if cur.val-1 <= val {
				n := node{
					x:     x,
					y:     y - 1,
					steps: steps + 1,
					val:   val}
				p := point{n.x, n.y}
				if _, ok := searched[p]; !ok {
					searched[p] = struct{}{}
					sq.Push(n)
				}
			}
		}
	}
}
