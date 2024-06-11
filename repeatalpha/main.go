package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	input := os.Args[1]
	result := ""
	for _, char := range input {
		if (char >= 'a' && char <= 'z')|| (char >='A' && char <='Z'){
		repeatCount := 0
		if char >= 'a' && char <= 'z' {
			repeatCount = int(char - 'a' + 1)
		} else if char >= 'A' && char <= 'Z' {
			repeatCount = int(char - 'A' + 1)
		}
			for i := 0; i < repeatCount; i++ {
				result += string(char)
			}

		} else {
			result += string(char)
		}

	}
	fmt.Println(result)
}
