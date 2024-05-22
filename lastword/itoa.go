package main

import (
	"github.com/01-edu/z01"
)

func Itoa(n int) string {
	if n == 0 {
		z01.PrintRune('0')
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	var result []rune
	for n > 0 {
		digit := n % 10
		result = append([]rune{rune('0' + digit)}, result...)
		n /= 10
	}
	return string(result)
}
