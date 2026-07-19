package practice

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	cases := []struct {
		name string
		n    int
		want []int
	}{
		{name: "one to five", n: 5, want: []int{1, 2, 3, 4, 5}},
		{name: "single", n: 1, want: []int{1}},
		{name: "zero yields nothing", n: 0, want: nil},
		{name: "negative yields nothing", n: -3, want: nil},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collectTimeout(t, Generate(tc.n))
			if !reflect.DeepEqual(got, tc.want) && !(len(got) == 0 && len(tc.want) == 0) {
				t.Errorf("Generate(%d) yielded %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}
