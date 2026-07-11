package practice

import "testing"

func TestUniqueValueCount(t *testing.T) {
	cases := []struct {
		name string
		m    map[string]int
		want int
	}{
		{
			name: "all distinct",
			m:    map[string]int{"a": 1, "b": 2, "c": 3},
			want: 3,
		},
		{
			name: "some duplicates",
			m:    map[string]int{"a": 1, "b": 1, "c": 2, "d": 2, "e": 3},
			want: 3,
		},
		{
			name: "empty map",
			m:    map[string]int{},
			want: 0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := UniqueValueCount(tc.m)
			if got != tc.want {
				t.Errorf("UniqueValueCount(%v) = %d, want %d", tc.m, got, tc.want)
			}
		})
	}
}
