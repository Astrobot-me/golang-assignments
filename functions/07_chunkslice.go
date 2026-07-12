package practice

// ChunkSlice splits nums into consecutive chunks of the given size. The
// last chunk may be smaller than size if len(nums) doesn't divide evenly.
// If size <= 0, return nil.
func ChunkSlice(nums []int, size int) [][]int {
	// TODO: implement

	ans := make([][]int, 0)
	if len(nums) <= 0 {
		return nil
	}

	if len(nums) < size {
		ans = append(ans, ans...)
		return ans
	}

	for i := 0; i < len(nums); i += size {
		j := i + size
		j = min(j, len(nums))
		ans = append(ans, nums[i:j])
	}

	return ans

}
