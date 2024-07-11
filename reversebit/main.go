package main

import "fmt"

func RevBits(n uint) uint {
	var res uint

	for i := 0; i < 8; i++ {
		res |= (n >> i & 1) << (7 - i)
	}
	return res
}

func main() {
	ret := RevBits(19)

	fmt.Printf("%08b\n", ret)
}
