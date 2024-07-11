package main

import "fmt"

func SwapBits(n uint, i, j uint) uint {
	bit_i := (n >> i) & 1
	bit_j := (n >> j) & 1

	if bit_i != bit_j {
		n ^= (1 << i) | (1 << j)
	}
	return n

}

func Rev(n uint) uint {
	var res uint

	for i := 0; i < 8; i++ {
		res |= (n >> i & 1) << (7 - i)
	}
	return res
}

func main() {
	i := uint(1)
	j := uint(4)
	n := uint(38)

	swapped := SwapBits(n, i, j)
	fmt.Printf("OriginalBits: %08b\n", n)
	reversed := Rev(38)
	fmt.Printf("Reversed Bits: %08b\n", reversed)
	fmt.Printf("SwapBits: %08b\n", swapped)
}
