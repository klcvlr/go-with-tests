package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		actual := Sum(numbers)
		expected := 15

		if expected != actual {
			t.Errorf("Expected %d, but got %d given %v", expected, actual, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	actual := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	checkSums(t, actual, expected)
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum all tails of non-empty arrays", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		checkSums(t, actual, expected)
	})

	t.Run("safely run on empty slices", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}

		checkSums(t, actual, expected)

	})
}

func TestPlayingWithSlices(t *testing.T) {
	t.Run("subSlice points to original underlying slice", func(t *testing.T) {
		original := []int{5, 5, 5}
		expected := []int{5, 6, 5}

		subSlice := original[1:]
		subSlice[0] = 6

		if !reflect.DeepEqual(original, expected) {
			t.Errorf("Expected original %v to be equal to %v", original, expected)
		}
	})

	t.Run("subSlice length is different from the original array", func(t *testing.T) {
		original := []int{5, 5, 5}
		subSlice := original[1:]
		subSliceLength := len(subSlice)
		expected := 2

		if subSliceLength != expected {
			t.Errorf("Expected subSubslice %v to have length %d", subSlice, expected)
		}
	})

}

func checkSums(t *testing.T, actual, expected []int) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
