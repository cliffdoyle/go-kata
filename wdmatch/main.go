package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	result := ""

	input1 := os.Args[1]
	input2 := os.Args[2]
//iterate over each character in input1
	for _, char1 := range input1 {
		found := false
		//Iterate over each character in input2 starting from  the current index
		for i, char2 := range input2 {
			if char1 == char2 {
				result += string(char2)
				//update input2 to start from the next character after the match
				input2 = input2[i+1:]
				found = true
				break
			}
		}
		//if any character is not found in sequence in input2, break
		if !found {
			result = ""
			break
		}
	}
	if result == input1 {
		fmt.Println(result)
	
}
}
