package reversestring

import "testing"

func TestReverseString(t *testing.T) {
	result := StringReverse("cat")
	expected := "tac"

	if result != expected {
		t.Errorf("RESULT: %v - EXPECTED: %v", result, expected)
	}
}
