package integers

import "testing"

func TestAdder(t *testing.T) {
	actual := Add(10, 25)
	expected := 35

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
