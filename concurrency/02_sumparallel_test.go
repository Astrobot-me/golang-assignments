package practice

import "testing"

func TestSumParallel(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want int
	}{
		{name: "basic", nums: []int{1, 2, 3, 4}, want: 10},
		{name: "odd length", nums: []int{5, 10, 15}, want: 30},
		{name: "single", nums: []int{42}, want: 42},
		{name: "negatives", nums: []int{-5, 5, -10, 10}, want: 0},
		{name: "empty", nums: nil, want: 0},
		{name: "larger", nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, want: 55},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SumParallel(tc.nums); got != tc.want {
				t.Errorf("SumParallel(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}
