package structsmethodsinterfaces

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Length float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}

func (r Rectangle) Perimeter() float64 {
	return r.Width*2 + r.Length*2
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return math.Pi * c.Radius * 2
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
	return 0.5
}
