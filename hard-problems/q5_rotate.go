package main

import "fmt"

// Concepts: slices, loops
//
// Task:
// Implement RotateLeft which rotates a slice of ints to the LEFT by k
// positions and returns the rotated slice. Handle k > len(nums) and k < 0
// correctly (use modulo). Build the result using loops and slicing only —
// no built-in rotate helpers.
//
// Example: RotateLeft([1,2,3,4,5], 2) -> [3,4,5,1,2]

func RotateLeft(nums []int, k int) []int {
	// TODO: implement
	return nil
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Original:      ", nums)
	fmt.Println("Rotated by 2:  ", RotateLeft(nums, 2))
	fmt.Println("Rotated by 7:  ", RotateLeft(nums, 7))
	fmt.Println("Rotated by -1: ", RotateLeft(nums, -1))
}
