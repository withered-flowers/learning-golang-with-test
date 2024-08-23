package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Someone", "english")
		want := "Hello World (Someone)"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying Hello World when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola (Elodie)"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
