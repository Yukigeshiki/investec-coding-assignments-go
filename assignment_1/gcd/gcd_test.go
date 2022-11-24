package gcd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/// Note: Tests are not exhaustive, they are just to demonstrate the assignment functions

func TestCalculateGCDArray(t *testing.T) {
	assert.Equal(t, CalculateGCDArray(4, 64, 32, 120), 4)
	assert.Equal(t, CalculateGCDArray(11, 33, 99, 198), 11)
}
