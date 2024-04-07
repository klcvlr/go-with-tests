package maps

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {

	t.Run("find word in dictionary", func(t *testing.T) {
		definition := "this is just a test"
		word := "test"
		dictionary := Dictionary{word: definition}

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{}
		actual, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("Expected to have an error but there was none", err)
		}
		assertStrings(t, "", actual)
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertEmptyError(t, err)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "existing definition"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new definition")

		assertError(t, err, ErrExistingDefinition)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("update word", func(t *testing.T) {
		word := "test"
		newDefinition := "new definition"
		dictionary := Dictionary{word: "existing definition"}

		err := dictionary.Update(word, newDefinition)

		assertEmptyError(t, err)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("update non existing word", func(t *testing.T) {
		word := "test"
		newDefinition := "new definition"
		dictionary := Dictionary{}

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "some definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	assertError(t, err, ErrNotFound)
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	result, err := dictionary.Search(word)

	assertEmptyError(t, err)
	assertStrings(t, result, definition)
}

func assertStrings(t testing.TB, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func assertError(t testing.TB, actualError error, expectedError error) {
	t.Helper()
	if !errors.Is(actualError, expectedError) {
		t.Errorf("Expected %v, but got %v", actualError, expectedError)
	}
}

func assertEmptyError(t testing.TB, err error) {
	if err != nil {
		t.Fatal("Got an error but did not expect one", err)
	}
}
