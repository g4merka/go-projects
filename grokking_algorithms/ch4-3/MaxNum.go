package main

import (
	"fmt"
)

func Max(arr []int) int {
	biggest := arr[0]

	for _, value := range arr {
		if value > biggest {
			biggest = value
		}
	}
	return biggest
}

func main() {
	fmt.Println(Max([]int{7, 10, 9}))
}