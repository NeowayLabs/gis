package geohash

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

const csvfile = "/src/github.com/NeowayLabs/gis/encoding/geohash/_testdata/geohash.csv"

type testcase struct {
	lon, lat     float64
	expectedHash string
	precision    uint
}

func RandomPoint() (float64, float64) {
	lat := -90 + 180*rand.Float64()
	lon := -180 + 360*rand.Float64()
	return lon, lat
}

func TestEncode(t *testing.T) {
	for _, tc := range []testcase{
		{
			expectedHash: "d9ez4jg6pmp4",
			lon:          -62.133888, lat: 9.699925,
			precision: 12,
		},
		{
			expectedHash: "d",
			lon:          -62.133888, lat: 9.699925,
			precision: 1,
		},
		{
			expectedHash: "d9",
			lon:          -62.133888, lat: 9.699925,
			precision: 2,
		},
		{
			expectedHash: "d9e",
			lon:          -62.133888, lat: 9.699925,
			precision: 3,
		},
		{
			expectedHash: "6gj7p2820", // florianópolis
			lon:          -48.5482,
			lat:          -27.5949,
			precision:    9,
		},
		{
			expectedHash: "6gyf4bf8m", // são paulo
			lon:          -46.6333,
			lat:          -23.5505,
			precision:    9,
		},
	} {
		got := Encode(tc.lon, tc.lat, tc.precision)
		if tc.expectedHash != got {
			t.Fatalf("mismatch: %s != %s", tc.expectedHash, got)
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lon, lat := RandomPoint()
		Encode(lon, lat, 12)
	}
}

func testEncode(t *testing.T, lon, lat float64, expected string) {
	got := Encode(lon, lat, 12)
	if !strings.HasPrefix(got, expected) {
		t.Fatalf("mismatch: %s != %s", got, expected)
	}
}

func TestEncodeFromCSV(t *testing.T) {
	fullpath := os.Getenv("GOPATH") + csvfile
	content, err := ioutil.ReadFile(fullpath)
	if err != nil {
		t.Fatal(err)
	}

	type testcase struct {
		lon, lat float64
		hash     string
	}

	tcases := []testcase{}

	start := 0
	for i := range content {
		if content[i] != '\n' {
			continue
		}
		line := string(content[start:i])
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			t.Fatal("Invalid line in geohash.csv")
		}
		lonstr := parts[0]
		latstr := parts[1]
		geohash := parts[2]
		lon, err := strconv.ParseFloat(lonstr, 64)
		if err != nil {
			t.Fatal(err)
		}
		lat, err := strconv.ParseFloat(latstr, 64)
		if err != nil {
			t.Fatal(err)
		}
		tcases = append(tcases, testcase{
			lon:  lon,
			lat:  lat,
			hash: geohash,
		})
		start = i + 1
	}

	startTime := time.Now()

	for _, tc := range tcases {
		testEncode(t, tc.lon, tc.lat, tc.hash)
	}

	elapsed := time.Since(startTime)
	t.Logf("%d geohashes processed in %s", len(tcases), elapsed)
}
