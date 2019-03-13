package main

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()

		// Assert the message
		if got != want {
			errMsg := fmt.Sprintf("got -> %d, want -> %d", got, want)
			// Print the error with colors
			t.Errorf(colorError(errMsg))
		}
	}

	t.Run("adding 2 integers", func(t *testing.T) {
		got := add(3, 65)
		want := 68

		assertCorrectMessage(t, got, want)
	})
}
