package geohash

// stripe bits of val by adding zero bits between every bit.
func stripe(val uint32) uint64 {
	X := uint64(val)
	X = (X | (X << 16)) & 0x0000ffff0000ffff
	X = (X | (X << 8)) & 0x00ff00ff00ff00ff
	X = (X | (X << 4)) & 0x0f0f0f0f0f0f0f0f
	X = (X | (X << 2)) & 0x3333333333333333
	X = (X | (X << 1)) & 0x5555555555555555
	return X
}

// interleave x's bits with y's. Bits of x occupy even positions.
func interleave(x, y uint32) uint64 {
	return stripe(x) | (stripe(y) << 1)
}
