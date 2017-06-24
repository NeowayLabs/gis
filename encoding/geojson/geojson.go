package geojson

import (
	"encoding/json"
)

type (
	// Position is the point coordinates in the
	// geographic coordinate system.
	// See: https://tools.ietf.org/html/rfc7946#section-3.1.1
	Position struct {
		Lon float64 // Longitude
		Lat float64 // Latitude
		Alt float64 // Altitude (optional)
	}

	FeatureCollection struct {
		Features []Feature
		bbox     []Position
	}

	Feature struct {
		ID         *json.Number
		properties interface{}
		geometry   Geometry
		bbox       []Position
	}

	Object interface {
		Type() string // Type of the GeoJSON object
	}

	Geometry interface {
		Type() string // Type of a geometry
	}

	Point struct {
		Coordinate Position
	}

	MultiPoint struct {
		coordinates []Position
		bbox        []Position
	}

	LineString struct {
		coordinates []Position
		bbox        []Position
	}

	MultiLineString struct {
		coordinates [][]Position
		bbox        []Position
	}

	// Polygon is comprised of linear rings (closed LinearString).
	// The first LinearString coordinates are the exterior ring and
	// others (if any) are the interior rings (holes within the
	// exterior ring).
	// See: https://tools.ietf.org/html/rfc7946#section-3.1.6
	Polygon struct {
		coordinates [][]Position
		bbox        []Position
	}

	// MultiPolygon is a list of polygons
	MultiPolygon struct {
		coordinates [][][]Position
		bbox        []Position
	}

	GeometryCollection struct {
		geometries []Geometry
		bbox       []Position
	}

	geojsonType []string
)

var (
	geometryTypes = []string{
		"Point",
		"MultiPoint",
		"LineString",
		"MultiLineString",
		"Polygon",
		"MultiPolygon",
		"GeometryCollection",
	}
	geojsonTypes geojsonType = append([]string{
		"Feature",
		"FeatureCollection",
	}, geometryTypes...)
)

func (typs geojsonType) Valid(typ string) bool {
	for _, v := range typs {
		if v == typ {
			return true
		}
	}

	return false
}

func (p Point) Type() string { return "Point" }
