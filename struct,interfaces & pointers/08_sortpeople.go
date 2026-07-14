package practice

// Person has a name and an age.
type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface for []Person, sorting by Age ascending.
// (In main you'd call sort.Sort(ByAge(people)) or sort.Sort(people) if
// people is already typed as ByAge.)
type ByAge []Person

func (a ByAge) Len() int {
	// TODO: implement
	return 0
}

func (a ByAge) Less(i, j int) bool {
	// TODO: implement
	return false
}

func (a ByAge) Swap(i, j int) {
	// TODO: implement
}
