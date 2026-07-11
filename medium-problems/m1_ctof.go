package main

import "fmt"

// Concepts: variables, constants
//
// Task:
// Declare a constant for the conversion factor (9.0/5.0) and implement
// CelsiusToFahrenheit which converts a Celsius temperature to Fahrenheit.
// Formula: F = C * 9/5 + 32

const (
	// TODO: define conversion constant(s)
)

func CelsiusToFahrenheit(celsius float64) float64 {
	// TODO: implement
	return 0
}

func main() {
	temps := []float64{0, 37, 100, -40}
	for _, c := range temps {
		fmt.Printf("%.1f°C = %.1f°F\n", c, CelsiusToFahrenheit(c))
	}
}
