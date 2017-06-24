package geojson

import (
	"encoding/json"
)

type (
	// Position is the point coordinates in WGS84
	// geographic coordinate system.
	// See: https://tools.ietf.org/html/rfc7946#section-3.1.1
	Position struct {
		Lon float64 // Longitude
		Lat float64 // Latitude
		Alt float64 // Altitude (optional)
	}

	FeatureCollection struct {
		Features []Feature  `json:"features"`
		bbox     []Position `json:"bbox,omitempty"`
	}

	Feature struct {
		ID         *json.Number `json:"id"`
		properties interface{}  `json:"properties"`
		geometry   Geometry     `json:"geometry"`
		bbox       []Position   `json:"bbox,omitempty"`
	}

	Object interface {
		Type() string // Type of the GeoJSON object
	}

	Geometry interface {
		Type() string // Type of a geometry
	}

	Point struct {
		// Coordinates (sic) of the point
		Coordinates Position `json:"coordinates"`
	}

	MultiPoint struct {
		coordinates []Position `json:"coordinates"`
		bbox        []Position `json:"bbox,omitempty"`
	}

	LineString struct {
		coordinates []Position
		bbox        []Position `json:"bbox,omitempty"`
	}

	MultiLineString struct {
		coordinates [][]Position
		bbox        []Position `json:"bbox,omitempty"`
	}

	// Polygon is comprised of linear rings (closed LinearString).
	// The first LinearString coordinates are the exterior ring and
	// others (if any) are the interior rings (holes within the
	// exterior ring).
	// See: https://tools.ietf.org/html/rfc7946#section-3.1.6
	Polygon struct {
		Coordinates [][]Position `json:"coordinates"`
		bbox        []Position   `json:"bbox,omitempty"`
	}

	// MultiPolygon is a list of polygons
	MultiPolygon struct {
		coordinates [][][]Position `json:"coordinates"`
		bbox        []Position     `json:"bbox,omitempty"`
	}

	GeometryCollection struct {
		geometries []Geometry `json:"geometries"`
		bbox       []Position `json:"bbox,omitempty"`
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
