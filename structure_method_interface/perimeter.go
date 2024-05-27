package perimeter

import "math"

// INTERFACE
type Shape interface {
	Area() float64
	Perimeter() float64
}

// -------------------------------------------------
// SHAPE
type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	height float64
	width  float64
}

// -------------------------------------------------
// PERIMETER
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (t Triangle) Perimeter() float64 {
	return 3 * t.width
}

// -------------------------------------------------
// AREA
func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func (t Triangle) Area() float64 {
	return t.height * t.width * 0.5
}
