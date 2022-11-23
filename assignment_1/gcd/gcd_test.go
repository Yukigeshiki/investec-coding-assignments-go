package gcd

import "testing"

func TestCalculateGCDArray(t *testing.T) {
	val1 := CalculateGCDArray([]int{4, 64, 32, 120})
	val2 := CalculateGCDArray([]int{11, 33, 99, 198})

	exp1 := 4
	exp2 := 11

	if val1 != exp1 {
		t.Fatalf("GCD Test 1: Expected %v but got %v", exp1, val1)
	}
	if val2 != exp2 {
		t.Fatalf("GCD Test 2: Expected %v but got %v", exp2, val2)
	}
}
