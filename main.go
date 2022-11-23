package main

import (
	"fmt"
	"robothouse.io/investec-coding-assignments/assignment_1/gcd"
)

func main() {
	testVal := []int{11, 33, 99, 198}

	fmt.Println(gcd.CalculateGCDArray(testVal))
}
