package slice_test

import (
	"testing"

	"github.com/kynrai/advent_of_code/libs/slice"
	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	s := slice.From(1, 2, 3, 4, 5)
	assert.Equal(t, 5, s.Max())

	s = slice.From(5, 4, 3, 2, 1)
	assert.Equal(t, 5, s.Max())

	s = slice.From(5, 4, 10, 2, 1)
	assert.Equal(t, 10, s.Max())

	s = slice.From(5, 1)
	assert.Equal(t, 5, s.Max())

	s = slice.From(5, 10)
	assert.Equal(t, 10, s.Max())

	s = slice.From(10)
	assert.Equal(t, 10, s.Max())

	s = slice.From[int]()
	assert.Equal(t, 0, s.Max())
}

func TestMin(t *testing.T) {
	s := slice.From(1, 2, 3, 4, 5)
	assert.Equal(t, 1, s.Min())

	s = slice.From(5, 4, 3, 2, 1)
	assert.Equal(t, 1, s.Min())

	s = slice.From(5, -4, 10, 2, 1)
	assert.Equal(t, -4, s.Min())

	s = slice.From(-5, 1)
	assert.Equal(t, -5, s.Min())

	s = slice.From(5, -10)
	assert.Equal(t, -10, s.Min())

	s = slice.From(10)
	assert.Equal(t, 10, s.Min())

	s = slice.From[int]()
	assert.Equal(t, 0, s.Min())
}

func TestSum(t *testing.T) {
	s := slice.From(1, 2, 3, 4, 5)
	assert.Equal(t, 15, s.Sum())

	s = slice.From(5, -4, 10, 2, 1)
	assert.Equal(t, 14, s.Sum())

	s = slice.From(-5, 1)
	assert.Equal(t, -4, s.Sum())

	s = slice.From(10)
	assert.Equal(t, 10, s.Sum())

	s = slice.From[int]()
	assert.Equal(t, 0, s.Sum())
}

func TestCount(t *testing.T) {
	s := slice.From(1, 1, 3, 1, 5)
	assert.Equal(t, 3, s.Count(1))

	s = slice.From(10)
	assert.Equal(t, 1, s.Count(10))

	s = slice.From[int]()
	assert.Equal(t, 0, s.Count(1))
}

func TestPop(t *testing.T) {
	s := slice.From(1, 2, 3, 4, 5)
	assert.Equal(t, 5, s.Pop())

	s = slice.From(5, 4, 3, 2, 1)
	assert.Equal(t, 1, s.Pop())

	s = slice.From(5, 4, 10, 2, 1)
	assert.Equal(t, 1, s.Pop())

	s = slice.From(5, 1)
	assert.Equal(t, 1, s.Pop())

	s = slice.From(5, 10)
	assert.Equal(t, 10, s.Pop())

	s = slice.From(10)
	assert.Equal(t, 10, s.Pop())

	s = slice.From[int]()
	assert.Equal(t, 0, s.Pop())
}

func TestPush(t *testing.T) {
	s := slice.From(1, 2, 3, 4, 5)
	assert.Equal(t, slice.From(0, 1, 2, 3, 4, 5), s.Push(0))
}

func TestMinN(t *testing.T) {
	s := slice.From(5, 4, 3, 2, 1)
	assert.Equal(t, []int{1, 2}, s.MinN(2))
}

func TestMaxN(t *testing.T) {
	s := slice.From(1, 2, 3, 4, 5)
	assert.Equal(t, []int{5, 4}, s.MaxN(2))
}
