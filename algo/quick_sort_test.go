package algo

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
    in := []rune("QUICKSORTEXAMPLE")
    out := []rune("ACEEIKLMOPQRSTUX")
    quickSort(in)
    assert.Equal(t, out, in)
}