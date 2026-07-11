package main

import "fmt"

// Concepts: maps, slices, loops, conditionals
//
// Task:
// Implement TwoSum which takes a slice of ints and a target sum, and returns
// the INDICES of the two numbers that add up to target. Use a map for an
// O(n) single-pass solution — no nested loops. If no pair is found, return
// (-1, -1). Assume at most one valid pair exists.
//
// Example: TwoSum([2,7,11,15], 9) -> (0, 1)

func TwoSum(nums []int, target int) (int, int) {
	// TODO: implement
	return -1, -1
}

func main() {
	nums := []int{2, 7, 11, 15, 3}

	i, j := TwoSum(nums, 10)
	fmt.Printf("Indices for target 10: %d, %d\n", i, j)

	k, l := TwoSum(nums, 100)
	fmt.Printf("Indices for target 100: %d, %d\n", k, l)
}
