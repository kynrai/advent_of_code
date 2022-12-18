package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManhattan(t *testing.T) {
	p1 := Point{X: 8, Y: 7}
	p2 := Point{X: 2, Y: 10}
	assert.Equal(t, 9, Manhattan(p1, p2))
}
