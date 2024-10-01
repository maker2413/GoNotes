package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Ethan")
	want := "Hello, Ethan!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
