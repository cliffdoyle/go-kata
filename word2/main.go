package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	input1 := os.Args[1]
	input2 := os.Args[2]
	found := false
	result := ""
	for _, cha := range input1 {
		for i, char := range input2 {
			if cha == char {
				result += string(char)
				found = true
				input2 = input2[i+1:]
				break
			}
		}
	}
	if found {
		fmt.Println(result)
	}
}
