package main

import (
	"fmt"
)

func main() {
	fmt.Println(quickSort([]int{8, 3}))
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		pivot := arr[0]
		var less, greater []int
		for _, value := range arr[1:] {
			if value <= pivot {
				less = append(less, value)
			} else {
				greater = append(greater, value)
			}
		}
		res := append(append(quickSort(less), pivot), quickSort(greater)...)
		return res
	}
}