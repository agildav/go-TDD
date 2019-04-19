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
	d := Dictionary{"key": "word"}
	t.Run("add a word", func(t *testing.T) {
		d.Add("test", "dictionary test string")
		got, err := d.Search("test")
		want := "dictionary test string"

		if err != nil {
			t.Fatal("error adding word,", err)
		}

		assertStrings(t, got, want)
	})

	t.Run("should not add a word that already exists", func(t *testing.T) {
		err := d.Add("key", "new word test")
		want := errWordExists

		if err != nil {
			assertError(t, err, want)
		}
	})
}

func TestUpdate(t *testing.T) {
	d := Dictionary{"key": "word"}
	t.Run("update a word", func(t *testing.T) {
		err := d.Update("key", "dictionary update test string")
		want := "dictionary update test string"

		if err != nil {
			t.Fatal("error updating word,", err)
		}
		got, _ := d.Search("key")
		assertStrings(t, got, want)
	})

	t.Run("should not update a word that does not exist", func(t *testing.T) {
		err := d.Update("unknown", "new word test")
		want := errUpdWordDoesNotExist

		if err != nil {
			assertError(t, err, want)
		}
	})
}

func TestDelete(t *testing.T) {
	d := Dictionary{"key": "word"}
	t.Run("delete a word", func(t *testing.T) {
		d.Delete("key")

		_, err := d.Search("key")
		want := errWordNotFound
		assertError(t, err, want)
	})

	t.Run("delete a word that does not exist", func(t *testing.T) {
		d.Delete("newKey")

		_, err := d.Search("newKey")
		want := errWordNotFound
		assertError(t, err, want)
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
