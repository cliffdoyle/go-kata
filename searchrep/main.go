package main

import (
	"fmt"
	"os"
)

func main() {
	result := ""
	args := os.Args[1:]

	if len(os.Args) != 4 {
		return
	}
	input1 := args[0]
	input2str := args[1]
	input3str := args[2]

	input2 := rune(input2str[0])
	input3 := rune(input3str[0])

	for _, char1 := range input1 {
		if char1 == input2 {
			result += string(input3)
		} else {
			result += string(char1)
		}
	}
	fmt.Println(result)
}
