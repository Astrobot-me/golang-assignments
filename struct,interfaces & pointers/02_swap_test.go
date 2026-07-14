package practice

import "testing"

func TestSwapInts(t *testing.T) {
	x, y := 5, 10
	SwapInts(&x, &y)
	if x != 10 || y != 5 {
		t.Errorf("got x=%d, y=%d, want x=10, y=5", x, y)
	}
}

func TestSwapIntsSameValue(t *testing.T) {
	x, y := 7, 7
	SwapInts(&x, &y)
	if x != 7 || y != 7 {
		t.Errorf("got x=%d, y=%d, want x=7, y=7", x, y)
	}
}
