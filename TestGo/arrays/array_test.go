package arrays

import (
	"testing"
)

func TestArray(t *testing.T) {

	t.Run("Collection of fixed size (Array)", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("Expected %d, got %d", expected, got)
		}
	})

	t.Run("Collection of variable size (Slice)", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumSlice(numbers)
		expected := 6

		if got != expected {
			t.Errorf("Expected %d, got %d", expected, got)
		}
	})
}
