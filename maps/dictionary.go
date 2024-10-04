package main

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("could not add specified word, word already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type Dictionary map[string]string
type DictionaryErr string

// https://dave.cheney.net/2016/04/07/constant-errors
func (e DictionaryErr) Error() string {
	return string(e)
}

// We can have two different return values in Go. In this case we are returning
// the results of our search and an error that can be nil.
func (d Dictionary) Search(word string) (string, error) {
	// The map lookup can return 2 values. The second value is a boolean which
	// indicates if the key was found successfully.
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		// maps can be added onto by simply adding a value to a new key.
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
