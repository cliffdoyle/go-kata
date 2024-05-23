package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	count := len(args)
	fmt.Println(count)
}
