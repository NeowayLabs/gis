package geojson

import "testing"

func invalidType(t *testing.T, geojson string) {
	_, err := Decode([]byte(geojson))
	if err == nil {
		t.Fatalf("Geojson '%s' is invalid", geojson)
	}
}

func testSinglePointsOK(t *testing.T, geojson string, expected Point) {
	obj, err := Decode([]byte(geojson))
	if err != nil {
		t.Fatal(err)
	}

	assertPoint(t, obj, expected)
}

func TestDecodeInvalidTypes(t *testing.T) {
	for _, value := range []string{
		`{"some": "val"}`,     // no type
		`{"type": "invalid"}`, // invalid type
		`{"type": ""}`,        // empty type
		// TODO(i4k): Add invalid types for nested geojson objects also
	} {
		invalidType(t, value)
	}
}

func TestDecodePointOK(t *testing.T) {
	for _, tc := range []struct {
		geojson string
		point   Point
	}{
		{
			geojson: `{
			"type": "Point",
			"coordinates": [10, 10]
		}`,
			point: Point{
				Coordinate: Position{
					Lat: 10,
					Lon: 10,
				},
			},
		},
		{
			geojson: `{
			"type": "Point",
			"coordinates": [1, 0]
		}`,
			point: Point{
				Coordinate: Position{
					Lat: 0,
					Lon: 1,
				},
			},
		},
	} {
		tc := tc
		testSinglePointsOK(t, tc.geojson, tc.point)
	}
}
