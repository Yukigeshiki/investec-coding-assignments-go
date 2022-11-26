package address

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

/// Note: Tests are not exhaustive, they are just to demonstrate the assignment functions

// TestPrettyPrintAddresses tests b.
func TestPrettyPrintAddresses(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	addrs.PrettyPrintAddresses()
	_ = w.Close()
	result, _ := io.ReadAll(r)
	os.Stdout = stdOut

	assert.Equal(
		t,
		string(result),
		"Physical Address: Address 1, Line 2 - City 1 - Eastern Cape - 1234 - South Africa\n"+
			"Postal Address: Not available - City 2 - Not available - 2345 - Lebanon\n"+
			"Business Address: Address 3 - City 3 - Not available - 3456 - South Africa\n",
	)
}

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
