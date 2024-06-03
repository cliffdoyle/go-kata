package main

import "fmt"

func ReduceInt(a []int, f func(int, int) int) {
	init := a[0]

	for i := 1; i < len(a); i++ {
		init = f(init, a[i])
	}

	fmt.Println(init)
}

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}

	div := func(acc int, cur int) int {
		return acc / cur
	}

	modulo := func(acc int, cur int) int {
		return acc % cur
	}
	
	add := func(acc int, cur int) int {
		return acc + cur
	}
	

	slice := []int{900,500,287}

	ReduceInt(slice, modulo)
	ReduceInt(slice, div)
	ReduceInt(slice, mul)
	ReduceInt(slice, add)
	
	
}
