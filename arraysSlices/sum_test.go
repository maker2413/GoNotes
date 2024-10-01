package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	got := Sum(numbers)
	want := 6

	if got != want {
		// %v is used to print the "default" format, which works well for arrays
		t.Errorf("expected: '%d', but got: '%d', input: %v", want, got, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// The reflect package is used to reflect data between data types.
	// The DeepEqual function will compare the values within our slices in this case.
	// NOTE: It's important to note that reflect.DeepEqual is not "type safe" - the code
	// will compile even if you did something a bit silly.
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected: '%v', but got: '%v'", want, got)
	}
}

func TestSumAllTails(t *testing.T) {
	// We could've created a new function checkSums like we normally do, but in this
	// case, we're showing a new technique, assigning a function to a variable. It might
	// look strange but, it's no different to assigning a variable to a string, or an
	// int, functions in effect are values too.

	// It's not shown here, but this technique can be useful when you want to bind a
	// function to other local variables in "scope" (e.g between some {}). It also allows
	// you to reduce the surface area of your API.
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected: '%v', but got: '%v'", want, got)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9})
		want := []int{5, 9}
		// The reflect package is used to reflect data between data types.
		// The DeepEqual function will compare the values within our slices in this case.
		// NOTE: It's important to note that reflect.DeepEqual is not "type safe" - the code
		// will compile even if you did something a bit silly.

		checkSums(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
