package practice

// MapTransform applies transform to every element of nums and returns the
// results as a new slice, in the same order. Do not modify the input slice.
func MapTransform(nums []int, transform func(int) int) []int {
	// TODO: implement
	ans := make([]int, 0)

	for _, i := range nums {

		ans = append(ans, transform(i))
	}
	return ans
}
