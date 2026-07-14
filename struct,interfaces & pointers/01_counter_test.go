package practice

import "testing"

func TestCounterIncrement(t *testing.T) {
	c := Counter{}
	c.Increment()
	c.Increment()
	c.Increment()
	if got := c.Value(); got != 3 {
		t.Errorf("Value() = %d, want 3", got)
	}
}

func TestCounterIncrementViaPointer(t *testing.T) {
	c := &Counter{Count: 10}
	c.Increment()
	if c.Count != 11 {
		t.Errorf("Count = %d, want 11", c.Count)
	}
}
