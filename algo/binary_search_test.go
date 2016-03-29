package algo

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestBinarySearchOk(t *testing.T) {
    a := []int{10, 11, 12, 13, 14, 15, 16, 17}
    v, ok := BinarySearch(a, 12)
    assert.Equal(t, 2, v)
    assert.Equal(t, true, ok)
}

func TestBinarySearchFail(t *testing.T) {
    a := []int{10, 11, 12, 13, 14, 15, 16, 17}
    _, ok := BinarySearch(a, 9)
    assert.Equal(t, false, ok)
}

func TestBinarySearchEmpty(t *testing.T) {
    a := []int{}
    _, ok := BinarySearch(a, 9)
    assert.Equal(t, false, ok)
}

func TestBinarySearchOneOk(t *testing.T) {
    a := []int{10}
    v, ok := BinarySearch(a, 10)
    assert.Equal(t, 0, v)
    assert.Equal(t, true, ok)
}

func TestBinarySearchOneFalse(t *testing.T) {
    a := []int{10}
    _, ok := BinarySearch(a, 11)
    assert.Equal(t, false, ok)
}
