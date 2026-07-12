package main

import "fmt"

// Concepts: loops
//
// Task:
// Implement SumOfDigits which returns the sum of the digits of a
// non-negative integer using a for loop (no string conversion).
// Example: SumOfDigits(1234) -> 10

func SumOfDigits(n int) int {
	// TODO: implement
	var sum int
	sum = 0
	for n > 0 {
		sum = sum + n%10
		n = n / 10
	}
	return sum
}

func main() {
	nums := []int{1234, 0, 9, 98765}
	for _, n := range nums {
		fmt.Printf("SumOfDigits(%d) = %d\n", n, SumOfDigits(n))
	}
}
