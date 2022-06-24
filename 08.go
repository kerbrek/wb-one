package main

import "fmt"

// Дана переменная int64. Разработать программу которая устанавливает
// i-й бит в 1 или 0.

// Two's complement and fmt.Printf
// https://stackoverflow.com/questions/37582550/twos-complement-and-fmt-printf

func setBitAt(position uint, bit bool, num int64) int64 {
	if bit {
		// mask := (1 << position)
		// fmt.Printf("mask: %b\n", uint64(mask))
		return num | (1 << position)
	}
	// mask := ^(1 << position)
	// fmt.Printf("mask: %b\n", uint64(mask))
	return num & ^(1 << position)
}

func btou(b bool) uint {
	if b {
		return 1
	}
	return 0
}

func main() {
	x := int64(42)

	pos, bit := uint(4), true
	fmt.Printf("set bit at position %d to %d\n", pos, btou(bit))
	y := setBitAt(pos, bit, x)
	fmt.Printf("before: %d -- %b\n", x, x)
	fmt.Printf("after : %d -- %b\n", y, y)
	fmt.Printf("\n")

	pos, bit = uint(3), false
	fmt.Printf("set bit at position %d to %d\n", pos, btou(bit))
	z := setBitAt(pos, bit, x)
	fmt.Printf("before: %d -- %b\n", x, x)
	fmt.Printf("after : %d -- %b\n", z, z)
	fmt.Printf("\n")

	fmt.Printf("\n")

	a := int64(-42)

	pos, bit = uint(3), true
	fmt.Printf("set bit at position %d to %d\n", pos, btou(bit))
	b := setBitAt(pos, bit, a)
	fmt.Printf("before: %d -- %b\n", a, uint64(a))
	fmt.Printf("after : %d -- %b\n", b, uint64(b))
	fmt.Printf("\n")

	pos, bit = uint(2), false
	fmt.Printf("set bit at position %d to %d\n", pos, btou(bit))
	c := setBitAt(pos, bit, a)
	fmt.Printf("before: %d -- %b\n", a, uint64(a))
	fmt.Printf("after : %d -- %b\n", c, uint64(c))
}
