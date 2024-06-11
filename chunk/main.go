package main

import "fmt"

func Chunk(slice []int, size int) {
	result := [][]int{}

	if size == 0 {
		return 
	}
	for firstIndex := 0; firstIndex < len(slice)-1; firstIndex += size {
		lastIndex := firstIndex + size
		if lastIndex > len(slice) {
			lastIndex = len(slice)
		}
		result = append(result, slice[firstIndex:lastIndex])
	}
	fmt.Println(result)
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
