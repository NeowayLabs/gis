package geojson

import "testing"

func assertPosition(t *testing.T, pos1, pos2 Position) {
	if pos1.Lat != pos2.Lat ||
		pos1.Lon != pos2.Lon ||
		pos1.Alt != pos2.Alt {
		t.Fatalf("Position mismatch: (%#v) != (%#v)", pos1, pos2)
	}
}

func assertObject(t *testing.T, obj1, obj2 Object) {
	if obj1.Type() != obj2.Type() {
		t.Fatalf("types mismatch: %s != %s", obj1.Type(), obj2.Type())
	}
}

func assertPoint(t *testing.T, p1, p2 Object) {
	assertObject(t, p1, p2)
	point1, ok := p1.(Point)
	point2, ok2 := p2.(Point)
	if !ok {
		t.Fatalf("Object (%#v) is not a point", p1)
	}
	if !ok2 {
		t.Fatalf("Object (%#v) is not a point", p2)
	}

	coord1, coord2 := point1.Coordinates, point2.Coordinates
	assertPosition(t, coord1, coord2)
}
