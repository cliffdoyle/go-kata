package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	str1 := args[0]
	str2 := args[1]

	if len(args) != 2 {
		return
	}

	result := " "

	for _, char1 := range str1 {
		found := false
		for i, char2 := range str2 {
			if char1 == char2 {
				result += string(char2)
				str2 = str2[i+1:]
				found = true
				break

			}
		}
		if !found{
			return
		}
	}
	//if result==str1{
		fmt.Println(result)
	//}
}
