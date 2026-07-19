package practice

// Concepts: goroutines, sync.WaitGroup
//
// Task:
// Implement SquareAll which squares every number concurrently — one goroutine
// per element. Each goroutine writes its result into the correct index of a
// pre-allocated result slice (writing to *different* indexes of a slice from
// different goroutines is safe). Use a sync.WaitGroup to wait for all
// goroutines to finish before returning.
//
// The result must be in the same order as the input.
//
// Watch out for the classic loop-variable capture bug: pass the index/value
// into the goroutine explicitly.
//
// Verify with the race detector: go test -race -run TestSquareAll ./...

func SquareAll(nums []int) []int {
	// TODO: implement
	return nil
}
