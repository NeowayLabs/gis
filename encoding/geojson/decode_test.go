package geojson

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var (
	geodataPath string
)

func init() {
	geodataPath = os.Getenv("GOPATH") + "/src/github.com/NeowayLabs/gis/_testdata/geodata"
}

func getGeometries(t *testing.T, path string) []string {
	var geojsons []string

	path = geodataPath + "/wgs84/geojson/" + path
	files, err := ioutil.ReadDir(path)
	if err != nil {
		t.Fatal(err)
	}
	for _, fileInfo := range files {
		filename := path + "/" + fileInfo.Name()
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Fatal(err)
		}
		geojsons = append(geojsons, string(content))
	}
	return geojsons
}

func invalidType(t *testing.T, geojson string) {
	_, err := Decode([]byte(geojson))
	if err == nil {
		t.Fatalf("Geojson '%s' is invalid", geojson)
	}
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

func TestDecodeDecoder(t *testing.T) {
	d := NewDecoder(strings.NewReader(`{
		"type": "Point",
		"coordinates": [10, 10]
	}`))
	obj, err := d.Decode()
	if err != nil {
		t.Fatal(err)
	}
	if obj.Type() != "Point" {
		t.Fatalf("geometry type mismatch: got '%s' but expected 'Point'",
			obj.Type())
	}
}
