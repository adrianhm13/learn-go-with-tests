package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("In spanish", func(t *testing.T) {
		got := Hello("Adrian", "Spanish")
		want := "Hola, Adrian"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In french", func(t *testing.T) {
		got := Hello("Adrian", "French")
		want := "Bonjour, Adrian"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Should say Hello with the name provided", func(t *testing.T) {
		got := Hello("Adrian", "")
		want := "Hello, Adrian"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Should say Hello, World if name is not provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	// Method to let know the line where the test broke instead of the one from the helper function
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
