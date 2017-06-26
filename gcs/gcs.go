package gcs

type GCS interface {
	Radius() float64
	Distance(p1, p2 Point) float64
}
