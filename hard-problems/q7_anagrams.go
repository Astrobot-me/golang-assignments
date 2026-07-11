package main

import "fmt"

// Concepts: maps, slices, arrays, range, functions
//
// Task:
// Implement GroupAnagrams which takes a slice of lowercase words and groups
// words that are anagrams of each other into a map[string][]string, where
// the key is a canonical signature for the anagram group and the value is
// the slice of words in that group.
//
// You may NOT use the "sort" package. Instead build the canonical signature
// yourself: use a [26]int array to count letter occurrences (a-z) in a word,
// then loop through that array and build a string key manually
// (e.g. "a2b0c1...z0").
//
// Example input: "eat", "tea", "tan", "ate", "nat", "bat"
//   -> groups: ["eat","tea","ate"], ["tan","nat"], ["bat"]

func GroupAnagrams(words []string) map[string][]string {
	// TODO: implement
	return nil
}

func main() {
	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groups := GroupAnagrams(words)
	for key, group := range groups {
		fmt.Printf("%-30s -> %v\n", key, group)
	}
}
