package vector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSub(t *testing.T) {
	tests := []struct {
		a      Vector
		b      Vector
		result Vector
	}{
		{
			a:      Vector{X: 1, Y: 1, Z: 0},
			b:      Vector{X: 1, Y: 1, Z: 0},
			result: Vector{X: 0, Y: 0, Z: 0},
		},
		{
			a:      Vector{X: 2, Y: 1, Z: 0},
			b:      Vector{X: 1, Y: 1, Z: 0},
			result: Vector{X: 1, Y: 0, Z: 0},
		},
	}

	for _, test := range tests {
		result := Sub(test.a, test.b)
		assert.Equal(t, test.result, result)
	}
}

func TestCrossProduct(t *testing.T) {
	tests := []struct {
		a      Vector
		b      Vector
		result Vector
	}{
		{
			a:      Vector{X: 1, Y: 1, Z: 0},
			b:      Vector{X: 1, Y: 1, Z: 0},
			result: Vector{X: 0, Y: 0, Z: 0},
		},
		{
			a:      Vector{X: 2, Y: 1, Z: 0},
			b:      Vector{X: 1, Y: 1, Z: 0},
			result: Vector{X: 0, Y: 0, Z: 1},
		},
	}

	for _, test := range tests {
		result := CrossProduct(test.a, test.b)
		assert.Equal(t, test.result, result)
	}
}
