package gcd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/// Note: Tests are not exhaustive, they are just to demonstrate the assignment functions

func TestCalculateGCDArray(t *testing.T) {
	assert.Equal(t, 6, CalculateGCDArray(6))
	assert.Equal(t, 4, CalculateGCDArray(4, 64, 32, 120))
	assert.Equal(t, 11, CalculateGCDArray(11, 33, 99, 198))
}
