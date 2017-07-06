package geohash

const (
	base32   = "0123456789bcdefghjkmnpqrstuvwxyz"
	maxint32 = 1 << 32
)

var base32rev ['z' + 1]byte

func init() {
	for i := 0; i < len(base32); i++ {
		base32rev[int(base32[i])] = byte(i)
	}
}

func Encode(lon, lat float64, precision uint) string {
	bits := uint(5 * precision)
	alon, alat := adjust(lon, 180), adjust(lat, 90)
	code := encode(alon, alat, bits)
	return tobase32(code)[12-precision:]
}

func encode(lonInt, latInt uint32, bits uint) uint64 {
	code := interleave(latInt, lonInt)
	return code >> (64 - bits)
}

func adjust(angle, r float64) uint32 {
	p := (angle + r) / (2 * r)
	return uint32(p * maxint32)
}

func tobase32(code uint64) string {
	var b [12]byte

	// loop unrolling
	// Start by last position to skip bounds checking in next
	// assignments in b. Take a look in "Bounds Checking Elimination"
	// or BCE in Go >= 1.7

	b[11] = base32[code&0x1f]
	code >>= 5
	b[10] = base32[code&0x1f]
	code >>= 5
	b[9] = base32[code&0x1f]
	code >>= 5
	b[8] = base32[code&0x1f]
	code >>= 5
	b[7] = base32[code&0x1f]
	code >>= 5
	b[6] = base32[code&0x1f]
	code >>= 5
	b[5] = base32[code&0x1f]
	code >>= 5
	b[4] = base32[code&0x1f]
	code >>= 5
	b[3] = base32[code&0x1f]
	code >>= 5
	b[2] = base32[code&0x1f]
	code >>= 5
	b[1] = base32[code&0x1f]
	code >>= 5
	b[0] = base32[code&0x1f]
	return string(b[:])
}
