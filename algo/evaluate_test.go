package algo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvaluateSimple(t *testing.T) {
	exp := "( 1 + 2 )"
	assert.Equal(t, 3, evaluate(exp), exp)
}

func TestEvaluateComplex(t *testing.T) {
	exp := "( 1 + ( ( 2 + 3 ) * ( 4 * 5 ) ) )"
	assert.Equal(t, 101, evaluate(exp), exp)
}
