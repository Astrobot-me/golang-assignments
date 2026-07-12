package practice

// MergeMaps combines a and b into a new map. If a key exists in both, the
// values are summed. Neither a nor b should be modified.
func MergeMaps(a, b map[string]int) map[string]int {
	// TODO: implement
	ans := make(map[string]int)
	for k, v := range a {
		ans[k] = v
	}

	for k, v := range b {
		ans[k] = ans[k] + v
	}
	return ans
}

// func main() {
// 	a := map[string]int{"apple": 3, "banana": 5}
// 	b := map[string]int{"cherry": 2, "apple": 5}

// 	fmt.Println(MergeMaps(a, b))
// }
