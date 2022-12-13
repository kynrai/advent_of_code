package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"

	"github.com/kynrai/advent_of_code/libs/read"
	"github.com/kynrai/advent_of_code/libs/slice"
	. "github.com/kynrai/advent_of_code/libs/utils"
)

func main() {
	one("input.txt")
	two("input.txt")
}

func isKind(v any, k reflect.Kind) bool {
	return reflect.TypeOf(v).Kind() == k
}

func inOrder(l, r any) (bool, bool) {
	// fmt.Printf("L %t, R %t\n", l, r)
	if isKind(l, reflect.Float64) && isKind(r, reflect.Float64) {
		lf, _ := l.(float64)
		rf, _ := r.(float64)
		if lf == rf {
			return false, true
		}
		return lf < rf, false
	}

	if isKind(l, reflect.Int) && isKind(r, reflect.Int) {
		lf, _ := l.(int)
		rf, _ := r.(int)
		if lf == rf {
			return false, true
		}
		return lf < rf, false
	}

	if isKind(l, reflect.Slice) && isKind(r, reflect.Slice) {
		lf, _ := l.([]any)
		rf, _ := r.([]any)
		for i := 0; i < Min(len(lf), len(rf)); i++ {
			if ord, eq := inOrder(lf[i], rf[i]); !eq {
				return ord, false
			}
		}
		return inOrder(len(lf), len(rf))
	}

	if isKind(l, reflect.Float64) {
		return inOrder([]any{l}, r)
	} else if isKind(r, reflect.Float64) {
		return inOrder(l, []any{r})
	}

	if isKind(l, reflect.Int) {
		return inOrder([]any{l}, r)
	}
	return inOrder(l, []any{r})
}

func strToList(s string) []any {
	l := make([]any, 0)
	json.Unmarshal([]byte(s), &l)
	return l
}

func one(path string) {
	input := read.InputAsBatches(path)
	orderIdx := []int{}
	for idx, pair := range input {
		leftLists, rightLists := strToList(pair[0]), strToList(pair[1])
		order, _ := inOrder(leftLists, rightLists)
		// fmt.Println(leftLists, rightLists, order)
		if order {
			orderIdx = append(orderIdx, idx+1)
			continue
		}

	}
	fmt.Println(slice.From(orderIdx...).Sum())
}

func two(path string) {
	input := read.InputAsStr(path)
	entries := []any{}
	for _, v := range input {
		if v == "" {
			continue
		}
		entries = append(entries, strToList(v))
	}
	entries = append(entries, strToList("[[2]]"), strToList("[[6]]"))

	sort.Slice(entries, func(i, j int) bool {
		ord, _ := inOrder(entries[i], entries[j])
		return ord
	})
	for k, v := range entries {
		fmt.Println(k+1, v)
	}
}
