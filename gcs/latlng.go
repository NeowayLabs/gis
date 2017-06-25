package gcs

import "math"

type (
	Angle float64 // Angle in radians
)

func radians(angle float64) Angle {
	return Angle(angle * math.Pi / 180)
}

// Radians returns the value in radians
func (a Angle) Radians() float64 {
	return float64(a)
}

// Degrees returns the value in degrees
func (a Angle) Degrees() float64 {
	return float64(a * 180 / math.Pi)
}
