package strings

import "testing"

func TestSayHello(t *testing.T) {
	t.Run("Say hello with empty parameter", func(t *testing.T) {
		actual := Hello("", "")
		expected := "Hello, World!"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("Say hello with white space parameter", func(t *testing.T) {
		actual := Hello(" 	", "") // space and tab
		expected := "Hello, World!"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("Say hello to person", func(t *testing.T) {
		actual := Hello("John", "")
		expected := "Hello, John!"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("Say hello to person, in english", func(t *testing.T) {
		actual := Hello("John", "english")
		expected := "Hello, John!"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("Say hello to person, in spanish", func(t *testing.T) {
		actual := Hello("John", "spanish")
		expected := "Hola, John!"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("Say hello to person, in french", func(t *testing.T) {
		actual := Hello("John", "french")
		expected := "Bonjour, John!"

		assertCorrectMessage(t, actual, expected)
	})
}

func assertCorrectMessage(t testing.TB, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}
