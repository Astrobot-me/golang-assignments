package practice

import "fmt"

// Money stores an amount in CENTS to avoid floating point rounding issues.
type Money struct {
	Cents int
}

// String implements the fmt.Stringer interface. Format the amount as
// dollars with exactly two decimal places, e.g.:
//   Money{Cents: 150}  -> "$1.50"
//   Money{Cents: 5}    -> "$0.05"
//   Money{Cents: -250} -> "-$2.50"
func (m Money) String() string {
	// TODO: implement
	return ""
}

// Compile-time check that Money satisfies fmt.Stringer.
var _ fmt.Stringer = Money{}
