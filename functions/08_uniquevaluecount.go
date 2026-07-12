package practice

import (
	"slices"
)

// UniqueValueCount returns the number of distinct values that appear in m
// (keys are ignored; only the values matter).
func UniqueValueCount(m map[string]int) int {
	// TODO: implement
	//
	ans := make([]int, 0)
	for _, v := range m {
		if !slices.Contains(ans, v) {
			ans = append(ans, v)
		}
	}
	return len(ans)
}
