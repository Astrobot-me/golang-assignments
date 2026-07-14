package practice

// Describe takes any value and returns a short description of its concrete
// type using a type switch. Handle at least: int, string, bool, []int, and
// a default case for anything else.
//   Describe(5)          -> "int: 5"
//   Describe("hi")       -> "string: hi"
//   Describe(true)       -> "bool: true"
//   Describe([]int{1,2}) -> "int slice of length 2"
//   Describe(3.14)       -> "unknown type"
func Describe(i interface{}) string {
	// TODO: implement
	return ""
}
