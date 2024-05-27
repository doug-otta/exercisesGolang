package module01

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	result := fizzBuzz(5)
	expected := "1, 2, Fizz, 4, Buzz"

	if reflect.DeepEqual(result, expected) {
		t.Errorf("RESULT: %s - EXPECTED: %s", result, expected)
	}

}
