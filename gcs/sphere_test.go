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
