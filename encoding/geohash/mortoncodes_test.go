package geohash

import "testing"

func TestMortonSplits(t *testing.T) {
	for _, tc := range []struct {
		x, expected uint32
	}{
		{x: 0, expected: 0},               // 0b000000000000000000000000000000
		{x: 1, expected: 1},               // 0b000000000000000000000000000001
		{x: 2, expected: 0x4},             // 0b000000000000000000000000000100
		{x: 3, expected: 0x5},             // 0b000000000000000000000000000101
		{x: 4, expected: 0x10},            // 0b000000000000000000000000010000
		{x: 0xff, expected: 0x5555},       // 0b000000000000000101010101010101
		{x: 0xffff, expected: 0x55555555}, // 0b101010101010101101010101010101
	} {
		if got := split(tc.x); got != tc.expected {
			t.Fatalf("Got 0b%0.4b but expected 0b%0.4b", got, tc.expected)
		}
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		split(uint32(i))
	}
}

func BenchmarkInterleaves(b *testing.B) {
	for i := 0; i < b.N; i++ {
		interleave(uint32(i), uint32(i+10))
	}
}

func TestMortonInterleaves(t *testing.T) {
	for _, tc := range []struct {
		x, y, expected uint32
	}{
		{x: 0, y: 0, expected: 0},
		{x: 0xff, y: 0xff, expected: 0xffff},
		{x: 0xffff, y: 0xffff, expected: 0xffffffff},
	} {
		if got := interleave(tc.x, tc.y); got != tc.expected {
			t.Fatalf("Got %0.8b but expects %0.8b", got, tc.expected)
		}
	}
}

func TestMortonGenerateTables(t *testing.T) {
	for i := 0; i < 256; i++ {

	}
}
