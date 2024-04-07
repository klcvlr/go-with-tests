package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "John")

	actual := buffer.String()
	expected := "Hello, John"

	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}
