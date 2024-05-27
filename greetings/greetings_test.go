package main

import (
	"testing"
)

func TestGreetings(t *testing.T) {
	checkCorrectMessage := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("Result: '%s' - Expected: '%s'", result, expected)
		}
	}

	t.Run("say Hello to people", func(t *testing.T) {
		result := greetings("Douglas", "english")
		expected := "Hello, Douglas"

		checkCorrectMessage(t, result, expected)
	})

	t.Run("'World' as standard to 'empty' string", func(t *testing.T) {
		result := greetings("", "")
		expected := "Hello, World"
		checkCorrectMessage(t, result, expected)
	})

	t.Run("Hello in portuguese", func(t *testing.T) {
		result := greetings("Maria", "portuguese")
		expected := "Olá, Maria"

		checkCorrectMessage(t, result, expected)
	})

	t.Run("Hello in french", func(t *testing.T) {
		result := greetings("Joana", "french")
		expected := "Bonjour, Joana"

		checkCorrectMessage(t, result, expected)
	})
}
