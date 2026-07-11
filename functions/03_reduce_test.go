package practice

import "testing"

func TestReduceSlice(t *testing.T) {
	sum := func(acc, cur int) int { return acc + cur }
	product := func(acc, cur int) int { return acc * cur }

	cases := []struct {
		name    string
		nums    []int
		initial int
		reducer func(acc, cur int) int
		want    int
	}{
		{name: "sum", nums: []int{1, 2, 3, 4}, initial: 0, reducer: sum, want: 10},
		{name: "product", nums: []int{1, 2, 3, 4}, initial: 1, reducer: product, want: 24},
		{name: "empty slice", nums: []int{}, initial: 42, reducer: sum, want: 42},
		{name: "sum with non-zero initial", nums: []int{5, 5}, initial: 100, reducer: sum, want: 110},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ReduceSlice(tc.nums, tc.initial, tc.reducer)
			if got != tc.want {
				t.Errorf("ReduceSlice(%v, %d) = %d, want %d", tc.nums, tc.initial, got, tc.want)
			}
		})
	}
}
