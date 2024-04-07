package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repeat a character 5 times", func(t *testing.T) {
		actual := Repeat("a", 5)
		expected := "aaaaa"

		if expected != actual {
			t.Errorf("Expected %q, but got %q", expected, actual)
		}
	})

	t.Run("Repeat with only 1 time", func(t *testing.T) {
		actual := Repeat("a", 1)
		expected := "a"

		if expected != actual {
			t.Errorf("Expected %q, but got %q", expected, actual)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 1000)
	}
}
