// https://www.codewars.com/kata/55ecd718f46fba02e5000029

package main

import (
	"fmt"
)

func main() {
	fmt.Println(Between(90, 100))
}

func Between(a, b int) []int {
	var arr []int
	for a <= b {
		arr = append(arr, a)
		a++
	}
	return arr
}