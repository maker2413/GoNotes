package main

import (
	"errors"
	"reflect"
	"testing"
)

func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got: '%v', want: '%v'", got, want)
	}
}

func assertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("got: '%v', want: '%v'", got, want)
	}
}

func assertDeepEqual[T comparable](t *testing.T, got, want []T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: '%v', want: '%v'", got, want)
	}
}

func assertNotDeepEqual[T comparable](t *testing.T, got, want []T) {
	t.Helper()
	if reflect.DeepEqual(got, want) {
		t.Errorf("got: '%v', want: '%v'", got, want)
	}
}

func assertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got: %v, want: true", got)
	}
}

func assertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got: %v, want: false", got)
	}
}

func assertError(t *testing.T, got error) {
	t.Helper()
	if got == nil {
		t.Errorf("Expected to receive an error and got: %v", got)
	}
}

func assertNotError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("Expected not to receive an error and got: %v", got)
	}
}

func TestAssertFunctions(t *testing.T) {
	t.Run("Assert Equal", func(t *testing.T) {
		assertEqual(t, 1, 1)
	})
	t.Run("Assert Not Equal", func(t *testing.T) {
		assertNotEqual(t, 1, 2)
	})
	t.Run("Assert True", func(t *testing.T) {
		assertTrue(t, 1 == 1)
	})
	t.Run("Assert False", func(t *testing.T) {
		assertFalse(t, 1 == 2)
	})
	t.Run("Assert Deep Equal", func(t *testing.T) {
		assertDeepEqual(t, []int{1, 2, 3}, []int{1, 2, 3})
	})
	t.Run("Assert Not Deep Equal", func(t *testing.T) {
		assertNotDeepEqual(t, []string{"a", "b", "c"}, []string{"1", "2", "3"})

		var got []int
		assertNotDeepEqual(t, got, []int{})
	})
	t.Run("Assert Error", func(t *testing.T) {
		assertError(t, errors.New("This is an error"))
	})
	t.Run("Assert Not Error", func(t *testing.T) {
		assertNotError(t, nil)
	})
}
