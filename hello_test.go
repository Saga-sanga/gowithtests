package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Sanga", "")
		want := "Hello Sanga!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello World!' when an empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Karlos", "Spanish")
		want := "Hola Karlos!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Karles", "French")
		want := "Bonjour Karles!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Mizo", func(t *testing.T) {
		got := Hello("Karles", "Mizo")
		want := "Chibai Karles!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
