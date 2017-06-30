package geohash

import (
	"encoding/base32"
	"math"
)

const alphabet = "0123456789bcdefghjkmnpqrstuvwxyz"

func Encode(lon, lat float64, precision int) string {
	bits := uint(5 * precision)
	x, y := bisect(lon, 180), bisect(lat, 90)
	hashint := interleave(y, x)
	hashint >>= (64 - bits)
	src := []byte{
		byte(hashint >> 24),
		byte((hashint >> 16) & 0xFF),
		byte((hashint >> 8) & 0xFF),
		byte(hashint & 0xFF),
	}
	var dst [8]byte
	enc := base32.NewEncoding(alphabet)
	enc.Encode(dst[:], src)
	return string(dst[:])
}

func bisect(angle, r float64) uint32 {
	v := (angle + r) / 2 * r
	return uint32(v * math.Exp2(32))
}
