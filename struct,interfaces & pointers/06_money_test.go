package practice

import "testing"

func TestMoneyString(t *testing.T) {
	cases := []struct {
		name  string
		cents int
		want  string
	}{
		{"whole dollars", 150, "$1.50"},
		{"single cent", 5, "$0.05"},
		{"zero", 0, "$0.00"},
		{"large amount", 123456, "$1234.56"},
		{"negative", -250, "-$2.50"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m := Money{Cents: tc.cents}
			if got := m.String(); got != tc.want {
				t.Errorf("String() = %q, want %q", got, tc.want)
			}
		})
	}
}
