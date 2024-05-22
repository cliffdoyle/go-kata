package main

import "fmt"

func Ascii(str string) []byte {
	var result []byte
	if str==""{
		return nil
	}

	for i := 0; i < len(str); i++ {
		result = append(result, str[i])
	}
	return result
}

func main() {
	l := Ascii("Hello")
	fmt.Println(l)
	s:=Ascii("")
	fmt.Println(s)
}
