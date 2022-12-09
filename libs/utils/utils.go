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
