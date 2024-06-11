package main

import "fmt"

func LastRune(s string) rune {
	var lastrune rune

	for _, char := range s {
		lastrune = char
		

	}
	return lastrune
}

func main() {
	fmt.Println(string(LastRune("Hello")))
	fmt.Println(LastRune("Salut"))
	fmt.Println(LastRune("Ola"))
	fmt.Println()
}
