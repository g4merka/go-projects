// https://www.codewars.com/kata/57f781872e3d8ca2a000007e

package main

import (
	"fmt"
)

func main() {
	fmt.Println(Maps([]int{1,2,3}))
}

func Maps(x []int) []int {
	var arr []int
	for _, v := range x {
		arr = append(arr, v*2)
	}
	return arr
}