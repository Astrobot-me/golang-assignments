package main

import "fmt"

// Concepts: arrays, loops, conditionals
//
// Task:
// Implement SecondMinMax which takes a fixed-size array of 10 ints and returns
// the second largest and second smallest DISTINCT values in the array.
// Do NOT use the sort package. Aim for a single pass over the array.
// If the array has fewer than 2 distinct values, return (0, 0).

func SecondMinMax(arr [10]int) (int, int) {
	// TODO: implement
	maxE := arr[0]
	sMaxE := arr[0]
	minE := arr[0]
	sMinE := arr[0]

	for _, e := range arr {

		sMaxE = maxE
		maxE = max(maxE, e)

		sMinE = minE
		minE = min(minE, e)
	}
	return sMaxE, sMinE
}

func main() {
	arr := [10]int{4, 2, 9, 9, 7, 2, 1, 8, 8, 3}
	sMin, sMax := SecondMinMax(arr)
	fmt.Printf("Second smallest: %d, Second largest: %d\n", sMin, sMax)
}
