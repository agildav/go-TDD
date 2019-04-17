package maps

import (
	"errors"
)

type Dictionary map[string]string

// // // // // // // // // // // // // // // // // // // // // // // // // //

// Dictionary Errors
var errWordNotFound = errors.New("could not find word, does not exist")

// // // // // // // // // // // // // // // // // // // // // // // // // //

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", errWordNotFound
	}

	return definition, nil
}
