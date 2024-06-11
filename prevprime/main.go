package main

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}



func main() {
	fmt.Println(isPrime(4))
	fmt.Println(isPrime(7))
}
