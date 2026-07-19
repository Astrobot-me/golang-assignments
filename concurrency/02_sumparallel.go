package practice

// Concepts: unbuffered channels, send/receive, goroutines
//
// Task:
// Implement SumParallel which splits nums into two halves and sums each half
// in its own goroutine. Each goroutine sends its partial sum on a channel;
// the calling goroutine receives exactly two values and returns their total.
//
// An empty or nil slice returns 0. An odd-length slice can split anywhere
// (e.g. len/2).
//
// Note how an unbuffered channel both transfers the data AND synchronises:
// the receive cannot complete before the goroutine has finished its half.

func SumParallel(nums []int) int {
	// TODO: implement
	return 0
}
