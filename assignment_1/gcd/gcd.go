package gcd

func CalculateGCDArray(integers []int) int {
	res := integers[0]

	for _, i := range integers {
		res = calculateGCD(res, i)
	}

	return res
}

func calculateGCD(a int, b int) int {

	if a < b {
		a, b = b, a
	}
	for b != 0 {
		b, a = a%b, b
	}

	return a
}
