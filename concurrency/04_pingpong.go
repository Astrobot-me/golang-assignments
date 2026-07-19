package practice

// Concepts: goroutine coordination, unbuffered channels as turn signals
//
// Task:
// Implement PingPong which returns a slice of 2*rounds strings strictly
// alternating: "ping", "pong", "ping", "pong", ...
//
// Rules:
//   - One goroutine appends every "ping"; a DIFFERENT goroutine appends
//     every "pong". They share the same result slice.
//   - Use unbuffered channel(s) to make them take strict turns, so only one
//     goroutine touches the slice at a time (that's what makes the sharing
//     safe — verify with `go test -race -run TestPingPong ./...`).
//   - PingPong must not return until both goroutines are done.
//
// If rounds <= 0, return an empty (or nil) slice.

func PingPong(rounds int) []string {
	// TODO: implement
	return nil
}
