package practice

import (
	"math"
	"testing"
)

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-6
}

func TestCircleAreaAndPerimeter(t *testing.T) {
	c := Circle{Radius: 2}
	if got := c.Area(); !almostEqual(got, math.Pi*4) {
		t.Errorf("Area() = %v, want %v", got, math.Pi*4)
	}
	if got := c.Perimeter(); !almostEqual(got, 2*math.Pi*2) {
		t.Errorf("Perimeter() = %v, want %v", got, 2*math.Pi*2)
	}
}

func TestRectangleAreaAndPerimeter(t *testing.T) {
	r := Rectangle{Width: 4, Height: 5}
	if got := r.Area(); !almostEqual(got, 20) {
		t.Errorf("Area() = %v, want 20", got)
	}
	if got := r.Perimeter(); !almostEqual(got, 18) {
		t.Errorf("Perimeter() = %v, want 18", got)
	}
}

func TestTotalArea(t *testing.T) {
	shapes := []Shape{
		Circle{Radius: 1},              // π
		Rectangle{Width: 2, Height: 3}, // 6
	}
	want := math.Pi + 6
	if got := TotalArea(shapes); !almostEqual(got, want) {
		t.Errorf("TotalArea() = %v, want %v", got, want)
	}
}
