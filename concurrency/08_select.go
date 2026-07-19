package practice

import "time"

// Concepts: select, time.After, timeouts
//
// Task:
// 1. FirstValue waits on both channels simultaneously and returns whichever
//    value arrives first. Use select — do NOT receive from them one after
//    the other.
//
// 2. RecvTimeout waits up to d for a value from ch.
//    - If a value arrives in time, return (value, true).
//    - If d elapses first, return (0, false).
//    Use select with time.After.

func FirstValue(a, b <-chan string) string {
	// TODO: implement
	return ""
}

func RecvTimeout(ch <-chan int, d time.Duration) (int, bool) {
	// TODO: implement
	return 0, false
}
