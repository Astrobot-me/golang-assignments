package practice

// Shared test helpers. Receiving from a nil or never-closed channel blocks
// forever, so every test that reads from an exercise's channel goes through
// these guards — an unimplemented or deadlocked exercise fails fast instead
// of hanging the whole test run.

import (
	"testing"
	"time"
)

// recvTimeout receives one value from ch, failing the test if nothing
// arrives within a second. The second return value is false if the channel
// was closed.
func recvTimeout[T any](t *testing.T, ch <-chan T) (T, bool) {
	t.Helper()
	var zero T
	if ch == nil {
		t.Fatal("channel is nil — is the function implemented?")
		return zero, false
	}
	select {
	case v, ok := <-ch:
		return v, ok
	case <-time.After(1 * time.Second):
		t.Fatal("timed out waiting on channel — possible deadlock or missing send/close")
		return zero, false
	}
}

// collectTimeout receives from ch until it is closed and returns everything
// received, failing the test if the channel is not closed within 2 seconds.
func collectTimeout[T any](t *testing.T, ch <-chan T) []T {
	t.Helper()
	if ch == nil {
		t.Fatal("channel is nil — is the function implemented?")
		return nil
	}
	var out []T
	deadline := time.After(2 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return out
			}
			out = append(out, v)
		case <-deadline:
			t.Fatalf("timed out after receiving %d value(s) — is the channel ever closed?", len(out))
			return nil
		}
	}
}

// sendAndClose returns a channel pre-loaded with vals and already closed —
// handy as a ready-made input for fan-in / pipeline exercises.
func sendAndClose(vals ...int) <-chan int {
	ch := make(chan int, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
	return ch
}
