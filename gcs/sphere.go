package gcs

import "math"

type (
	// Sphere is a geometry represented as the set of points
	// that are all at the same distance R from a given point,
	// called center or origin.
	Sphere struct {
		R float64 // radius in meters
	}

	// SPoint is a point in the surface of the Sphere, ie, its
	// distance from the origin is the radius of the sphere.
	SPoint struct {
		φ Angle // latitude in radians
		λ Angle // longitude in radians
	}

	// Point in a spherical coordinate system
	Point struct {
		R float64 // radial distance from origin
		φ Angle   // latitude in radians
		λ Angle   // longitude in radians
	}
)

const (
	EarthRadius = 6378100 // Earth radius at equator in meters
	MarsRadius  = 3396200 // Mars radius at equator in meters
)

var (
	// SphericalUnit is the unit sphere
	SphericalUnit = Sphere{
		R: 1,
	}

	// SphericalEarth is the spherical approximation of earth
	SphericalEarth = Sphere{
		R: EarthRadius,
	}

	// SphericalMars is the spherical approximation of mars
	SphericalMars = Sphere{
		R: MarsRadius,
	}
)

// NewSPoint creates a point in the surface of a sphere.
// Lon and lat are in degrees.
func NewSPoint(lon, lat float64) SPoint {
	return SPoint{
		λ: radians(lon),
		φ: radians(lat),
	}
}

// Area gives the surface area
func (s *Sphere) Area() float64 {
	return 4 * math.Pi * s.R * s.R
}

func (s *Sphere) Contains(p Point) bool {
	return p.R <= s.R
}

// func (s *Sphere) Distance(p1, p2 Point) float64 {

// }
