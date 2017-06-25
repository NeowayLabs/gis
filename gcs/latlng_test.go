package gcs

import "testing"

const (
	ε float64 = 0.000001 // TODO: increase
)

func floatEquals(a, b float64) bool {
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
		rad := radians(tc.degrees)
		if !floatEquals(float64(rad), float64(tc.expectedRadians)) {
			t.Fatalf("Expected %.8f radians but got %.12f",
				tc.expectedRadians, rad)
		}
	}
}
