package main

import "testing"

func TestFindNumberInList(t *testing.T) {
	t.Run("Existing Number", func(t *testing.T) {
		result := findNumber([]int{1, 2, 3, 4, 5, 6}, 3)
		expected := true

		if result != expected {
			t.Errorf("RESULT: %v - EXPECTED: %v", result, expected)
		}
	})

	t.Run("Not existing Number", func(t *testing.T) {
		result := findNumber([]int{1, 2, 3, 4}, 7)
		expected := false

		if result != expected {
			t.Errorf("RESULT: %v - EXPECTED: %v", result, expected)
		}
	})

	t.Run("Not existing List", func(t *testing.T) {
		result := findNumber(nil, 7)
		expected := false

		if result != expected {
			t.Errorf("RESULT: %v - EXPECTED: %v", result, expected)
		}
	})
}
