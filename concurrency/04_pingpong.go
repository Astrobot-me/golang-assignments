package practice

import "sync"

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

	if rounds <= 0 {
		return []string{}
	}

	// TODO: implement

	var w sync.WaitGroup

	ping := make(chan int)
	pong := make(chan int)

	result := make([]string, 0, 2*rounds)

	w.Add(2)

	go func() {

		defer w.Done()

		for range rounds {
			<-ping
			result = append(result, "ping")
			pong <- 1
		}

	}()

	go func() {

		defer w.Done()

		for i := range rounds {
			<-pong
			result = append(result, "pong")

			if i < rounds-1 {

				ping <- 1
			}
		}

	}()

	ping <- 1

	w.Wait()
	return result
}
