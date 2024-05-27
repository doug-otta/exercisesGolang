package iteration

import (
	"testing"
)

func TestIterarion(t *testing.T) {
	checkMessage := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("RESULT: %s - EXPECTED: %s", result, expected)
		}
	}

	t.Run("Attempt 5 c", func(t *testing.T) {
		result := Iteration(5, "c")
		expected := "ccccc"

		checkMessage(t, result, expected)
	})

	t.Run("Attempt 6 d", func(t *testing.T) {
		result := Iteration(7, "d")
		expected := "ddddddd"

		checkMessage(t, result, expected)
	})
}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iteration(5, "a")
	}
}
