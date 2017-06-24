package geojson

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type Decoder struct {
	content io.Reader
	Strict  bool
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		content: r,
		Strict:  true,
	}
}

func Decode(data []byte) (Object, error) {
	d := NewDecoder(bytes.NewReader(data))
	return d.Decode()
}

func (d *Decoder) Decode() (Object, error) {
	jsonDecoder := json.NewDecoder(d.content)
	jsonDecoder.UseNumber()
	objmap := map[string]*json.RawMessage{}
	err := jsonDecoder.Decode(&objmap)
	if err != nil {
		return nil, err
	}
	return d.decode(objmap)
}

func (d *Decoder) decode(objmap map[string]*json.RawMessage) (Object, error) {
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
		return d.decodePoint(objmap)
	}

	return nil, nil
}

func (d *Decoder) decodePoint(objmap map[string]*json.RawMessage) (Object, error) {
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

	if d.Strict {
		if lon < -180.0 || lon > 180.0 {
			return nil, fmt.Errorf("longitude must satisfy: -180.0 < lon < 180.0")
		}
		if lat < -90.0 || lat > 90.0 {
			return nil, fmt.Errorf("latitude must satisfy: -90.0 < lat < 90.0")
		}
	}

	pos := Position{
		Lon: lon,
		Lat: lat,
		Alt: alt,
	}

	return Point{
		Coordinates: pos,
	}, nil
}
