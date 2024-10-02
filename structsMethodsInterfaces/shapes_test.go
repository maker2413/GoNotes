package main

import "testing"

// Interfaces allow us to define functions that multiple types can use.
type Shape interface {
	Area() float64
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		// f is for float64 and the .2 means print out 2 decimal places.
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		// invoke the Area method for the rectangle struct
		got := shape.Area()
		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}

// This function is showing how to do table driven tests
func TestAreaTable(t *testing.T) {
	// areaTests is an "anonymous struct".
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		// Naming the input values is optional, but can be helpful for reading the code:
		// {"Triangle", shape: Triangle{Width: 12, Height: 6}, want: 32.0},
		{"Triangle", Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				// The v input is going to dictate which shape the test failed on.
				t.Errorf("%v got %g, want %g", tt.shape, got, tt.want)
			}
		})
	}
}
