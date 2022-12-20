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

func TestProduct(t *testing.T) {
	assert.Equal(t, slices.Product([]int{3, 2, 1}), 6)
}

func TestUnion(t *testing.T) {
	assert.Equal(t, []int{1, 2, 4, 5, 6, 7}, slices.Union([]int{2, 1}, []int{5, 7}, []int{4, 6}))
}

func TestDelete(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := []int{1, 2, 4, 5, 6, 7}
	assert.Equal(t, b, slices.Delete(a, 2))
}

func TestDeleteElement(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := []int{1, 2, 4, 5, 6, 7}
	assert.Equal(t, b, slices.DeleteElement(a, 3))
}

func TestDeleteElementAll(t *testing.T) {
	a := []int{1, 2, 3, 4, 3, 6, 3}
	b := []int{1, 2, 4, 6}
	assert.Equal(t, b, slices.DeleteElement(a, 3))
}

func TestPopLeft(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	assert.Equal(t, 1, slices.PopLeft(&s))
	assert.Equal(t, []int{2, 3, 4, 5}, s)
}
