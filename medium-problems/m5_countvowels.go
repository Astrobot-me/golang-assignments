package main

import "fmt"

// Concepts: loops, range, conditionals
//
// Task:
// Implement CountVowels which returns the count of vowels (a,e,i,o,u —
// case-insensitive) in a string. Use range to iterate over the string.

func CountVowels(s string) int {
	// TODO: implement
	return 0
}

func main() {
	words := []string{"Hello World", "GoLang", "xyz"}
	for _, w := range words {
		fmt.Printf("%-12s -> %d vowels\n", w, CountVowels(w))
	}
}
