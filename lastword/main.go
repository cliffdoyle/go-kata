package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	args1 := args[0]

	var word string
	for i := len(args1) - 1; i > 0; i-- {
		if word == "" && string(args1[i]) == " " {
			continue
		}
		if word != "" && string(args1[i]) == " " {
			break
		}
		word = string((args1[i])) + word
		// fmt.Print(string(args1[i]))
	}

	for _, ch := range word {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
