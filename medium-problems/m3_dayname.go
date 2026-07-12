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
	switch day {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return "Invalid day"
	}
}

func main() {
	for d := 0; d <= 8; d++ {
		fmt.Printf("%d -> %s\n", d, DayName(d))
	}
}
