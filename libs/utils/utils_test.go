package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInts(t *testing.T) {
	testStr := `Sensor at x=1999258, y=1017714: closest beacon is at x=3332075, y=-572515`
	var x1, y1, x2, y2 int
	Ints(testStr, &x1, &y1, &x2, &y2)
	assert.Equal(t, 1999258, x1)
	assert.Equal(t, 1017714, y1)
	assert.Equal(t, 3332075, x2)
	assert.Equal(t, -572515, y2)
}
