package main

import (
	"fmt"
)

func XOR32Hash(s string) string {
	var h uint32 = 0

	for i, c := range s {
		shift := uint((i % 4) * 8)
		h ^= uint32(c) << shift
	}

	return fmt.Sprintf("%8x", h)
}

func main() {
	input := "bitcoin0"
	target := XOR32Hash(input)

	collision := ""

	for mask := 1; mask < 16; mask++ {
		b := []byte(input)

		for i := range 4 {
			if mask&(1<<i) != 0 {
				b[i], b[i+4] = b[i+4], b[i]
			}
		}

		found := string(b)
		if XOR32Hash(found) == target {
			collision = found
			break
		}
	}

	fmt.Println(collision)
}

