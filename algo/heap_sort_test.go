package algo

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestHeapSort(t *testing.T) {
    in := []rune("SORTEXAMPLE")
    out := []rune("AEELMOPRSTX")
    
    heapSort(in)
    assert.Equal(t, string(out), string(in))
}