package practice

import "fmt"

// GroupByLength groups words by their length and returns a map from length
// to the slice of words with that length. Within each group, preserve the
// original relative order of the words.
func GroupByLength(words []string) map[int][]string {
	// TODO: implement

	ans := make(map[int][]string)

	for _, e := range words {

		key := len(e)
		// _, isPresent := ans[key]

		ans[key] = append(ans[key], e)

	}
	return ans
}

func main() {
	words := []string{"go", "is", "fun", "cat", "dog", "elephant"}

	fmt.Println(GroupByLength(words))
}
