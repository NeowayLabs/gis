package geohash

import "testing"

func TestMortonStripes(t *testing.T) {
	for _, tc := range []struct {
		x        uint32
		expected uint64
	}{
		{x: 0, expected: 0},               // 0b000000000000000000000000000000
		{x: 1, expected: 1},               // 0b000000000000000000000000000001
		{x: 2, expected: 0x4},             // 0b000000000000000000000000000100
		{x: 3, expected: 0x5},             // 0b000000000000000000000000000101
		{x: 4, expected: 0x10},            // 0b000000000000000000000000010000
		{x: 0xff, expected: 0x5555},       // 0b000000000000000101010101010101
		{x: 0xffff, expected: 0x55555555}, // 0b101010101010101101010101010101
	} {
		if got := stripe(tc.x); got != tc.expected {
			t.Fatalf("Got 0b%0.4b but expected 0b%0.4b", got, tc.expected)
		}
	}

	expected := uint64(0xffffffff)
	got := stripe(0xffff) | (stripe(0xffff) << 1)
	if got != expected {
		t.Fatalf("Expected %0.8b but got %0.8b", expected, got)
	}
}

func BenchmarkStripe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stripe(uint32(i))
	}
}

func BenchmarkInterleaves(b *testing.B) {
	for i := 0; i < b.N; i++ {
		interleave(uint32(i), uint32(i+10))
	}
}

func TestMortonInterleaves(t *testing.T) {
	for _, tc := range []struct {
		x, y     uint32
		expected uint64
	}{
		{x: 0, y: 0, expected: 0},
		{x: 0xff, y: 0xff, expected: 0xffff},
		{x: 0xffff, y: 0xffff, expected: 0xffffffff},

		{x: 0, y: 1, expected: 2},
		{x: 0, y: 2, expected: 8},
		{x: 0, y: 3, expected: 10},
		{x: 0, y: 4, expected: 32},
		{x: 0, y: 5, expected: 34},
		{x: 0, y: 6, expected: 40},

		{x: 1, y: 0, expected: 1},
		{x: 1, y: 1, expected: 3},
		{x: 1, y: 2, expected: 9},
		{x: 0x55, y: 0x55, expected: 0x3333},
		{x: 0x5555, y: 0x5555, expected: 0x33333333},
		{x: 0x555555, y: 0x555555, expected: 0x333333333333},
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
