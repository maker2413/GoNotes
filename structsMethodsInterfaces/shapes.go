package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Width  float64
	Height float64
}

func Perimeter(shape Rectangle) float64 {
	return 2 * (shape.Width + shape.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	// The math library has basic math functions like pi or Powers of.
	return math.Pi * (math.Pow(c.Radius, 2))
}

func (t Triangle) Area() float64 {
	return .5 * (t.Width * t.Height)
}

func main() {
	shape := Rectangle{10.0, 10.0}
	fmt.Println(Perimeter(shape))
}
