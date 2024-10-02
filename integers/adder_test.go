package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	assertCorrectAnswer(t, sum, expected)
}

func assertCorrectAnswer(t testing.TB, sum, expected int) {
	t.Helper()
	if sum != expected {
		// %d is used to format integers
		t.Errorf("expected: '%d', but got: '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
