package practice

// Shape is satisfied by anything that can report an area and a perimeter.
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

// Hint: you'll want to import "math" for math.Pi.

func (c Circle) Area() float64 {
	// TODO: implement
	return 0
}

func (c Circle) Perimeter() float64 {
	// TODO: implement
	return 0
}

func (r Rectangle) Area() float64 {
	// TODO: implement
	return 0
}

func (r Rectangle) Perimeter() float64 {
	// TODO: implement
	return 0
}

// TotalArea sums the area of every shape in the slice, using only the
// Shape interface — it should work for any type that satisfies Shape,
// not just Circle and Rectangle.
func TotalArea(shapes []Shape) float64 {
	// TODO: implement
	return 0
}
