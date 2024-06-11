package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
	}

	arg1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	result := 0
	for i := 2; i <= arg1; i++ {
		if isPrime(i) {
		result+=i
		}
	}
	fmt.Println(result)
}
