package practice

// Concepts: done channels, cancellation, select on send, goroutine leaks
//
// Task:
// Implement Produce which returns a channel emitting the successive integers
// 0, 1, 2, 3, ... forever — until the caller closes `done`.
//
// The producing goroutine must stop and close its output channel as soon as
// done is closed. Use a select that covers BOTH cases at once:
//
//   select {
//   case out <- next:   // consumer took a value
//   case <-done:        // caller cancelled — clean up and return
//   }
//
// Without the <-done case the goroutine would block forever on the send once
// the consumer stops reading — a goroutine leak. This done-channel pattern is
// how context.Context cancellation works under the hood (ctx.Done()).

func Produce(done <-chan struct{}) <-chan int {
	// TODO: implement
	return nil
}
