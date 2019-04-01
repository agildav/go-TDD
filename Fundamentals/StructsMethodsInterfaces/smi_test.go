package structsmethodsinterfaces

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t *testing.T, got, want float64) {
		t.Helper()

		// Assert the message
		if got != want {
			errMsg := fmt.Sprintf("got -> %f, want -> %f", got, want)
			// Print the error with colors
			t.Errorf(colorError(errMsg))
		}
	}

	t.Run("perimeter of a rectangle ", func(t *testing.T) {
		width := 20.0
		height := 10.0

		rectangle := Rectangle{width, height}
		got := rectangle.perimeter()
		want := 60.0

		checkPerimeter(t, got, want)
	})

	t.Run("zero values of a rectangle give zero perimeter", func(t *testing.T) {
		width := (-20.0)
		height := 10.0

		rectangle := Rectangle{width, height}
		got := rectangle.perimeter()
		want := 0.0

		checkPerimeter(t, got, want)
	})
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, s Shape, want float64) {
		t.Helper()
		got := s.area()

		// Assert the message
		if got != want {
			errMsg := fmt.Sprintf("got -> %f, want -> %f", got, want)
			// Print the error with colors
			t.Errorf(colorError(errMsg))
		}
	}

	// Table Tests, insert new ones here
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		// Rectangle
		{name: "Rectangle", shape: Rectangle{width: 12.0, height: 6.0}, hasArea: 72.0},
		{name: "Rectangle with negative width", shape: Rectangle{width: (-12.0), height: 6.0}, hasArea: 0.0},
		{name: "Rectangle with zero width", shape: Rectangle{width: (0.0), height: 6.0}, hasArea: 0.0},
		// Circle
		{name: "Circle", shape: Circle{radius: 10.0}, hasArea: 314.1592653589793},
		{name: "Circle with negative width", shape: Circle{radius: (-10.0)}, hasArea: 0.0},
		{name: "Circle with zero radius", shape: Circle{radius: 0.0}, hasArea: 0.0},
		// Triangle
		{name: "Triangle", shape: Triangle{base: 12.0, height: 6.0}, hasArea: 36.0},
		{name: "Triangle with negative base", shape: Triangle{base: (-12.0), height: 6.0}, hasArea: 0.0},
		{name: "Triangle with zero base", shape: Triangle{base: 0.0, height: 6.0}, hasArea: 0.0},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			checkArea(t, test.shape, test.hasArea)
		})
	}
}
