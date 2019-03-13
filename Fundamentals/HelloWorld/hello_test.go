package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		// Assert the message
		if got != want {
			errMsg := fmt.Sprintf("got -> %s, want -> %s", got, want)
			// Print the error with colors
			t.Errorf(colorError(errMsg))
		}
	}

	t.Run("saying hello", func(t *testing.T) {
		got := hello("John", "english")
		want := "Hello, John"

		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to world", func(t *testing.T) {
		got := hello("", "english")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in spanish", func(t *testing.T) {
		got := hello("John", "spanish")
		want := "Hola, John"

		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to mundo in spanish", func(t *testing.T) {
		got := hello("", "spanish")
		want := "Hola, mundo"

		assertCorrectMessage(t, got, want)
	})
}
