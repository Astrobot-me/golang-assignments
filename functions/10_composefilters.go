package practice

// ComposeFilters returns a new slice containing only the elements of nums
// that satisfy ALL of the given predicates (logical AND). If no predicates
// are given, return a copy of nums. If no elements satisfy every predicate,
// return nil.
func ComposeFilters(nums []int, predicates ...func(int) bool) []int {
	// TODO: implement
	if len(predicates) == 0 {
		return nums
	}
	ans := make([]int, 0)
	for _, num := range nums {
		sat := true
		for _, fn := range predicates {
			sat = sat && fn(num)
		}

		if sat {
			ans = append(ans, num)
		}
	}
	if len(ans) == 0 {
		return nil
	}
	return ans
}
