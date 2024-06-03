package main

import "fmt"

func FirstRune(s string) rune {
	var result rune
	for _, char := range s {
		result = char
		break
	}

	return result
}

func NRune(s string, n int) rune {
	if n > len(s) {
		fmt.Println("Out of range")
	}
	var result rune
	for i, char := range s {
		if i == n {
			result = char
		}
	}

	return result
}

func LastRune(s string) rune {
	var result rune
	for _, char := range s {
		result = char
	}

	return result
}

func main() {
	fmt.Println(string(FirstRune("hello world")))
	fmt.Println(string(FirstRune("gello")))
	fmt.Println(string(NRune("gello", 7)))
	fmt.Println(string(LastRune("Hello you")))
}
