package main

import (
	"reflect"
	"testing"
)

func TestSumIndex(t *testing.T) {
	t.Run("First test", func(t *testing.T) {
		numToSum := []int{2, 7, 11, 15}
		result := twoSum(numToSum, 9)
		expected := []int{0, 1}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("RESULT: %v - EXPECTED: %v", result, expected)
		}
	})

	t.Run("Second test", func(t *testing.T) {
		numToSum := []int{3, 2, 4}
		result := twoSum(numToSum, 6)
		expected := []int{1, 2}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("RESULT: %v - EXPECTED: %v", result, expected)
		}
	})

	t.Run("Third test", func(t *testing.T) {
		numToSum := []int{3, 3}
		result := twoSum(numToSum, 6)
		expected := []int{0, 1}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("RESULT: %v - EXPECTED: %v", result, expected)
		}
	})

}
