package gcd

// CalculateGCDArray calculates the GCD for an array of integers
func CalculateGCDArray(args ...int) (res int) {
	if len(args) == 0 {
		return -1
	}

	res = args[0]
	for _, i := range args {
		res = calculateGCD(res, i)
	}

	return
}

// calculateGCD calculates the GCD for 2 integers.
func calculateGCD(a, b int) int {

	if a < b {
		a, b = b, a
	}
	for b != 0 {
		b, a = a%b, b
	}

	return a
}
