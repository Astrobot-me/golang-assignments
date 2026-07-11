package main

import "fmt"

// Concepts: slices, range
//
// Task:
// Implement FindMax which returns the largest value in a slice of ints.
// If the slice is empty, return 0.

func FindMax(nums []int) int {
	// TODO: implement
	return 0
}

func main() {
	fmt.Println("Max:", FindMax([]int{3, 7, 2, 9, 4}))
	fmt.Println("Max of empty:", FindMax([]int{}))
}
