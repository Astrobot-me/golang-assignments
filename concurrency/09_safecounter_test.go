package practice

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("sequential", func(t *testing.T) {
		var c Counter
		for i := 0; i < 5; i++ {
			c.Inc()
		}
		if got := c.Value(); got != 5 {
			t.Errorf("Value() = %d, want 5", got)
		}
	})

	t.Run("concurrent increments", func(t *testing.T) {
		var c Counter
		const goroutines = 100
		const perGoroutine = 100

		var wg sync.WaitGroup
		for i := 0; i < goroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < perGoroutine; j++ {
					c.Inc()
				}
			}()
		}
		wg.Wait()

		if got := c.Value(); got != goroutines*perGoroutine {
			t.Errorf("Value() = %d, want %d (lost updates — is the count guarded?)", got, goroutines*perGoroutine)
		}
	})
}
