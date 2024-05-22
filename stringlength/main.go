package main

import "fmt"

func Strln(s string) int {
	return len(s)
}

func main() {
	l := Strln("Hello World!")
	fmt.Println(l)
}
