package address

import (
	"reflect"
	"testing"
)

/// Note: Tests are not exhaustive, they are just to demonstrate the assignment functions

// TestValidateAddresses tests e.
func TestValidateAddresses(t *testing.T) {
	val := addrs.ValidateAddresses()
	exp := []string{
		"Address for ID: 2 is invalid. Validation errors: [You must include valid address details (line 1 and/or 2 must be filled in)]",
		"Address for ID: 3 is invalid. Validation errors: [You must include a province if your country is ZA]",
	}

	if !reflect.DeepEqual(exp, val) {
		t.Fatalf("TestValidateAddresses: Expected \"%v\" but got \"%v\"", exp, val)
	}
}
