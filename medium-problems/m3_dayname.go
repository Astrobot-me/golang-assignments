package main

import "fmt"

// Concepts: switch
//
// Task:
// Implement DayName which takes an int (1-7) and returns the day name
// ("Monday" ... "Sunday"). Return "Invalid day" for anything outside 1-7.
// Use a switch statement with a default case.

func DayName(day int) string {
	// TODO: implement
	return ""
}

func main() {
	for d := 0; d <= 8; d++ {
		fmt.Printf("%d -> %s\n", d, DayName(d))
	}
}
