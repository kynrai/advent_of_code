package sliceconv

import (
	"log"
	"strconv"
)

func StrToInt(s []string) []int {
	ints := make([]int, 0, len(s))
	for _, v := range s {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, i)
	}
	return ints
}
