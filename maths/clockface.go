package clockface

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

// A Point represents a two-dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

// secondsInRadians returns the angle of the second hand from 12 o'clock in radians.
func secondsInRadians(t time.Time) float64 {
	// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/math#floats-are-horrible
	return (math.Pi / (secondsInHalfClock / (float64(t.Second()))))
}

// secondHandPoint is the unit vector of the second hand at time `t`,.
// represented a Point.
func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

// minutesInRadians returns the angle of the minute hand from 12 o'clock in radians.
func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / minutesInClock) +
		(math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

// minuteHandPoint is the unit vector of the minute hand at time `t`,.
// represented a Point.
func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

// hoursInRadians returns the angle of the hour hand from 12 o'clock in radians.
func hoursInRadians(t time.Time) float64 {
	// We modulo to 12 because this is not a 24 clock
	return (minutesInRadians(t) / hoursInClock) +
		(math.Pi / (hoursInHalfClock / float64(t.Hour()%hoursInClock)))
}

// hourHandPoint is the unit vector of the hour hand at time `t`,.
// represented a Point.
func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
