package algo

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	in := []int{4, 2, 3, 16, 1, 28, 12}
	out := []int{1, 2, 3, 4, 12, 16, 28}
	sort(in)
	assert.Equal(t, out, in)
}

func TestMergeSortSimple(t *testing.T) {
	in := []int{3, 1}
	out := []int{1, 3}
	sort(in)
	assert.Equal(t, out, in)
}

func TestMergeSortSolo(t *testing.T) {
	in := []int{1}
	out := []int{1}
	sort(in)
	assert.Equal(t, out, in)
}

func TestMergeSortEmpty(t *testing.T) {
	in := []int{}
	out := []int{}
	sort(in)
	assert.Equal(t, out, in)
}
