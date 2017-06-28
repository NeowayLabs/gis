package geojson

import "testing"

func testLineStringOK(t *testing.T, geojson string, expected LineString) {
	obj, err := Decode([]byte(geojson))
	if err != nil {
		t.Fatalf("Parsing: %s\nerror: %s", geojson, err)
	}

	assertLineString(t, obj, expected)
}

func TestLineStringOK(t *testing.T) {
	for _, tc := range []struct {
		geojson string
		line    LineString
	}{
		{
			geojson: `{
			"type": "LineString",
			"coordinates": [[10, 10], [12, 11]]
		}`,
			line: LineString{
				Coordinates: []Position{
					{Lat: 10, Lon: 10},
					{Lat: 11, Lon: 12},
				},
			},
		},
		{
			geojson: `{
        "type": "LineString",
        "coordinates": [
          [
            -78.3984375,
            69.53451763078358
          ],
          [
            55.54687499999999,
            70.25945200030641
          ],
          [
            54.84375,
            52.26815737376817
          ],
          [
            -74.8828125,
            29.53522956294847
          ],
          [
            -77.34374999999999,
            5.9657536710655235
          ],
          [
            50.9765625,
            11.867350911459308
          ],
          [
            68.90625,
            -9.44906182688142
          ],
          [
            -109.6875,
            -14.604847155053886
          ],
          [
            -94.5703125,
            45.82879925192134
          ],
          [
            32.34375,
            58.44773280389084
          ],
          [
            35.15625,
            64.62387720204691
          ],
          [
            -102.3046875,
            63.074865690586634
          ],
          [
            -80.85937499999999,
            69.03714171275197
          ]
        ]
      }`,
			line: LineString{
				Coordinates: []Position{
					{Lon: -78.3984375, Lat: 69.53451763078358},
					{Lon: 55.54687499999999, Lat: 70.25945200030641},
					{Lon: 54.84375, Lat: 52.26815737376817},
					{Lon: -74.8828125, Lat: 29.53522956294847},
					{Lon: -77.34374999999999, Lat: 5.9657536710655235},
					{Lon: 50.9765625, Lat: 11.867350911459308},
					{Lon: 68.90625, Lat: -9.44906182688142},
					{Lon: -109.6875, Lat: -14.604847155053886},
					{Lon: -94.5703125, Lat: 45.82879925192134},
					{Lon: 32.34375, Lat: 58.44773280389084},
					{Lon: 35.15625, Lat: 64.62387720204691},
					{Lon: -102.3046875, Lat: 63.074865690586634},
					{Lon: -80.85937499999999, Lat: 69.03714171275197},
				},
			},
		},
		{
			geojson: `{
			"type": "LineString",
			"coordinates": [
				[10, 10],
				[12, 11],
				[-1, 0],
				[-24, 12.2]
			]}`,
			line: LineString{
				Coordinates: []Position{
					{Lat: 10, Lon: 10},
					{Lat: 11, Lon: 12},
					{Lat: 0, Lon: -1},
					{Lat: 12.2, Lon: -24},
				},
			},
		},
	} {
		tc := tc
		testLineStringOK(t, tc.geojson, tc.line)
	}
}

func TestDecodeGeodataValidLineStrings(t *testing.T) {
	for _, geojson := range getGeometries(t, "samples/linestring/valids") {
		_, err := Decode([]byte(geojson))
		if err != nil {
			t.Fatalf("Parsing: %s\nerror: %s", geojson, err)
		}
	}
}

func TestDecodeGeodataInvalidLineStrings(t *testing.T) {
	for _, geojson := range getGeometries(t, "samples/linestring/invalids") {
		_, err := Decode([]byte(geojson))
		if err == nil {
			t.Fatalf("GeoJSON below is invalid...\n%s", geojson)
		}
	}
}
