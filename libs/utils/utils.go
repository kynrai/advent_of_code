package utils

import (
	"log"
	"math"
	"strconv"
)

func Int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func Abs(i int) int {
	return int(math.Abs(float64(i)))
}

func Chars(s string) []string {
	r := make([]string, len(s))
	for k, v := range s {
		r[k] = string(v)
	}
	return r
}

func AsInts(s string) []int {
	ints := make([]int, len(s))
	for k, v := range s {
		i, _ := strconv.Atoi(string(v))
		ints[k] = i
	}
	return ints
}

func Max(ints ...int) int {
	if len(ints) == 1 {
		return ints[0]
	}
	max := ints[0]
	for _, i := range ints {
		if i > max {
			max = i
		}
	}
	return max
}

func Min(ints ...int) int {
	if len(ints) == 1 {
		return ints[0]
	}
	min := ints[0]
	for _, i := range ints {
		if i < min {
			min = i
		}
	}
	return min
}
