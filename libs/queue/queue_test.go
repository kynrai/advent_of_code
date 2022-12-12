package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := Queue[int]{}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	n, _ := q.Pop()
	assert.Equal(t, 1, n)
	n, _ = q.Pop()
	assert.Equal(t, 2, n)
	n, _ = q.Pop()
	assert.Equal(t, 3, n)
	_, ok := q.Pop()
	assert.Equal(t, false, ok)
}
