package practice

import "testing"

func TestDescribe(t *testing.T) {
	cases := []struct {
		name  string
		input interface{}
		want  string
	}{
		{"int", 5, "int: 5"},
		{"string", "hi", "string: hi"},
		{"bool", true, "bool: true"},
		{"int slice", []int{1, 2, 3}, "int slice of length 3"},
		{"unknown", 3.14, "unknown type"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Describe(tc.input); got != tc.want {
				t.Errorf("Describe(%v) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
