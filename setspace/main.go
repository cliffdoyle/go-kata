package main

import "fmt"

func SetSpace(s string) string {
	result := ""
	if s == "" {
		return ""
	}

	for i, char := range s {
		if char > '0' && char < '9' {
			return "Error"
		}
		if char==' '{
			return "Error"
		}

		if char >= 'A' && char <= 'Z' && i > 0 {
			result += " " + string(char)
		}else{
			result += string(char)
		}
			
		
	}
	fmt.Println(result)
	return result
}

func main() {
	fmt.Println(SetSpace("HelloWorld"))
	fmt.Println(SetSpace("HelloWorld12"))
	fmt.Println(SetSpace("Hello World"))
	fmt.Println(SetSpace(""))
	fmt.Println(SetSpace("LoremIpsumWord"))
}
