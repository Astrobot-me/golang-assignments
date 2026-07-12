package main

import "fmt"

// Concepts: constants (iota), switch, conditionals
//
// Task:
// Using iota, define constants for the four seasons: Winter, Spring, Summer, Autumn.
// Implement GetSeason which takes a month (1-12) and returns the season name as a string.
//   Dec, Jan, Feb  -> "Winter"
//   Mar, Apr, May  -> "Spring"
//   Jun, Jul, Aug  -> "Summer"
//   Sep, Oct, Nov  -> "Autumn"
// If month is outside 1-12, return "Invalid month".
// Use a switch statement with multiple values per case (e.g. case 12, 1, 2:).

const (
	// TODO: define season constants using iota
	Winter = iota
	Spring
	Summer
	Autumn
)

func GetSeason(month int) string {
	// TODO: implement
	switch month {
	case 1, 2, 12:
		return "Winter"

	case 3, 4, 5:
		return "Spring"

	case 6, 7, 8:
		return "Summer"
	case 9, 10, 11:
		return "Autumn"
	default:
		return "Invalid Month"
	}
	return ""
}

func main() {
	tests := []int{1, 4, 7, 10, 13, 0}
	for _, m := range tests {
		fmt.Printf("Month %d -> %s\n", m, GetSeason(m))
	}
}
