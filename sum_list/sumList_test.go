package sumlist

import "testing"

func TestSumList(t *testing.T) {
	t.Run("Sum sorted numbers", func(t *testing.T) {
		result := sumList([]int{1, 2, 3, 4, 5})
		expected := 15

		if result != expected {
			t.Errorf("RESULT: %v - EXPECTED: %v", result, expected)
		}
	})

	t.Run("Sum unsorted numbers", func(t *testing.T) {
		result := sumList([]int{7, 28, 45, 23})
		expected := 103

		if result != expected {
			t.Errorf("RESULT: %v - EXPECTED: %v", result, expected)
		}
	})

	t.Run("Sum with negative numbers", func(t *testing.T) {
		result := sumList([]int{7, -28, 45, -23})
		expected := 1

		if result != expected {
			t.Errorf("RESULT: %v - EXPECTED: %v", result, expected)
		}
	})

	t.Run("NULL value", func(t *testing.T) {
		result := sumList([]int{})
		expected := 0

		if result != expected {
			t.Errorf("RESULT: %v - EXPECTED: %v", result, expected)
		}
	})
}
