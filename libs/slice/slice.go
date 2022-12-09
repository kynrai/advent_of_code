package slice

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type S[T constraints.Ordered] []T

func From[T constraints.Ordered](t ...T) *S[T] {
	s := make(S[T], len(t))
	copy(s, t)
	return &s
}

func (s *S[T]) Max() T {
	var max T
	if len(*s) > 0 {
		max = (*s)[0]
	}
	for i := 0; i < len(*s); i++ {
		if (*s)[i] > max {
			max = (*s)[i]
		}
	}
	return max
}

func (s *S[T]) MaxN(n int) []T {
	mins := make([]T, len(*s))
	copy(mins, *s)
	sort.Slice(mins, func(i, j int) bool {
		return mins[i] > mins[j]
	})
	return mins[:n]
}

func (s *S[T]) Min() T {
	var min T
	if len(*s) > 0 {
		min = (*s)[0]
	}
	for i := 0; i < len(*s); i++ {
		if (*s)[i] < min {
			min = (*s)[i]
		}
	}
	return min
}

func (s *S[T]) MinN(n int) *S[T] {
	mins := make([]T, len(*s))
	copy(mins, *s)
	sort.Slice(mins, func(i, j int) bool {
		return mins[i] < mins[j]
	})
	return From(mins[:n]...)
}

func (s *S[T]) Sum() T {
	var sum T
	for i := 0; i < len(*s); i++ {
		sum += (*s)[i]
	}
	return sum
}

func (s *S[T]) Count(t T) int {
	var count int
	for i := 0; i < len(*s); i++ {
		if (*s)[i] == t {
			count++
		}
	}
	return count
}

func (s *S[T]) Pop() T {
	var t T
	if len(*s) == 0 {
		return t
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element
}

func (s *S[T]) Push(t T) *S[T] {
	s1 := From(t)
	*s1 = append(*s1, *s...)
	return s1
}

func (s *S[T]) Slice() []T {
	s1 := make([]T, len(*s))
	copy(s1, *s)
	return s1
}
