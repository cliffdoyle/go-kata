package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1]
	result := ""

	// result := ""

	// for i, char := range s {
	// 	if char >= 'a' && char <= 'z' && i > 0 {
	// 		char = ('a'+25)%26 + 'a'
	// 		result += string(char)

	// 	} else {
	// 		if char > 'A' && char < 'Z' && i > 0 {
	// 			char = ('A'+25)%26 + 'A'
	// 			result += string(char)
	// 		}
	// 	}
	// }

	// fmt.Println(result)
	if len(os.Args) != 2 {
		// fmt.Println()
		return
	}
	for _, char := range args {
		if char >= 'a' && char <= 'z' {
			char = 'z' - char + 'a'
			 result += string(char)

		} else if char >= 'A' && char <= 'Z' {
				char = 'Z' - char + 'A'
				result += string(char)
			}else{
				result+=string(char)
			}
		
		}
		fmt.Println(result)
	}
	

