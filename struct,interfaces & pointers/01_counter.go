package practice

// Counter has a single field Count. Implement Increment as a POINTER
// receiver method so it mutates the original Counter (calling it multiple
// times should accumulate). Implement Value as a value receiver method
// that just returns the current count.
//
// Interview angle: explain why Increment MUST use a pointer receiver to
// work correctly, and what would happen if it used a value receiver instead.
type Counter struct {
	Count int
}

func (c *Counter) Increment() {
	// TODO: implement
}

func (c Counter) Value() int {
	// TODO: implement
	return 0
}
