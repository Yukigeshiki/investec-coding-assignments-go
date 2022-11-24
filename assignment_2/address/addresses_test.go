package address

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/// Note: Tests are not exhaustive, they are just to demonstrate the assignment functions

// TestValidateAddresses tests e.
func TestValidateAddresses(t *testing.T) {
	assert.Equal(
		t,
		[]string{
			"Address for ID: 2 is invalid. Validation errors: [You must include valid address details (line 1 and/or 2 must be filled in)]",
			"Address for ID: 3 is invalid. Validation errors: [You must include a province if your country is ZA]",
		},
		addrs.ValidateAddresses(),
	)
}
