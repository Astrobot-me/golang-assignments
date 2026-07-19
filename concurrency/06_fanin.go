package practice

// Concepts: fan-in, sync.WaitGroup combined with channels
//
// Task:
// Implement Merge which fans in any number of input channels into a single
// output channel: every value received from any input is forwarded to the
// output. The output channel must be closed only after ALL inputs are
// exhausted (closed and drained).
//
// Classic shape:
//   - one forwarding goroutine per input channel
//   - a sync.WaitGroup counting the forwarders
//   - one extra goroutine that waits on the WaitGroup, then closes the output
//
// Merge() with no arguments returns a channel that is closed immediately.
// Output order across different inputs is unspecified.

func Merge(chs ...<-chan int) <-chan int {
	// TODO: implement
	return nil
}
