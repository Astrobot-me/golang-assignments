package practice

import (
	"reflect"
	"testing"
)

func copyMap(m map[string]int) map[string]int {
	c := make(map[string]int, len(m))
	for k, v := range m {
		c[k] = v
	}
	return c
}

func TestMergeMaps(t *testing.T) {
	cases := []struct {
		name string
		a    map[string]int
		b    map[string]int
		want map[string]int
	}{
		{
			name: "no overlap",
			a:    map[string]int{"apple": 3, "banana": 5},
			b:    map[string]int{"cherry": 2},
			want: map[string]int{"apple": 3, "banana": 5, "cherry": 2},
		},
		{
			name: "overlapping keys sum",
			a:    map[string]int{"apple": 3, "banana": 5},
			b:    map[string]int{"apple": 10, "cherry": 2},
			want: map[string]int{"apple": 13, "banana": 5, "cherry": 2},
		},
		{
			name: "one side empty",
			a:    map[string]int{},
			b:    map[string]int{"apple": 1},
			want: map[string]int{"apple": 1},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			aBefore := copyMap(tc.a)
			bBefore := copyMap(tc.b)

			got := MergeMaps(tc.a, tc.b)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("MergeMaps(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
			if !reflect.DeepEqual(tc.a, aBefore) {
				t.Errorf("input map a was mutated: got %v, want %v", tc.a, aBefore)
			}
			if !reflect.DeepEqual(tc.b, bBefore) {
				t.Errorf("input map b was mutated: got %v, want %v", tc.b, bBefore)
			}
		})
	}
}
