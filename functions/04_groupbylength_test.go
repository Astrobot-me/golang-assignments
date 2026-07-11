package practice

import (
	"reflect"
	"testing"
)

func TestGroupByLength(t *testing.T) {
	cases := []struct {
		name  string
		words []string
		want  map[int][]string
	}{
		{
			name:  "mixed lengths",
			words: []string{"go", "is", "fun", "cat", "dog", "elephant"},
			want: map[int][]string{
				2: {"go", "is"},
				3: {"fun", "cat", "dog"},
				8: {"elephant"},
			},
		},
		{
			name:  "all same length",
			words: []string{"ant", "bee", "cow"},
			want: map[int][]string{
				3: {"ant", "bee", "cow"},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := GroupByLength(tc.words)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GroupByLength(%v) = %v, want %v", tc.words, got, tc.want)
			}
		})
	}
}
