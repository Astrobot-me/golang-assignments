package practice

import (
	"reflect"
	"sort"
	"testing"
)

func TestWorkerPool(t *testing.T) {
	square := func(n int) int { return n * n }

	cases := []struct {
		name    string
		jobs    []int
		workers int
		fn      func(int) int
		want    []int // compared ignoring order
	}{
		{name: "squares with 4 workers", jobs: []int{1, 2, 3, 4, 5, 6, 7, 8}, workers: 4, fn: square, want: []int{1, 4, 9, 16, 25, 36, 49, 64}},
		{name: "single worker", jobs: []int{3, 1, 2}, workers: 1, fn: square, want: []int{9, 1, 4}},
		{name: "more workers than jobs", jobs: []int{2, 4}, workers: 10, fn: square, want: []int{4, 16}},
		{name: "double with 3 workers", jobs: []int{10, 20, 30, 40, 50}, workers: 3, fn: func(n int) int { return n * 2 }, want: []int{20, 40, 60, 80, 100}},
		{name: "no jobs", jobs: nil, workers: 2, fn: square, want: nil},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := WorkerPool(tc.jobs, tc.workers, tc.fn)
			if len(got) != len(tc.want) {
				t.Fatalf("WorkerPool returned %d results (%v), want %d", len(got), got, len(tc.want))
			}
			gotSorted := append([]int(nil), got...)
			wantSorted := append([]int(nil), tc.want...)
			sort.Ints(gotSorted)
			sort.Ints(wantSorted)
			if !reflect.DeepEqual(gotSorted, wantSorted) && len(tc.want) > 0 {
				t.Errorf("WorkerPool results (sorted) = %v, want %v", gotSorted, wantSorted)
			}
		})
	}
}
