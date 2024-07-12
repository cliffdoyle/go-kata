package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	num, _ := strconv.Atoi(os.Args[1])

	hexValues := "0123456789abcdef"
	result := ""
	for num > 0 {
		digit := num%16
		result = string(hexValues[digit]) + result
		num /= 16

	}
	fmt.Println(result)
}