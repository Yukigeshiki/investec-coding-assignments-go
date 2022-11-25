package address

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/// Note: Tests are not exhaustive, they are just to demonstrate the assignment functions

// TestGetPrettyPrinting tests a. and by extension b.
func TestGetPrettyPrinting(t *testing.T) {
	assert.Equal(
		t,
		addrs[0].GetPrettyPrinting(),
		"Physical Address: Address 1, Line 2 - City 1 - Eastern Cape - 1234 - South Africa",
	)
	assert.Equal(
		t,
		addrs[1].GetPrettyPrinting(),
		"Postal Address: Not available - City 2 - Not available - 2345 - Lebanon",
	)
	assert.Equal(
		t,
		addrs[2].GetPrettyPrinting(),
		"Business Address: Address 3 - City 3 - Not available - 3456 - South Africa",
	)
}

// TestIsValid tests d.
func TestIsValid(t *testing.T) {
	assert.True(t, addrs[0].IsValid())
	assert.False(t, addrs[1].IsValid())
	assert.False(t, addrs[2].IsValid())
}
