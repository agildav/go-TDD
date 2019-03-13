package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		// Color formatting
		const (
			BeginErrorColor = "\033[1;31m"
			EndErrorColor   = "\033[0m"
		)

		// Color the error message
		colorError := func(s string) string {
			return BeginErrorColor + s + EndErrorColor
		}

		// Assert the message
		if got != want {
			errMsg := fmt.Sprintf("got %s, want %s", got, want)
			t.Errorf(colorError(errMsg))
		}
	}

	t.Run("saying hello", func(t *testing.T) {
		got := Hello("John")
		want := "Hello, John"

		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to world", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in spanish", func(t *testing.T) {
		got := Hello("John")
		want := "Hola, John"

		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to mundo in spanish", func(t *testing.T) {
		got := Hello("")
		want := "Hola, mundo"

		assertCorrectMessage(t, got, want)
	})
}
