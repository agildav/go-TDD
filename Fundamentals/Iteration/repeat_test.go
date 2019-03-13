package main

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		// Assert the message
		if got != want {
			errMsg := fmt.Sprintf("got -> %s, want -> %s", got, want)
			// Print the error with colors
			t.Errorf(colorError(errMsg))
		}
	}

	t.Run("repeating 7 times", func(t *testing.T) {
		const n = 7
		got := repeat("a", n)
		want := "aaaaaaa"

		assertCorrectMessage(t, got, want)
	})
}
