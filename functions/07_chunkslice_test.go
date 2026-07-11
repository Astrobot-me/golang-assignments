package practice

import (
	"reflect"
	"testing"
)

func TestChunkSlice(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		size int
		want [][]int
	}{
		{
			name: "divides evenly",
			nums: []int{1, 2, 3, 4, 5, 6},
			size: 2,
			want: [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name: "leftover chunk",
			nums: []int{1, 2, 3, 4, 5},
			size: 2,
			want: [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			name: "size larger than slice",
			nums: []int{1, 2, 3},
			size: 10,
			want: [][]int{{1, 2, 3}},
		},
		{
			name: "invalid size",
			nums: []int{1, 2, 3},
			size: 0,
			want: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ChunkSlice(tc.nums, tc.size)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ChunkSlice(%v, %d) = %v, want %v", tc.nums, tc.size, got, tc.want)
			}
		})
	}
}
