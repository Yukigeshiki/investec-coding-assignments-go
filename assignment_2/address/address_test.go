package address

import "testing"

/// Note: Tests are not exhaustive, they are just to demonstrate the assignment functions

// TestGetPrettyPrinting tests a. and by extension b.
func TestGetPrettyPrinting(t *testing.T) {

	val0 := (*addrs)[0].GetPrettyPrinting()
	val1 := (*addrs)[1].GetPrettyPrinting()
	val2 := (*addrs)[2].GetPrettyPrinting()

	exp0 := "Physical Address: Address 1, Line 2 - City 1 - Eastern Cape - 1234 - South Africa"
	exp1 := "Postal Address: Not available - City 2 - Not available - 2345 - Lebanon"
	exp2 := "Business Address: Address 3 - City 3 - Not available - 3456 - South Africa"

	if val0 != exp0 {
		t.Fatalf("TestGetPrettyPrinting: Expected \"%s\" but got \"%s\"", exp0, val0)
	}
	if val1 != exp1 {
		t.Fatalf("TestGetPrettyPrinting: Expected \"%s\" but got \"%s\"", exp1, val1)
	}
	if val2 != exp2 {
		t.Fatalf("TestGetPrettyPrinting: Expected \"%s\" but got \"%s\"", exp2, val2)
	}
}

// TestIsValid tests d.
func TestIsValid(t *testing.T) {
	if !(*addrs)[0].IsValid() {
		t.Fatalf("TestIsValid: Expected a valid address but got invalid")
	}
	if (*addrs)[1].IsValid() {
		t.Fatalf("TestIsValid: Expected a invalid address but got valid")
	}
	if (*addrs)[2].IsValid() {
		t.Fatalf("TestIsValid: Expected a invalid address but got valid")
	}
}
