package data

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := stack{first: nil}
	
    s.push(1)
    assert.Equal(t, s.pop(), 1)

	s.push(2)
	s.push(3)
    assert.Equal(t, 3, s.pop())
    assert.Equal(t, 2, s.pop())
}

func TestResizingArrayStack(t *testing.T) {
    s := newResizingArrayStack()
    
    s.push(1)
    assert.Equal(t, 1, s.pop())
    
    s.push(2)
    s.push(3)
    assert.Equal(t, 3, s.pop())
    assert.Equal(t, 2, s.pop())
}

func TestResizingArrayStackShrinking(t *testing.T) {
    s := newResizingArrayStack()
    
    for i := 0; i < 100; i++ {
        s.push(i)
    }
    
    for i := 99; i >= 0; i-- {
        assert.Equal(t, i, s.pop())
    }    
}