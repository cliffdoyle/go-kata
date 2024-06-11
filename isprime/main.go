package main

import "fmt"

func IsPrime(nb int) bool {
	if nb < 0 {
		return false
	}
	if nb == 1 {
		return false
	}

	if nb%2 != 0 && nb%nb == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(4))
	fmt.Println(IsPrime(3))
}
