package main

import "fmt"

// Concepts: functions (multiple return values), loops
//
// Task:
// Implement MinMax which returns both the minimum and maximum values in a
// slice of ints as two separate return values.

func MinMax(nums []int) (min int, max int) {
	// TODO: implement
	return 0, 0
}

func main() {
	mn, mx := MinMax([]int{5, 1, 9, 3, 7})
	fmt.Printf("Min: %d, Max: %d\n", mn, mx)
}
