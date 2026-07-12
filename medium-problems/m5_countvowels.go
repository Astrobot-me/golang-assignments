package main

import "fmt"

// Concepts: loops, range, conditionals
//
// Task:
// Implement CountVowels which returns the count of vowels (a,e,i,o,u —
// case-insensitive) in a string. Use range to iterate over the string.

func CountVowels(s string) int {
	// TODO: implement
	count := 0
	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			count++
		}
	}
	return count
}

func main() {
	words := []string{"Hello World", "GoLang", "xyz"}
	for _, w := range words {
		fmt.Printf("%-12s -> %d vowels\n", w, CountVowels(w))
	}
}
