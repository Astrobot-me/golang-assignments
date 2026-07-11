package practice

import (
	"reflect"
	"testing"
)

func TestApplyToEach(t *testing.T) {
	t.Run("collects squares in order", func(t *testing.T) {
		var got []int
		ApplyToEach([]int{1, 2, 3, 4}, func(n int) {
			got = append(got, n*n)
		})
		want := []int{1, 4, 9, 16}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("counts calls", func(t *testing.T) {
		count := 0
		ApplyToEach([]int{10, 20, 30, 40, 50}, func(n int) {
			count++
		})
		if count != 5 {
			t.Errorf("action was called %d times, want 5", count)
		}
	})

	t.Run("empty slice never calls action", func(t *testing.T) {
		called := false
		ApplyToEach([]int{}, func(n int) {
			called = true
		})
		if called {
			t.Errorf("action should not be called for an empty slice")
		}
	})
}
