package maps

import (
	"errors"
)

var ErrNotFound = DictionaryErr("could not find the word you were looking for")
var ErrExistingDefinition = DictionaryErr("could not add definition: word already exist")
var ErrWordDoesNotExist = DictionaryErr("could not update word because it does not exist")

type Dictionary map[string]string
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch {
	case errors.Is(err, ErrNotFound):
		d[word] = definition
		return nil
	case err == nil:
		return ErrExistingDefinition
	default:
		return err
	}
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch {
	case errors.Is(err, ErrNotFound):
		return ErrWordDoesNotExist
	case err == nil:
		d[word] = definition
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
