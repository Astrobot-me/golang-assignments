package main

import "fmt"

// Concepts: slices, maps, range, functions
//
// Task:
// Implement RemoveDuplicates which takes a slice of ints and returns a new
// slice with duplicate values removed, preserving the ORIGINAL order of first
// occurrence. Use a map to track which values have already been seen.

func RemoveDuplicates(nums []int) []int {
	// TODO: implement
	return nil
}

func main() {
	nums := []int{4, 5, 4, 2, 9, 5, 9, 1, 4}
	fmt.Println("Original:", nums)
	fmt.Println("Deduped :", RemoveDuplicates(nums))
}
