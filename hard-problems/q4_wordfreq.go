package main

import "fmt"

// Concepts: maps, loops, switch, conditionals
//
// Task:
// Implement WordFrequencyCategory which takes a slice of words, counts how
// many times each word appears, and returns a map[string]string categorizing
// each UNIQUE word by its frequency:
//   count == 1   -> "rare"
//   count 2-3    -> "common"
//   count >= 4   -> "frequent"
// Use a switch statement (switch true, with range conditions in each case)
// for the categorization step — not a chain of if/else.

func WordFrequencyCategory(words []string) map[string]string {
	// TODO: implement
	return nil
}

func main() {
	words := []string{"go", "is", "fun", "go", "go", "is", "fast", "go"}
	result := WordFrequencyCategory(words)
	for word, category := range result {
		fmt.Printf("%-8s -> %s\n", word, category)
	}
}
