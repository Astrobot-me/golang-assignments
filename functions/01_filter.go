package practice

// FilterSlice returns a new slice containing only the elements of nums for
// which predicate returns true, preserving the original order. Do not
// modify the input slice. If no elements match, return nil.
func FilterSlice(nums []int, predicate func(int) bool) []int {
	// TODO: implement
	result := make([]int, 0)
	for _, e := range nums {
		if predicate(e) {
			result = append(result, e)
		}
	}
	if len(result) == 0 {
		return nil

	}
	return result
}
