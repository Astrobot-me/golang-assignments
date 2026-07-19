package practice

import "sync"

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

	var w sync.WaitGroup
	// TODO: implement
	slice1 := nums[:len(nums)/2]
	slice2 := nums[len(nums)/2:]

	var sum int

	ch := make(chan int, 2)

	if len(slice1) > 0 {
		w.Add(1)

		go func() {
			ch <- SumAll(slice1, &w)
		}()
	}

	if len(slice2) > 0 {
		w.Add(1)

		go func() {
			ch <- SumAll(slice2, &w)
		}()
	}

	w.Wait()
	close(ch)

	for i := range ch {
		sum += i
	}
	return sum
}

func SumAll(nums []int, w *sync.WaitGroup) int {
	defer w.Done()
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}
