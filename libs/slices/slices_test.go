package slices_test

import (
	"testing"

	"github.com/kynrai/advent_of_code/libs/slices"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, slices.Sum([]int{1, 2, 3}), 6)
}

func TestMax(t *testing.T) {
	assert.Equal(t, slices.Max([]int{1, 2, 3}), 3)
}

func TestMaxN(t *testing.T) {
	assert.Equal(t, slices.MaxN([]int{1, 2, 3}, 2), []int{3, 2})
}

func TestMin(t *testing.T) {
	assert.Equal(t, slices.Min([]int{3, 2, 1}), 1)
}

func TestReverse(t *testing.T) {
	assert.Equal(t, slices.Reverse([]int{3, 2, 1}), []int{1, 2, 3})
}
