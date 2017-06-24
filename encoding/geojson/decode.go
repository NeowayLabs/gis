package geojson

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

func Decode(data []byte) (Object, error) {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	objmap := map[string]*json.RawMessage{}
	err := decoder.Decode(&objmap)
	if err != nil {
		return nil, err
	}
	return decode(objmap)
}

func decode(objmap map[string]*json.RawMessage) (Object, error) {
	typ, ok := objmap["type"]
	if !ok || typ == nil {
		return nil, errors.New("missing required 'type' member")
	}

	var typstr string
	err := json.Unmarshal(*typ, &typstr)
	if err != nil {
		return nil, err
	}

	if !geojsonTypes.Valid(typstr) {
		return nil, fmt.Errorf("invalid geojson object: %s", typstr)
	}

	switch typstr {
	case "Point":
		return decodePoint(objmap)
	}

	return nil, nil
}

func decodePoint(objmap map[string]*json.RawMessage) (Object, error) {
	coordinates, ok := objmap["coordinates"]
	if !ok {
		return nil, errors.New("missing required 'coordinates' member")
	}

	var coords []json.Number
	err := json.Unmarshal(*coordinates, &coords)
	if err != nil {
		return nil, err
	}
	if len(coords) < 2 ||
		len(coords) > 3 {
		return nil, fmt.Errorf("invalid coordinate for point: %#v", coords)
	}

	lon, err := coords[0].Float64()
	if err != nil {
		return nil, fmt.Errorf("invalid longitude: %s", coords[0])
	}
	lat, err := coords[1].Float64()
	if err != nil {
		return nil, fmt.Errorf("invalid latitude: %s", coords[1])
	}

	var alt float64

	if len(coords) == 3 {
		alt, err = coords[2].Float64()
		if err != nil {
			return nil, fmt.Errorf("invalid altitude: %s", coords[2])
		}
	}

	pos := Position{
		Lon: lon,
		Lat: lat,
		Alt: alt,
	}

	return Point{
		Coordinate: pos,
	}, nil
}
