package main

import "fmt"

// Concepts: 2D arrays, nested loops, conditionals
//
// Task:
// Implement DiagonalSum which takes a 5x5 array of ints and returns the sum
// of both diagonals (top-left to bottom-right, and top-right to bottom-left).
// Because the size is odd, the CENTER element sits on both diagonals — make
// sure it's counted only ONCE.

func DiagonalSum(matrix [5][5]int) int {
	// TODO: implement
	return 0
}

func main() {
	matrix := [5][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}
	fmt.Println("Diagonal sum:", DiagonalSum(matrix))
}
