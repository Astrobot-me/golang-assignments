package practice

import (
	"reflect"
	"testing"
)

func TestSquareAll(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "basic", nums: []int{1, 2, 3, 4}, want: []int{1, 4, 9, 16}},
		{name: "single", nums: []int{7}, want: []int{49}},
		{name: "negatives", nums: []int{-3, 0, 5}, want: []int{9, 0, 25}},
		{name: "bigger slice keeps order", nums: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, want: []int{100, 81, 64, 49, 36, 25, 16, 9, 4, 1}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := SquareAll(tc.nums)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("SquareAll(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}

	t.Run("empty", func(t *testing.T) {
		if got := SquareAll(nil); len(got) != 0 {
			t.Errorf("SquareAll(nil) = %v, want empty", got)
		}
	})
}
