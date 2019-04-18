package maps

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	d := Dictionary{"test": "dictionary test string"}
	t.Run("known word", func(t *testing.T) {
		got, _ := d.Search("test")
		want := "dictionary test string"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := d.Search("unknown")

		assertError(t, err, errWordNotFound)
	})
}

func TestAdd(t *testing.T) {
	d := Dictionary{}
	t.Run("add a word", func(t *testing.T) {
		d.Add("test", "dictionary test string")
		got, err := d.Search("test")
		want := "dictionary test string"

		if err != nil {
			t.Fatal("error adding word,", err)
		}

		assertStrings(t, got, want)
	})
}

// // // // // // // // // // // // // // // // // // // // // // // // // //

// Helpers
func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		errMsg := fmt.Sprintf("got -> %s, want -> %s", got, want)
		t.Errorf(colorError(errMsg))
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		errMsg := fmt.Sprintf("got -> %s, want -> %s", got, want)
		t.Errorf(colorError(errMsg))
	}
}
