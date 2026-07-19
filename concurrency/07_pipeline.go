package practice

// Concepts: pipelines — chaining channel stages
//
// Task:
// Build a three-stage pipeline, then chain it:
//
//   1. emit(nums)  — returns a channel that yields the values of nums in
//                    order, then closes.
//   2. square(in)  — returns a channel that yields x*x for every x received
//                    from in, and closes when in closes.
//   3. total(in)   — receives until in is closed and returns the sum.
//
// Then implement SumOfSquares by chaining: emit → square → total.
// Each stage owns (and closes) the channel it returns.
//
// e.g. SumOfSquares([]int{1, 2, 3}) == 1 + 4 + 9 == 14

func emit(nums []int) <-chan int {
	// TODO: implement
	return nil
}

func square(in <-chan int) <-chan int {
	// TODO: implement
	return nil
}

func total(in <-chan int) int {
	// TODO: implement
	return 0
}

func SumOfSquares(nums []int) int {
	// TODO: chain emit → square → total
	return 0
}
