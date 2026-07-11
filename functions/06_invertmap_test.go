package practice

import (
	"reflect"
	"testing"
)

func TestInvertMap(t *testing.T) {
	cases := []struct {
		name string
		m    map[string]int
		want map[int]string
	}{
		{
			name: "simple",
			m:    map[string]int{"one": 1, "two": 2, "three": 3},
			want: map[int]string{1: "one", 2: "two", 3: "three"},
		},
		{
			name: "single entry",
			m:    map[string]int{"only": 99},
			want: map[int]string{99: "only"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := InvertMap(tc.m)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("InvertMap(%v) = %v, want %v", tc.m, got, tc.want)
			}
		})
	}
}
