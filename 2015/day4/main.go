package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/kynrai/advent_of_code/libs/read"
)

func main() {
	one("input.txt")
	two("input.txt")
}

func one(path string) {
	input := read.InputAsStr(path)[0]
	i := 0
	for {
		s := fmt.Sprintf("%s%d", input, i)
		v := md5.Sum([]byte(s))
		start := hex.EncodeToString(v[:])[:5]

		if start == "00000" {
			break
		}
		i++
	}
	fmt.Println(i)
}

func two(path string) {
	input := read.InputAsStr(path)[0]
	i := 0
	for {
		s := fmt.Sprintf("%s%d", input, i)
		v := md5.Sum([]byte(s))
		start := hex.EncodeToString(v[:])[:6]

		if start == "000000" {
			break
		}
		i++
	}
	fmt.Println(i)
}
