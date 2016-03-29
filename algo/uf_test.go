package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickFindUF(t *testing.T) {
	uf := newQuickFindUF(10)
	uf.union(4, 3)
	uf.union(3, 8)
	uf.union(6, 5)
	uf.union(9, 4)
	uf.union(2, 1)
	uf.union(5, 0)
	uf.union(7, 2)
	uf.union(6, 1)

	assert.Equal(t, 2, uf.count())
}

func TestQuickUnionUF(t *testing.T) {
	uf := newQuickUnionUF(10)
	uf.union(4, 3)
	uf.union(3, 8)
	uf.union(6, 5)
	uf.union(9, 4)
	uf.union(2, 1)
	uf.union(5, 0)
	uf.union(7, 2)
	uf.union(6, 1)

	assert.Equal(t, 2, uf.count())
}

func TestQuickUnionWeightedUF(t *testing.T) {
	uf := newQuickUnionWeightedUF(10)
	uf.union(4, 3)
	uf.union(3, 8)
	uf.union(6, 5)
	uf.union(9, 4)
	uf.union(2, 1)
	uf.union(5, 0)
	uf.union(7, 2)
	uf.union(6, 1)

	assert.Equal(t, 2, uf.count())
}