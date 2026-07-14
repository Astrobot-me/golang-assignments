package practice

// Base holds fields/methods common to several record types.
type Base struct {
	ID int
}

// Describe returns a short description using ID, e.g. Base{ID: 7} -> "record #7".
// This method lives on Base — embedding should make it callable directly on
// User and Product too, without redefining it on either.
func (b Base) Describe() string {
	// TODO: implement
	return ""
}

// User embeds Base and adds a Name field.
type User struct {
	Base
	Name string
}

// Product embeds Base and adds a Price field.
type Product struct {
	Base
	Price float64
}
