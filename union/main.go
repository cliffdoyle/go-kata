package main

import (
	"fmt"
	"os"
)

// This function looks for unique characters in two strings then combines them
func uniqueCharacters(s1, s2 string) {
	// Initialize a map to keep track of characters that are unique
	seen := make(map[rune]bool)

	combined := s1 + s2
	result := ""
	for _, char := range combined {
		if !seen[char] {
			seen[char] = true
			result += string(char)
		}
	}
	fmt.Println(result)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}
	input1 := os.Args[1]
	input2 := os.Args[2]

	uniqueCharacters(input1, input2)
}
