package data

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
    q := Queue{first: nil, last: nil}
	q.Enqueue(1)
	q.Enqueue(2)
    assert.Equal(t, 1, q.Dequeue())
    
	q.Enqueue(3)
    assert.Equal(t, 2, q.Dequeue())
    assert.Equal(t, 3, q.Dequeue())
}