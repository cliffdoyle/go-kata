package main

import "fmt"

func RepeatAlpha(s string) string {
	res := ""

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			res += Repeat(char, int(char-'a'+1))
		} else if char >= 'A' && char <= 'Z' {
			res += Repeat(char, int(char-'A'+1))
		} else {
			res += string(char)
		}

	}
	fmt.Println(res)

	return res
}

func Repeat(r rune, n int) string {
	res := ""

	for i := 0; i < n; i++ {
		res += string(r)
	}

	return res
}

func main(){
	RepeatAlpha("chuomi.")
}
