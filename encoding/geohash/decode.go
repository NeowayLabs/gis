package geohash

import (
	"encoding/base32"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	// Geohash alphabet
	alphabet = "0123456789bcdefghjkmnpqrstuvwxyz"
)

// Decode a geohash into a longitude and latitude
func Decode(hash string) (float64, float64, error) {
	if len(hash)%8 != 0 {
		npad := 8 - (len(hash) % 8)
		for i := 0; i < npad; i++ {
			hash += "="
		}
	}
	return decode(hash)
}

func decode(hash string) (lng float64, lat float64, err error) {
	enc := base32.NewEncoding(alphabet)
	b32dec := base32.NewDecoder(enc, strings.NewReader(hash))
	var data []byte

	for {
		var tmp [3]byte
		n, err := b32dec.Read(tmp[:])
		if err != nil {
			if err == io.EOF {
				break
			}
			return lng, lat, err
		}

		data = append(data, tmp[:n]...)
	}

	fmt.Println("original: ", strconv.FormatUint(uint64(data[0]), 2))

	code := uint64(data[0]>>3) << 58
	code |= uint64((data[0]<<5)>>5) << 55

	fmt.Println(strconv.FormatUint(code, 2))
	fmt.Println(strconv.FormatUint(uint64(0x5555555555555555), 2))
	even := code & 0x5555555555555555
	fmt.Println(strconv.FormatUint(even, 2))
	return lat, lng, nil
}

func Encode(lon, lat float64, precision int) string {
	minLat, maxLat := -90.0, 90.0
	minLon, maxLon := -180.0, 180.0
	result := uint64(0)
	for i := 0; i < precision; i++ {
		if i%2 == 0 {
			midpoint := (minLon + maxLon) / 2
			if lon < midpoint {
				result <<= 1
				maxLon = midpoint
			} else {
				result = result<<1 | 1
				minLon = midpoint
			}
		} else {
			midpoint := (minLat + maxLat) / 2
			if lat < midpoint {
				result <<= 1
				maxLat = midpoint
			} else {
				result = result<<1 | 1
				minLat = midpoint
			}
		}
	}

	// DO NOT USE encoding/base32 ....
	var dst [32]byte
	enc := base32.NewEncoding(alphabet)
	src := []byte{
		byte(result >> 52),
		byte((result & 0xf) << 8),
	}
	fmt.Println("#")
	fmt.Println(strconv.FormatUint(uint64(result), 2))
	fmt.Printf("src0 %s\n", strconv.FormatUint(uint64(src[0]), 2))

	enc.Encode(dst[:], src)
	return string(dst[:])
}
