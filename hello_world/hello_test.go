package main

import "testing"

func TestHello(t *testing.T) {
	result := hello()
	expected := "Hello World"

	if result != expected {
		t.Errorf("result is: '%s', expected is: '%s'", result, expected)
	}
}
