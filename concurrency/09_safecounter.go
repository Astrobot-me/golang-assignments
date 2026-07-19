package practice

// Concepts: data races, sync.Mutex, methods on pointer receivers
//
// Task:
// Implement Counter so it is safe to use from many goroutines at once:
//   - Inc() adds 1 to the count.
//   - Value() returns the current count.
// Guard the count with a sync.Mutex (lock in BOTH methods — reads race too).
//
// First try it WITHOUT the mutex and run:
//   go test -race -run TestCounter ./...
// Watch the race detector catch it, then add the mutex.

type Counter struct {
	// TODO: add fields — the count and a sync.Mutex guarding it
}

func (c *Counter) Inc() {
	// TODO: implement
}

func (c *Counter) Value() int {
	// TODO: implement
	return 0
}
