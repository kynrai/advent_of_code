package main

import (
	"fmt"

	"github.com/kynrai/advent_of_code/libs/read"
	"github.com/kynrai/advent_of_code/libs/slices"
	. "github.com/kynrai/advent_of_code/libs/utils"
)

func main() {
	one("input.txt")
	two("input.txt")
}

func del[T any](a []T, i int) []T {
	var t T
	copy(a[i:], a[i+1:])
	a[len(a)-1] = t
	return a[:len(a)-1]
}

func insert[T any](a []T, x T, i int) []T {
	return append(a[:i], append([]T{x}, a[i:]...)...)
}

func mod(a, b int) int {
	return (a%b + b) % b
}

type Num struct{ I, N int }

func one(path string) {
	original := read.InputAsInt(path)
	nums := make([]Num, 0, len(original))
	for k, v := range original {
		nums = append(nums, Num{I: k, N: v})
	}
	for oriidx, ori := range original {
		old_idx := 0
		for numIdx, num := range nums {
			if num.I == oriidx {
				old_idx = numIdx
			}
		}
		new_idx := (old_idx + ori)
		switch {
		case new_idx > len(original):
			new_idx %= (len(original) - 1)
		case new_idx <= 0:
			new_idx = len(original) - Abs(new_idx)%(len(original)-1) - 1
		default:
			new_idx %= (len(original) - 1)
		}
		oldNum := nums[old_idx]
		nums = del(nums, old_idx)
		nums = insert(nums, oldNum, new_idx)
	}
	sum := 0
	zeroIdx := 0
	for numIdx, num := range nums {
		if num.N == 0 {
			zeroIdx = numIdx
		}
	}
	for _, v := range []int{1000, 2000, 3000} {
		i := zeroIdx + v
		i = i % len(nums)
		sum += nums[i].N
	}
	fmt.Println(sum)
}

func two(path string) {
	original := read.InputAsInt(path)
	nums := make([]Num, 0, len(original))
	for k, v := range original {
		nums = append(nums, Num{I: k, N: v * 811589153})
	}
	count := 0
	for count < 10 {
		for oriidx := range original {
			old_idx := 0
			for numIdx, num := range nums {
				if num.I == oriidx {
					old_idx = numIdx
				}
			}
			nums = slices.Rotate(nums, -old_idx)
			v := slices.PopLeft(&nums)
			nums = slices.Rotate(nums, -v.N)
			nums = slices.PushLeft(nums, v)
			// new_idx := mod(old_idx+ori, len(original)-1)
			// oldNum := nums[old_idx]
			// nums = del(nums, old_idx)
			// nums = insert(nums, oldNum, new_idx)
		}
		count++
	}

	sum := 0
	zeroIdx := 0
	for numIdx, num := range nums {
		if num.N == 0 {
			zeroIdx = numIdx
		}
	}
	for _, v := range []int{1000, 2000, 3000} {
		i := zeroIdx + v
		i = i % len(nums)
		sum += nums[i].N
	}
	fmt.Println(sum)
}
