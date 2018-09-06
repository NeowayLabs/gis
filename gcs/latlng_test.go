package gcs

import "testing"

func floatEquals(a, b, ε float64) bool {
	if (a-b) < ε && (b-a) < ε {
		return true
	}
	return false
}

func TestGCSAngleDegree(t *testing.T) {
	for _, tc := range []struct {
		degrees         float64
		expectedRadians Angle
	}{
		{
			degrees:         0,
			expectedRadians: 0,
		},
		{
			degrees:         90,
			expectedRadians: 1.570796,
		},
	} {
		ε := 0.000001 // error factor
		rad := toRadians(tc.degrees)
		if !floatEquals(float64(rad), float64(tc.expectedRadians), ε) {
			t.Fatalf("Expected %.8f radians but got %.12f",
				tc.expectedRadians, rad)
		}
	}
}
