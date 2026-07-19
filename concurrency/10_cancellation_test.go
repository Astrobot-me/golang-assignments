package practice

import (
	"testing"
	"time"
)

func TestProduce(t *testing.T) {
	t.Run("emits successive integers from zero", func(t *testing.T) {
		done := make(chan struct{})
		defer close(done)

		ch := Produce(done)
		for want := 0; want < 5; want++ {
			got, ok := recvTimeout(t, ch)
			if !ok {
				t.Fatalf("channel closed early, wanted value %d", want)
			}
			if got != want {
				t.Fatalf("received %d, want %d", got, want)
			}
		}
	})

	t.Run("closes output after done is closed", func(t *testing.T) {
		done := make(chan struct{})
		ch := Produce(done)

		// Consume a few values, then cancel.
		for i := 0; i < 3; i++ {
			if _, ok := recvTimeout(t, ch); !ok {
				t.Fatal("channel closed before done was closed")
			}
		}
		close(done)

		// One value may already be in flight; after that the channel must
		// close. Drain with a deadline so a leaky implementation fails
		// instead of hanging.
		deadline := time.After(1 * time.Second)
		received := 0
		for {
			select {
			case _, ok := <-ch:
				if !ok {
					return // closed — correct
				}
				received++
				if received > 2 {
					t.Fatal("channel kept producing after done was closed")
				}
			case <-deadline:
				t.Fatal("output channel never closed after cancellation (goroutine leak)")
			}
		}
	})
}
