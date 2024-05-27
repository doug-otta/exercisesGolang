package main

import (
	"bytes"
	"testing"
)

func TestGreetings(t *testing.T) {
	buffer := bytes.Buffer{}
	Greetings(&buffer, "Chris")

	result := buffer.String()
	expected := "Hello, Chris"

	t.Errorf("RESULT: '%s' - EXPECTED: '%s'", result, expected)

	if result != expected {
		t.Errorf("RESULT: '%s' - EXPECTED: '%s'", result, expected)
	}
}
