package main

import "fmt"

// Concepts: maps, range
//
// Task:
// Implement CharFrequency which returns a map[rune]int counting how many
// times each character appears in a string (spaces excluded).

func CharFrequency(s string) map[rune]int {
	// TODO: implement
	return nil
}

func main() {
	freq := CharFrequency("go gopher")
	for ch, count := range freq {
		fmt.Printf("%q -> %d\n", ch, count)
	}
}
