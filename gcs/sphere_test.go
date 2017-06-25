package gcs

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

type (
	ttSphereArea struct {
		r        float64 // radius
		expected float64
	}

	ttSphereContains struct {
		point    Point
		sphere   Sphere
		expected bool
	}
)

func testSphereArea(t *testing.T, testcase ttSphereArea) {
	sp := Sphere{
		R: testcase.r,
	}
	area := sp.Area()
	if area != testcase.expected {
		t.Fatalf("Sphere area differs: %f != %f", area, testcase.expected)
	}
}

func testSphereContains(t *testing.T, testcase ttSphereContains) {
	got := testcase.sphere.Contains(testcase.point)
	if got != testcase.expected {
		t.Fatalf("expected %v but got %v", got, testcase.expected)
	}
}

func TestSphereArea(t *testing.T) {
	for _, tc := range []ttSphereArea{
		{
			r:        1.0,
			expected: 4 * math.Pi,
		},
		{
			r:        2.0,
			expected: 16 * math.Pi,
		},
		{
			r:        0.5,
			expected: math.Pi,
		},
	} {
		tc := tc
		t.Run(strconv.FormatFloat(tc.r, 'e', -1, 64), func(t *testing.T) {
			testSphereArea(t, tc)
		})
	}
}

func TestSphereContains(t *testing.T) {
	for _, tc := range []ttSphereContains{
		{
			sphere:   Sphere{R: 1.0},
			point:    Point{R: 1.0, λ: 1, φ: 1.5},
			expected: true,
		},
	} {
		tc := tc
		t.Run(fmt.Sprintf("Point %#v inside sphere %#v", tc.point, tc.sphere),
			func(t *testing.T) {
				testSphereContains(t, tc)
			})

	}
}

func TestSphereEarthDistance(t *testing.T) {
	for _, tc := range []struct {
		srcName  string
		dstName  string
		src      SPoint
		dst      SPoint
		expected float64
	}{
		{
			srcName:  "Florianópolis",
			dstName:  "São Paulo",
			src:      NewSPoint(-48.5482, -27.5949),
			dst:      NewSPoint(-46.6333, -23.5505),
			expected: 489532.64,
		},
		{
			srcName:  "Origin",
			dstName:  "Antipodal -180.0",
			src:      NewSPoint(0, 0),
			dst:      NewSPoint(-180, 0),
			expected: math.Pi * SphericalEarth.R, // 2πr/2
		},
		{
			srcName:  "Origin",
			dstName:  "Antipodal +180.0",
			src:      NewSPoint(0, 0),
			dst:      NewSPoint(180, 0),
			expected: math.Pi * SphericalEarth.R, // 2πr/2
		},
		{
			srcName:  "antipodal -180.0",
			dstName:  "antipodal +180.0",
			src:      NewSPoint(-180.0, 0),
			dst:      NewSPoint(180.0, 0),
			expected: 0, // -180.0 and +180.0 is the same longitude
		},
		{
			srcName: "arbitrary p1",
			dstName: "arbitrary p2",
			src:     NewSPoint(-100.0, 32.0),
			dst:     NewSPoint(40.0, -27.0),
			// online calculated at http://www.movable-type.co.uk/scripts/latlong.html
			expected: 16144146.90,
		},
	} {
		tc := tc
		desc := fmt.Sprintf("Distance from %s to %s",
			tc.srcName, tc.dstName)
		t.Run(desc, func(t *testing.T) {
			d := SphericalEarth.Distance(tc.src, tc.dst)
			if !floatEquals(d, tc.expected, 0.01) {
				t.Fatalf("Distance off by %.12f. Expected %.12f but got %.12f",
					math.Abs(d-tc.expected), tc.expected, d)
			}
		})
	}

}

func TestSphereMarsDistance(t *testing.T) {
	for _, tc := range []struct {
		srcName  string
		dstName  string
		src      SPoint
		dst      SPoint
		expected float64
	}{
		{
			srcName: "Cydonia mensae",
			dstName: "Tharsis Montes",
			src:     NewSPoint(12.80, 37.00),
			dst:     NewSPoint(113.30, 2.80),
			// Distance calculated here, but not reliable either:
			// http://answers.google.com/answers/threadview/id/443926.html
			expected: 5719015.85,
		},
	} {
		tc := tc
		desc := fmt.Sprintf("Distance from %s to %s",
			tc.srcName, tc.dstName)
		t.Run(desc, func(t *testing.T) {
			d := SphericalMars.Distance(tc.src, tc.dst)
			if !floatEquals(d, tc.expected, 0.01) {
				t.Fatalf("Distance off by %.12f. Expected %.12f but got %.12f",
					math.Abs(d-tc.expected), tc.expected, d)
			}
		})
	}

}
