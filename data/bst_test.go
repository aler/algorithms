package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBSTPutNew(t *testing.T) {
	bst := NewBST()
	bst.Put(10, 100)
	bst.Put(12, 120)
	bst.Put(1, 10)
	bst.Put(8, 80)

	assert.Equal(t, 4, bst.Size())
	assert.Equal(t, 12, bst.Max())
	assert.Equal(t, 1, bst.Min())

	get, ok := bst.Get(12)
	assert.Equal(t, 120, get)
	assert.Equal(t, true, ok)

	get, ok = bst.Get(1)
	assert.Equal(t, 10, get)
	assert.Equal(t, true, ok)

	get, ok = bst.Get(8)
	assert.Equal(t, 80, get)
	assert.Equal(t, true, ok)

	get, ok = bst.Get(10)
	assert.Equal(t, 100, get)
	assert.Equal(t, true, ok)
}

func TestBSTPutExisting(t *testing.T) {
	bst := NewBST()
	bst.Put(10, 100)
	bst.Put(12, 120)
	bst.Put(1, 10)
	bst.Put(8, 80)
	bst.Put(12, 210)

	get, ok := bst.Get(12)
	assert.Equal(t, 210, get)
	assert.Equal(t, true, ok)
}

func TestBSTGetNotFound(t *testing.T) {
	bst := NewBST()
	bst.Put(10, 100)
	bst.Put(12, 120)
	bst.Put(1, 10)
	bst.Put(8, 80)

	_, ok := bst.Get(66)
	assert.Equal(t, false, ok)
}

func TestBSTEmpty(t *testing.T) {
	bst := NewBST()
	assert.Equal(t, 0, bst.Size())
}

func TestBSTMinMax(t *testing.T) {
	bst := NewBST()
	bst.Put(10, 100)
	bst.Put(12, 120)
	bst.Put(1, 10)
	bst.Put(8, 80)

	assert.Equal(t, 12, bst.Max())
	assert.Equal(t, 1, bst.Min())
}

func TestBSTFloor(t *testing.T) {
	bst := NewBST()
	bst.Put(10, 100)
	bst.Put(12, 120)
	bst.Put(4, 10)
	bst.Put(8, 80)

    f, ok := bst.Floor(9)
	assert.Equal(t, true, ok)
	assert.Equal(t, 8, f)

    f, ok = bst.Floor(12)
	assert.Equal(t, true, ok)
	assert.Equal(t, 12, f)

    _, ok = bst.Floor(3)
	assert.Equal(t, false, ok)
}

func TestBSTKeys(t *testing.T) {
	bst := NewBST()
	bst.Put(10, 100)
	bst.Put(12, 120)
	bst.Put(4, 10)
	bst.Put(8, 80)
    
    a := []int{4, 8, 10, 12}
    assert.Equal(t, a, bst.Keys())
}

func TestBSTKeysEmpty(t *testing.T) {
	bst := NewBST()
    assert.Equal(t, 0, len(bst.Keys()))
}
