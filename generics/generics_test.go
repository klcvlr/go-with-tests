package generics

import "testing"

func TestStack(t *testing.T) {
	t.Run("can pop data from a stack", func(t *testing.T) {
		stack := NewStack[int]()
		want := 3
		stack.Push(3)

		got, err := stack.Pop()

		if err != nil {
			t.Fatal("got an error but didn't expect one", err)
		}
		if got != want {
			t.Errorf("wanted: %v, got %v", want, got)
		}
	})

	t.Run("stack works with strings", func(t *testing.T) {
		stack := NewStack[string]()
		want := "apple"
		stack.Push("apple")

		got, err := stack.Pop()

		if err != nil {
			t.Fatal("got an error but didn't expect one", err)
		}
		if got != want {
			t.Errorf("wanted: %v, got %v", want, got)
		}
	})

	t.Run("data are retrieved in reverse insertion order", func(t *testing.T) {
		stack := NewStack[int]()
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		want := [3]int{3, 2, 1}

		first, _ := stack.Pop()
		second, _ := stack.Pop()
		third, _ := stack.Pop()
		got := [3]int{first, second, third}

		if got != want {
			t.Errorf("wanted: %v, got %v", want, got)
		}
	})
}

func TestAssertFunctions(t *testing.T) {
	AssertEqual(t, 1, 1)
	AssertNotEqual(t, 1, 2)

	AssertEqual(t, "hello", "hello")
	AssertNotEqual(t, "hello", "world")
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want: %v", got)
	}
}
