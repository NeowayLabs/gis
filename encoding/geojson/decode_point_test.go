package geojson

import "testing"

func testSinglePointsOK(t *testing.T, geojson string, expected Point) {
	obj, err := Decode([]byte(geojson))
	if err != nil {
		t.Fatal(err)
	}

	assertPoint(t, obj, expected)
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
				Coordinates: Position{
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
				Coordinates: Position{
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

func TestDecodeGeodataInvalidPoints(t *testing.T) {
	for _, geojson := range getGeometries(t, "samples/points/invalids") {
		_, err := Decode([]byte(geojson))
		if err == nil {
			t.Fatalf("GeoJSON below is invalid: %s", string(geojson))
		}
	}
}

func TestDecodeGeodataValidPoints(t *testing.T) {
	for _, geojson := range getGeometries(t, "samples/points/valids") {
		_, err := Decode([]byte(geojson))
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestDecodeGeodataValidMultiPoints(t *testing.T) {
	for _, geojson := range getGeometries(t, "samples/multipoints/valids") {
		_, err := Decode([]byte(geojson))
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestDecodeGeodataInvalidMultiPoints(t *testing.T) {
	for _, geojson := range getGeometries(t, "samples/multipoints/invalids") {
		_, err := Decode([]byte(geojson))
		if err == nil {
			t.Fatalf("GeoJSON below is invalid...\n%s", geojson)
		}
	}
}
