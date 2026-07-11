package practice

import (
	"reflect"
	"testing"
)

func TestFilterSlice(t *testing.T) {
	cases := []struct {
		name      string
		nums      []int
		predicate func(int) bool
		want      []int
	}{
		{
			name:      "evens",
			nums:      []int{1, 2, 3, 4, 5, 6},
			predicate: func(n int) bool { return n%2 == 0 },
			want:      []int{2, 4, 6},
		},
		{
			name:      "greater than 10",
			nums:      []int{3, 15, 8, 22, 1},
			predicate: func(n int) bool { return n > 10 },
			want:      []int{15, 22},
		},
		{
			name:      "none match",
			nums:      []int{1, 2, 3},
			predicate: func(n int) bool { return n > 100 },
			want:      nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := FilterSlice(tc.nums, tc.predicate)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("FilterSlice(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}
