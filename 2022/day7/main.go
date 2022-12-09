package main

import (
	"fmt"
	"math"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kynrai/advent_of_code/libs/read"
)

func main() {
	one("input.txt")
	two("input.txt")
}

func isSubDir(path, sub string) bool {
	return path != sub && strings.HasPrefix(sub, path)
}

func totalSubDirs(dirSizes map[string]int, path string) int {
	total := 0
	for k, v := range dirSizes {
		if isSubDir(path, k) {
			// depth := strings.Count(k, "/")
			total += v
		}
	}
	return total
}

func totalDirSize(path string) map[string]int {
	input := read.InputAsStr(path)
	dirs := make(map[string]map[string]int)
	cwd := "/"
	for _, v := range input {
		parts := strings.Split(v, " ")
		switch parts[0] {
		case "$":
			switch parts[1] {
			case "cd":
				switch parts[2] {
				case "..":
					cwd = filepath.Dir(cwd)
				default:
					cwd = filepath.Join(cwd, parts[2])
				}
			case "ls":
				// listing a directory
				dirs[cwd] = make(map[string]int)
			}
			continue
		}
		i, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}
		dirs[cwd][filepath.Join(cwd, parts[1])] = i
	}
	dirSize := 0
	dirSizes := make(map[string]int)
	for d, v := range dirs {
		dirSize = 0
		for _, size := range v {
			dirSize += size
		}
		dirSizes[d] = dirSize
	}
	// find dirs < 100000 and add sub dirs
	totalDirSize := make(map[string]int)
	for k, v := range dirSizes {
		totalDirSize[k] = totalSubDirs(dirSizes, k) + v
	}
	return totalDirSize
}
func one(path string) {
	total := 0
	for _, v := range totalDirSize(path) {
		if v < 100000 {
			total += v
		}
	}
	fmt.Println(total)
}

func two(path string) {
	tds := totalDirSize(path)
	unused := 70000000 - tds["/"]
	min := math.MaxInt
	for _, v := range totalDirSize(path) {
		if (unused+v) > 30000000 && min > v {
			min = v
		}
	}
	fmt.Println(min)
}
