package utils

import (
	"log"
	"math"
	"regexp"
	"sort"
	"strconv"

	"golang.org/x/exp/constraints"
)

func Int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(s, err)
	}
	return i
}

func Ints(s string, as ...*int) {
	matches := regexp.MustCompile(`-?\d+`).FindAllString(s, -1)
	for k, v := range as {
		*v = Int(matches[k])
	}
}

func Range(s, e int) []int {
	res := make([]int, 0, e-s+1)
	for i := s; i <= e; i++ {
		res = append(res, i)
	}
	return res
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

func Keys[T constraints.Ordered, U any](m map[T]U) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

func Values[T comparable, U any](m map[T]U) []U {
	values := make([]U, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
