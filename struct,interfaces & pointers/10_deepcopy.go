package practice

// Profile has a slice field, which makes plain struct assignment (a
// "shallow copy") dangerous — the copy would share the same underlying
// array as the original. DeepCopy should return a Profile whose Tags slice
// is completely independent: mutating the copy's Tags must NOT affect the
// original's Tags, and vice versa.
type Profile struct {
	Name string
	Tags []string
}

func DeepCopy(p Profile) Profile {
	// TODO: implement
	return Profile{}
}
