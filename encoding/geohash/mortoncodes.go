package geohash

// interleave bits of x with zero bits at odd positions.
// ie. 0xff == 11111111 became 101010101010101
func split(x uint32) uint32 {
	x &= 0x0000ffff
	x = (x ^ (x << 8)) & 0x00ff00ff
	x = (x ^ (x << 4)) & 0x0f0f0f0f
	x = (x ^ (x << 2)) & 0x33333333
	x = (x ^ (x << 1)) & 0x55555555
	return x
}

func interleave(x, y uint32) uint32 {
	return split(y) | (split(x) << 1)
}
