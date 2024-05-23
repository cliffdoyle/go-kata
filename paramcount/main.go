package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	count := len(args)

	var result []rune

	if count == 0 {
		return
	} else {
		for count > 0 {
			digit := count % 10
			result = append([]rune{rune('0' + digit)}, result...)
			count /= 10
		}
	}

	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
