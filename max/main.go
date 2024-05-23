package main

import "fmt"

func Max(a []int) int {
	if a == nil {
		return 0
	}

	max := 0

	for _, num := range a {
		if num>max {
			max=num
		}
	}
	return max
}

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	b:=[]int{}
	max := Max(a)
	fmt.Println(max)
	fmt.Println(b)
}
