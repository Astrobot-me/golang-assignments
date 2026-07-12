package main

import "fmt"

// Concepts: constants, loops, switch, conditionals
//
// Task:
// Define grade boundary constants (e.g. GradeABoundary = 90, GradeBBoundary = 75, ...).
// Implement GradeReport which takes a slice of int scores, computes the
// average using a loop, and determines a letter grade using a switch
// statement built on those boundary constants:
//   avg >= 90         -> "A"
//   avg >= 75 && < 90 -> "B"
//   avg >= 60 && < 75 -> "C"
//   avg >= 40 && < 60 -> "D"
//   avg < 40          -> "F"
// Return both the average and the grade. If the slice is empty, return (0, "N/A").

const (
	GradeABoundary = 90
	GradeBBoundary = 75
	GradeCBoundary = 60
	GradeDBoundary = 40
)

func GradeReport(scores []int) (avg float64, grade string) {
	// TODO: implement
	if n := len(scores); n == 0 {
		return 0, "N/A"
	}
	sum := 0
	for _, e := range scores {
		sum += e
	}
	avg = float64(sum / len(scores))
	grade = ""

	if avg >= GradeABoundary {
		grade = "A"
	} else if avg >= GradeBBoundary && avg < GradeABoundary {
		grade = "B"
	} else if avg >= GradeCBoundary && avg < GradeBBoundary {
		grade = "C"

	} else {
		grade = "D"
	}

	return avg, grade
}

func main() {
	scores := []int{88, 92, 79, 65, 95}
	avg, grade := GradeReport(scores)
	fmt.Printf("Average: %.2f, Grade: %s\n", avg, grade)

	empty := []int{}
	avg2, grade2 := GradeReport(empty)
	fmt.Printf("Average: %.2f, Grade: %s\n", avg2, grade2)
}
