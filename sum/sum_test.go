package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Colection of any length", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		result := Sum(numbers)
		expected := 6

		if result != expected {
			t.Errorf("RESULT: %d - EXPECTED: %d", result, expected)
		}
	})
}

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("RESULT: %d - EXPECTED: %d", result, expected)
	}
}

func TestSumOthers(t *testing.T) {
	verifySum := func(t *testing.T, result, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("RESULT: %d - EXPECTED: %d", result, expected)
		}
	}
	t.Run("Do sum of some slice", func(t *testing.T) {
		result := SumOthers([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		verifySum(t, result, expected)
	})

	t.Run("Sum empty slice", func(t *testing.T) {
		result := SumOthers([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}

		verifySum(t, result, expected)
	})
}
