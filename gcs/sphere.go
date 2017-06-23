package gcs

import "math"

type (
	// Sphere is a geometry represented as the set of points
	// that are all at the same distance R from a given point,
	// called center or origin.
	Sphere struct {
		R float64 // radius
	}

	// SPoint is a point in the surface of the Sphere, ie, its
	// distance from the origin is the radius of the sphere.
	SPoint struct {
		φ float64 // latitude in degree
		λ float64 // longitude in degree
	}

	// Point in a spherical coordinate system
	Point struct {
		R float64 // radial distance from origin
		φ float64 // latitude in degree
		λ float64 // longitude in degree
	}
)

// Area gives the surface area
func (s *Sphere) Area() float64 {
	return 4 * math.Pi * s.R * s.R
}

func (s *Sphere) Contains(p Point) bool {
	return p.R <= s.R
}
