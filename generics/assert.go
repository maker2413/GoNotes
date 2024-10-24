package generics

import "testing"

// To write generic functions in Go, you need to provide "type parameters" which is just
// a fancy way of saying "describe your generic type and give it a label".
// In our case the type of our type parameter is comparable and we've given it the label
// of T. This label then lets us describe the types for the arguments to our
// function (got, want T). We're using comparable because we want to describe to the compiler
// that we wish to use the == and != operators on things of type T in our function, we want
// to compare!
func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %v", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
