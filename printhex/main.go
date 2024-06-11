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

	agr1, err := strconv.Atoi(os.Args[1])
	if err != nil{
		fmt.Println("ERROR")
		return
	}

	hexNums := "0123456789abcdef"
	result := ""

	for agr1 > 0 {
		rem := agr1 % 16
		result = string(hexNums[rem]) + result
		agr1 /= 16

	}
	fmt.Println(result)
}
