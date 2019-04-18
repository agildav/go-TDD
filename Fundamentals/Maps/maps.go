package maps

type Dictionary map[string]string

// // // // // // // // // // // // // // // // // // // // // // // // // //

// Dictionary Errors
type DictionaryErr string

const (
	errWordNotFound = DictionaryErr("could not find word, it does not exist")
	errWordExists   = DictionaryErr("could not add word, it already exists")
)

func (err DictionaryErr) Error() string {
	return string(err)
}

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
