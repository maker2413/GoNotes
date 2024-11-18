package main

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("Assert Equal", func(t *testing.T) {
		AssertEqual(t, 1, 1)
	})
	t.Run("Assert Not Equal", func(t *testing.T) {
		AssertNotEqual(t, 1, 2)
	})
	t.Run("Assert True", func(t *testing.T) {
		AssertTrue(t, 1 == 1)
	})
	t.Run("Assert False", func(t *testing.T) {
		AssertFalse(t, 1 == 2)
	})
}
