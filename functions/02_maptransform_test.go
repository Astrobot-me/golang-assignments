package practice

import (
	"reflect"
	"testing"
)

func TestMapTransform(t *testing.T) {
	cases := []struct {
		name      string
		nums      []int
		transform func(int) int
		want      []int
	}{
		{
			name:      "square",
			nums:      []int{1, 2, 3, 4},
			transform: func(n int) int { return n * n },
			want:      []int{1, 4, 9, 16},
		},
		{
			name:      "add ten",
			nums:      []int{-5, 0, 5, 10},
			transform: func(n int) int { return n + 10 },
			want:      []int{5, 10, 15, 20},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := MapTransform(tc.nums, tc.transform)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("MapTransform(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}
