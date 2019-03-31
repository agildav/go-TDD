package structsmethodsinterfaces

import "math"

type Shape interface {
	area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) perimeter() float64 {
	if r.width < 0 || r.height < 0 {
		return 0.0
	}
	return 2 * (r.width + r.height)
}

func (r Rectangle) area() float64 {
	if r.width < 0 || r.height < 0 {
		return 0.0
	}
	return r.width * r.height
}

func (c Circle) area() float64 {
	if c.radius < 0 {
		return 0.0
	}
	return math.Pi * c.radius * c.radius
}
