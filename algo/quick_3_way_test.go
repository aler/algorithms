package algo

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestQuick3WayDutchFlag(t *testing.T) {
    in := []rune("RBWWRWBRRWBR")
    out := []rune("BBBRRRRRWWWW")
    
    quick3WaySort(in)
    assert.Equal(t, string(out), string(in))
}


