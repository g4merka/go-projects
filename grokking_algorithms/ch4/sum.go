package main

import (
	"fmt"
)

func main() {
	fmt.Println(Sum([]int{0, 7, 10}))
}

func Sum(arr []int)  int {
	if len(arr) == 1 {
		return arr[0]
		 }
	n := len(arr)-1
	return arr[n] + Sum(arr[:n])
}