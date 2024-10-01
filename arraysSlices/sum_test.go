package main

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15
	assertCorrectAnswer(t, got, want, numbers)
}

func assertCorrectAnswer(t testing.TB, got, want int, numbers [5]int) {
	t.Helper()
	if got != want {
		// %v is used to print the "default" format, which works well for arrays
		t.Errorf("expected: '%d', but got: '%d', input: %v", want, got, numbers)
	}
}
