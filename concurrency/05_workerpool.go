package practice

// Concepts: worker pools, buffered channels, closing to signal "no more work"
//
// Task:
// Implement WorkerPool which processes every job using exactly `workers`
// worker goroutines:
//   1. Send all jobs on a jobs channel, then close it — a closed channel is
//      how workers learn there is no more work.
//   2. Each worker ranges over the jobs channel, applies fn to each job, and
//      sends the output on a results channel.
//   3. Collect every result and return them. Order does NOT matter.
//
// Make sure WorkerPool collects exactly len(jobs) results before returning,
// and that no goroutine is left blocked forever after it returns.
// Assume workers >= 1.

func WorkerPool(jobs []int, workers int, fn func(int) int) []int {
	// TODO: implement
	return nil
}
