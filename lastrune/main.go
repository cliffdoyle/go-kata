package main

import "fmt"


func LastRune(s string)rune{
	firstRune := []rune(s)

	return firstRune[len(s)-1]

	
}

func main() {
	fmt.Println(LastRune("Hello!"))
	fmt.Println(LastRune("Salut!"))
	fmt.Println(LastRune("Ola!"))
	fmt.Println()
}
