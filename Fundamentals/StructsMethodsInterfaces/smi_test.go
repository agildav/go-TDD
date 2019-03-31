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

	t.Run("negative values of a rectangle give zero perimeter", func(t *testing.T) {
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

	// Rectangles

	t.Run("area of a rectangle", func(t *testing.T) {
		width := 20.0
		height := 10.0

		rectangle := Rectangle{width, height}

		want := 200.0

		checkArea(t, rectangle, want)
	})

	t.Run("negative values of a rectangle give zero area", func(t *testing.T) {
		width := (-20.0)
		height := 10.0

		rectangle := Rectangle{width, height}
		want := 0.0

		checkArea(t, rectangle, want)
	})

	// Circles

	t.Run("area of a circle", func(t *testing.T) {
		radius := 10.0

		circle := Circle{radius}
		want := 314.1592653589793

		checkArea(t, circle, want)
	})

	t.Run("negative radius of a circle give zero area", func(t *testing.T) {
		radius := (-10.0)

		circle := Circle{radius}
		want := 0.0

		checkArea(t, circle, want)
	})

}
