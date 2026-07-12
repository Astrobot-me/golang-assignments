package main

import "fmt"

// Concepts: arrays, range
//
// Task:
// Implement ArraySum which returns the sum of all elements in a fixed-size
// array of 5 ints, using range.

func ArraySum(arr [5]int) int {
	// TODO: implement
	sum := 0
	for _, element := range arr {
		sum += element
	}
	return sum
}

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Sum:", ArraySum(arr))
}
