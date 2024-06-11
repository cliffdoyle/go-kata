package main

import (
	//"github.com/01-edu/z01"
	"fmt"
)

func NRune(s string, n int) string {
	nrune := []rune(s)
	if n < 0 || n > len(nrune) {
		return "0"
	}
	return string(nrune[n])
}

func main() {
	fmt.Println(NRune("Hello!", 3))
	fmt.Println(NRune("salut!", 2))
	
}
