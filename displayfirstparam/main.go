package main

import (
	"fmt"

	"os"
)


func main(){
	args := os.Args[1:]

	var firstParam string

	for _, char := range args{
		firstParam=char
		break
		
	}

	fmt.Println(firstParam)
}

