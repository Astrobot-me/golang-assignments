package practice

import (
	"reflect"
	"testing"
)

func TestEmit(t *testing.T) {
	got := collectTimeout(t, emit([]int{4, 5, 6}))
	if !reflect.DeepEqual(got, []int{4, 5, 6}) {
		t.Errorf("emit yielded %v, want [4 5 6]", got)
	}

	if got := collectTimeout(t, emit(nil)); len(got) != 0 {
		t.Errorf("emit(nil) yielded %v, want nothing", got)
	}
}

func TestSquare(t *testing.T) {
	got := collectTimeout(t, square(sendAndClose(2, 3, 4)))
	if !reflect.DeepEqual(got, []int{4, 9, 16}) {
		t.Errorf("square yielded %v, want [4 9 16]", got)
	}
}

func TestTotal(t *testing.T) {
	if got := total(sendAndClose(1, 2, 3, 4)); got != 10 {
		t.Errorf("total = %d, want 10", got)
	}
	if got := total(sendAndClose()); got != 0 {
		t.Errorf("total of empty channel = %d, want 0", got)
	}
}

func TestSumOfSquares(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want int
	}{
		{name: "basic", nums: []int{1, 2, 3}, want: 14},
		{name: "single", nums: []int{5}, want: 25},
		{name: "negatives square positive", nums: []int{-2, 2}, want: 8},
		{name: "empty", nums: nil, want: 0},
		{name: "larger", nums: []int{1, 2, 3, 4, 5}, want: 55},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SumOfSquares(tc.nums); got != tc.want {
				t.Errorf("SumOfSquares(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}
