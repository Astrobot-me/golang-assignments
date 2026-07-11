package practice

import (
	"reflect"
	"testing"
)

func TestComposeFilters(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	isPositive := func(n int) bool { return n > 0 }
	greaterThan10 := func(n int) bool { return n > 10 }

	cases := []struct {
		name       string
		nums       []int
		predicates []func(int) bool
		want       []int
	}{
		{
			name:       "even and positive",
			nums:       []int{-4, -3, -2, -1, 0, 1, 2, 3, 4},
			predicates: []func(int) bool{isEven, isPositive},
			want:       []int{2, 4},
		},
		{
			name:       "three predicates",
			nums:       []int{5, 12, 18, 21, 24, 30},
			predicates: []func(int) bool{isEven, isPositive, greaterThan10},
			want:       []int{12, 18, 24, 30},
		},
		{
			name:       "no predicates returns copy",
			nums:       []int{1, 2, 3},
			predicates: nil,
			want:       []int{1, 2, 3},
		},
		{
			name:       "nothing satisfies all",
			nums:       []int{1, 3, 5},
			predicates: []func(int) bool{isEven},
			want:       nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ComposeFilters(tc.nums, tc.predicates...)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ComposeFilters(%v, %d preds) = %v, want %v", tc.nums, len(tc.predicates), got, tc.want)
			}
		})
	}
}
