package maps

import (
	"errors"
)

type Dictionary map[string]string

// // // // // // // // // // // // // // // // // // // // // // // // // //

// Dictionary Errors
var errWordNotFound = errors.New("could not find word, it does not exist")
var errWordExists = errors.New("could not add word, it already exists")

// // // // // // // // // // // // // // // // // // // // // // // // // //

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", errWordNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, word string) error {
	_, err := d.Search(key)

	switch err {
	case errWordNotFound:
		d[key] = word
	case nil:
		return errWordExists
	default:
		return err
	}

	return nil
}
