package slices

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func Map[T, U any](s []T, fn func(t T) U) []U {
	res := make([]U, 0, len(s))
	for _, v := range s {
		res = append(res, fn(v))
	}
	return res
}

func Filter[T any](s []T, fn func(T) bool) []T {
	res := make([]T, 0, len(s))
	for _, e := range s {
		if fn(e) {
			res = append(res, e)
		}
	}
	return res
}

type Number interface {
	constraints.Integer | constraints.Float
}

func Sum[T Number](s []T) (sum T) {
	for _, v := range s {
		sum += v
	}
	return sum
}

func Max[T constraints.Ordered](s []T) T {
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	return s[0]
}

func MaxN[T constraints.Ordered](s []T, n uint) []T {
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	return s[:n]
}

func Min[T constraints.Ordered](s []T) T {
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return s[0]
}

func MinN[T constraints.Ordered](s []T, n uint) []T {
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return s[:n]
}

func Intersect[T comparable](as ...[]T) []T {
	maps := make([]map[T]struct{}, 0, len(as))
	for _, a := range as {
		m := make(map[T]struct{})
		for _, r := range a {
			m[r] = struct{}{}
		}
		maps = append(maps, m)
	}
	res := []T{}
	for k := range maps[0] {
		common := true
		for _, m := range maps[1:] {
			if _, ok := m[k]; !ok {
				common = false
			}
		}
		if common {
			res = append(res, k)
		}
	}
	return res
}

func Union[T constraints.Ordered](as ...[]T) []T {
	set := make(map[T]struct{})
	for _, a := range as {
		for _, r := range a {
			set[r] = struct{}{}
		}
	}
	res := make([]T, 0, len(set))
	for k := range set {
		res = append(res, k)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return res
}

func Reverse[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func Product[T Number](s []T) T {
	var prod = s[0]
	for _, v := range s[1:] {
		prod *= v
	}
	return prod
}

func Find[T comparable](a []T, e T) int {
	for i := range a {
		if a[i] == e {
			return i
		}
	}
	return -1
}

func Contains[T comparable](a []T, e T) bool {
	for i := range a {
		if a[i] == e {
			return true
		}
	}
	return false
}

func Delete[T any](a []T, i int) []T {
	var n T
	copy(a[i:], a[i+1:])
	a[len(a)-1] = n // or the zero value of T
	return a[:len(a)-1]
}

func DeleteElement[T comparable](a []T, e T) []T {
	new := []T{}
	for _, v := range a {
		if v != e {
			new = append(new, v)
		}
	}
	return new
}

func Rotate[T any](s []T, i int) []T {
	if len(s) == 0 {
		return s
	}
	r := mod(i, len(s))
	if r > 0 {
		for i := 0; i < r; i++ {
			last := (s)[len(s)-1]
			cutLast := (s)[:len(s)-1]
			s = append([]T{last}, cutLast...)
		}
		return s
	}
	if r < 0 {
		for i := 0; i < r; i++ {
			first := (s)[0]
			cutFirst := (s)[1:]
			s = append(cutFirst, first)
		}
	}
	return s
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func PopLeft[T any](s *[]T) T {
	var t T
	if len(*s) == 0 {
		return t
	}
	element := (*s)[0]
	*s = (*s)[1:]
	return element
}

func PushLeft[T any](s []T, t T) []T {
	return append([]T{t}, s...)
}
