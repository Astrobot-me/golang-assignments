package practice

// Concepts: channel ownership, returning receive-only channels, close, range
//
// Task:
// Implement Generate which returns a receive-only channel that yields the
// integers 1..n (inclusive) and is then closed. Start the producing goroutine
// inside Generate and return immediately — do NOT produce all values before
// returning.
//
// If n <= 0, return a channel that is closed without sending anything.
//
// Rule of thumb being practised here: the goroutine that SENDS on a channel
// is the one that closes it. The caller should be able to `for v := range ch`.

func Generate(n int) <-chan int {
	// TODO: implement
	return nil
}
