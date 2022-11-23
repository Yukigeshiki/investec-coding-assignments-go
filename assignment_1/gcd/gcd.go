package gcd

// CalculateGCDArray calculates the GCD for an array of integers
func CalculateGCDArray(integers []int) int {
	res := integers[0]

	for _, i := range integers {
		res = calculateGCD(res, i)
	}

	return res
}

// calculateGCD calculates the GCD for 2 integers.
func calculateGCD(a int, b int) int {

	if a < b {
		a, b = b, a
	}
	for b != 0 {
		b, a = a%b, b
	}

	return a
}
