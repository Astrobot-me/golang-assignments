package practice

import (
	"reflect"
	"testing"
)

func TestPingPong(t *testing.T) {
	cases := []struct {
		name   string
		rounds int
		want   []string
	}{
		{name: "one round", rounds: 1, want: []string{"ping", "pong"}},
		{name: "three rounds", rounds: 3, want: []string{"ping", "pong", "ping", "pong", "ping", "pong"}},
		{name: "zero rounds", rounds: 0, want: nil},
		{name: "negative rounds", rounds: -2, want: nil},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := PingPong(tc.rounds)
			if !reflect.DeepEqual(got, tc.want) && !(len(got) == 0 && len(tc.want) == 0) {
				t.Errorf("PingPong(%d) = %v, want %v", tc.rounds, got, tc.want)
			}
		})
	}

	// Many rounds makes broken turn-taking (interleaving bugs) show up.
	t.Run("fifty rounds stays strictly alternating", func(t *testing.T) {
		got := PingPong(50)
		if len(got) != 100 {
			t.Fatalf("PingPong(50) returned %d entries, want 100", len(got))
		}
		for i, s := range got {
			want := "ping"
			if i%2 == 1 {
				want = "pong"
			}
			if s != want {
				t.Fatalf("entry %d = %q, want %q (sequence broke alternation)", i, s, want)
			}
		}
	})
}
