package practice

import (
	"testing"
	"time"
)

func TestFirstValue(t *testing.T) {
	t.Run("first channel is ready", func(t *testing.T) {
		a := make(chan string, 1)
		a <- "from a"
		b := make(chan string) // never delivers
		if got := FirstValue(a, b); got != "from a" {
			t.Errorf("FirstValue = %q, want %q", got, "from a")
		}
	})

	t.Run("second channel is ready", func(t *testing.T) {
		a := make(chan string) // never delivers
		b := make(chan string, 1)
		b <- "from b"
		if got := FirstValue(a, b); got != "from b" {
			t.Errorf("FirstValue = %q, want %q", got, "from b")
		}
	})

	t.Run("slow sender still wins when it is the only one", func(t *testing.T) {
		a := make(chan string)
		b := make(chan string)
		go func() {
			time.Sleep(20 * time.Millisecond)
			a <- "late but only"
		}()
		if got := FirstValue(a, b); got != "late but only" {
			t.Errorf("FirstValue = %q, want %q", got, "late but only")
		}
	})
}

func TestRecvTimeout(t *testing.T) {
	t.Run("value already available", func(t *testing.T) {
		ch := make(chan int, 1)
		ch <- 42
		v, ok := RecvTimeout(ch, 500*time.Millisecond)
		if !ok || v != 42 {
			t.Errorf("RecvTimeout = (%d, %v), want (42, true)", v, ok)
		}
	})

	t.Run("times out on silent channel", func(t *testing.T) {
		ch := make(chan int)
		v, ok := RecvTimeout(ch, 30*time.Millisecond)
		if ok || v != 0 {
			t.Errorf("RecvTimeout = (%d, %v), want (0, false)", v, ok)
		}
	})

	t.Run("value arrives before deadline", func(t *testing.T) {
		ch := make(chan int)
		go func() {
			time.Sleep(10 * time.Millisecond)
			ch <- 7
		}()
		v, ok := RecvTimeout(ch, 500*time.Millisecond)
		if !ok || v != 7 {
			t.Errorf("RecvTimeout = (%d, %v), want (7, true)", v, ok)
		}
	})
}
