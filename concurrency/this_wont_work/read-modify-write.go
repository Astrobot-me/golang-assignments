package thiswontwork

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

	if len(slice1) > 0 {
		w.Add(1)

		go func() {
			sum += SumAll(slice1, &w)
		}()
	}

	if len(slice2) > 0 {
		w.Add(1)

		go func() {
			sum += SumAll(slice2, &w)
		}()
	}

	w.Wait()
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

// The problem is here:

// go func() {
//     sum += SumAll(slice1, &w)
// }()

// go func() {
//     sum += SumAll(slice2, &w)
// }()

// Problem

// Both goroutines modify the same sum variable concurrently.

// sum += x is effectively a read-modify-write operation:

// temp := sum
// temp = temp + x
// sum = temp

// Imagine both goroutines see sum == 0:

// Goroutine 1                 Goroutine 2

// reads sum = 0               reads sum = 0
// adds 6                      adds 15
// writes sum = 6              writes sum = 15

// The final result could incorrectly be 15 instead of 21. This is a data race.

// Solution

// 1. both have a differt variable to write upon
// var sum1, sum2 int

// 2. Put into channels and get result back from channels
// Channels — cleanest for returning results from goroutines

// Instead of both goroutines modifying a shared sum, each goroutine sends its result through a channel.
