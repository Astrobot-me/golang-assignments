package practice

import (
	"reflect"
	"sort"
	"testing"
)

func TestMerge(t *testing.T) {
	t.Run("three inputs", func(t *testing.T) {
		merged := Merge(
			sendAndClose(1, 2, 3),
			sendAndClose(10, 20),
			sendAndClose(100),
		)
		got := collectTimeout(t, merged)
		sort.Ints(got)
		want := []int{1, 2, 3, 10, 20, 100}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Merge yielded (sorted) %v, want %v", got, want)
		}
	})

	t.Run("single input", func(t *testing.T) {
		got := collectTimeout(t, Merge(sendAndClose(7, 8, 9)))
		sort.Ints(got)
		if !reflect.DeepEqual(got, []int{7, 8, 9}) {
			t.Errorf("Merge yielded (sorted) %v, want [7 8 9]", got)
		}
	})

	t.Run("empty input channel among others", func(t *testing.T) {
		got := collectTimeout(t, Merge(sendAndClose(), sendAndClose(5)))
		if !reflect.DeepEqual(got, []int{5}) {
			t.Errorf("Merge yielded %v, want [5]", got)
		}
	})

	t.Run("no inputs closes immediately", func(t *testing.T) {
		got := collectTimeout(t, Merge())
		if len(got) != 0 {
			t.Errorf("Merge() yielded %v, want nothing", got)
		}
	})
}
