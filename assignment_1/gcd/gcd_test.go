package gcd

import "testing"

/// Note: Tests are not exhaustive, they are just to demonstrate the assignment functions

func TestCalculateGCDArray(t *testing.T) {
	val1 := CalculateGCDArray(4, 64, 32, 120)
	val2 := CalculateGCDArray(11, 33, 99, 198)

	exp1 := 4
	exp2 := 11

	if val1 != exp1 {
		t.Fatalf("GCD Test 1: Expected \"%d\" but got \"%d\"", exp1, val1)
	}
	if val2 != exp2 {
		t.Fatalf("GCD Test 2: Expected \"%d\" but got \"%d\"", exp2, val2)
	}
}
