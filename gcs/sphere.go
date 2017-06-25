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
		λ: toRadians(lon),
		φ: toRadians(lat),
	}
}

// Area gives the surface area
func (s *Sphere) Area() float64 {
	return 4 * math.Pi * s.R * s.R
}

func (s *Sphere) Contains(p Point) bool {
	return p.R <= s.R
}

// Distance of p1 and p2 throught the surface of the sphere (orthodromic
// distance or great circle) returned in meters.
// The algorithm uses the haversine formula and because of that it has
// lower precision for computing distance of antipodal points.
func (s *Sphere) Distance(p1, p2 SPoint) float64 {
	return haversin(s.R, p1, p2)
}

// haversin function of angle θ
//   hsin(θ) = sin²(θ/2)
func hsin(θ float64) float64 {
	return math.Pow(math.Sin(θ/2), 2)
}

// haversin of α angle between points p1 and p2.
// The formula is:
//   haversin(α) = hsin(φ2-φ1)+cosφ1.cosφ2.hsin(λ1-λ2)
//   haversin(α) = (d/2R)²
// then:
//   d = 2Rsin⁻¹√(haversin(α))
// Based on:
//   https://www.math.ksu.edu/~dbski/writings/haversine.pdf
func haversin(R float64, p1, p2 SPoint) float64 {
	φ1, λ1 := float64(p1.φ), float64(p1.λ)
	φ2, λ2 := float64(p2.φ), float64(p2.λ)
	h := hsin(φ2-φ1) + math.Cos(φ1)*math.Cos(φ2)*hsin(λ2-λ1)
	return 2 * R * math.Asin(math.Sqrt(h))
}
