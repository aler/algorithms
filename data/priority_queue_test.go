package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriortyQueue(t *testing.T) {
	q := newPriorityQueue()
	q.insert(10)
	q.insert(12)
	q.insert(2)
	q.insert(5)
	q.insert(7)
	q.insert(1)
	q.insert(4)

	assert.Equal(t, 12, q.deleteMax())
	assert.Equal(t, 10, q.deleteMax())
	assert.Equal(t, 7, q.deleteMax())
	assert.Equal(t, 5, q.max())

	q.insert(1)
	assert.Equal(t, 5, q.max())
	q.insert(9)
	assert.Equal(t, 9, q.max())
}
