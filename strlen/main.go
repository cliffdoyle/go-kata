package main

import (
	//"fmt"

	"github.com/01-edu/z01"
)

func StrLen(s string) int {
	length := len(s)

	return length
}

func main() {
	l := StrLen("Hello World!")

	var result []rune

	if l == 0 {
		z01.PrintRune(0)
	} else {
		for l > 0 {
			digit := l % 10
			result = append([]rune{rune('0' + digit)}, result...)
			l /= 10

		}
	}
	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
