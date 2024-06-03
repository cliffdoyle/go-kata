package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	result := ""
	str := os.Args[1]

	for _, char := range str {

		if char >= 'a' && char <= 'z' {
			char = (char-'a'+13)%26 + 'a'
		}

		if char >= 'A' && char <= 'Z' {
			char = (char-'A'+13)%26 + 'A'
		}
		result += string(char)

	}
	fmt.Println(result)
}
