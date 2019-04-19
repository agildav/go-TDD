package maps

type Dictionary map[string]string

// // // // // // // // // // // // // // // // // // // // // // // // // //

// Dictionary Errors
type DictionaryErr string

const (
	errWordNotFound        = DictionaryErr("could not find word because it does not exist")
	errWordExists          = DictionaryErr("could not add word because it already exists")
	errUpdWordDoesNotExist = DictionaryErr("could not update word because it does not exist")
	errDelWordDoesNotExist = DictionaryErr("could not delete word because it does not exist")
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

func (d Dictionary) Update(key, word string) error {
	_, err := d.Search(key)

	switch err {
	case errWordNotFound:
		return errUpdWordDoesNotExist
	case nil:
		d[key] = word
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
