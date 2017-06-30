package geohash

import (
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	_, _, err := Decode("ezs42")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(Encode(49, 27, 60))
}
