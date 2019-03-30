package arraysslices

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()

		// Assert the message
		if got != want {
			errMsg := fmt.Sprintf("got -> %d, want -> %d", got, want)
			// Print the error with colors
			t.Errorf(colorError(errMsg))
		}
	}

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := sum(numbers)
		want := 15

		assertCorrectMessage(t, got, want)
	})

	t.Run("collection with negative numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 5, -4, 5, 9, -1, -2, -18, -2, -3, 5, 1}
		got := sum(numbers)
		want := 1

		assertCorrectMessage(t, got, want)
	})
}
