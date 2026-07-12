package main

import "fmt"

// Concepts: conditionals
//
// Task:
// Implement IsLeapYear which returns true if the given year is a leap year.
// Rules: divisible by 4, EXCEPT century years (divisible by 100) unless also
// divisible by 400.

func IsLeapYear(year int) bool {
	// TODO: implement

	if year%4 == 0 {
		if year%100 != 0 {
			return true
		} else if year%400 == 0 {
			return true
		}

		return false
	}

	return false
}

func main() {
	years := []int{2000, 1900, 2024, 2023}
	for _, y := range years {
		fmt.Printf("%d -> %t\n", y, IsLeapYear(y))
	}
}
