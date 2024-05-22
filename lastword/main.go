package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	args1 := args[0]

	word := ""
	for i := len(args1) - 1; i > 0; i-- {
		// if word == "" && string(args1[i]) == " " {
		// 	continue
		// }
		if word != "" && string(args1[i]) == " " {
			break
		}
		word = string(args1[i]) + word
		// fmt.Print(string(args1[i]))
	}

	fmt.Println(word)
}
