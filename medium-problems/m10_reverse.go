package main

import "fmt"

// Concepts: slices, loops
//
// Task:
// Implement ReverseSlice which returns a NEW slice with the elements of the
// input slice in reverse order (don't modify the original).

func ReverseSlice(nums []int) []int {

	ans := []int{}

	for i := len(nums) - 1; i >= 0; i-- {
		ans = append(ans, nums[i])
	}
	return ans
}

func main() {
	original := []int{1, 2, 3, 4, 5}
	fmt.Println("Original:", original)
	fmt.Println("Reversed:", ReverseSlice(original))
}
