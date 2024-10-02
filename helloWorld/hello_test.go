package main

import "testing"

func TestHello(t *testing.T) {
	//t.Run is how we can define subtests, which allow us to group together tests in a func
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Ethan", "")
		want := "Hello, Ethan!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, world!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in Spanish to people", func(t *testing.T) {
		got := Hello("Ethan", "Spanish")
		want := "Hola, Ethan!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in French to people", func(t *testing.T) {
		got := Hello("Ethan", "French")
		want := "Bonjour, Ethan!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in Japanese to people", func(t *testing.T) {
		got := Hello("Ethan", "Japanese")
		want := "Konnichiwa, Ethan!"
		assertCorrectMessage(t, got, want)
	})
}

// testing.TB is an interface that accepts tests and benchmarks
func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper will make the failure message show the line number test that called this function
	t.Helper()
	if got != want {
		// %q is used to format strings
		t.Errorf("got: %q, want: %q", got, want)
	}
}
