// https://www.codewars.com/kata/55d24f55d7dd296eb9000030

package main

import (
	"fmt"
)

func main() {
	fmt.Println(Summation(8))
}

func Summation(n int) int {
	if n == 1 {
		return 1
	} 
	return n + Summation(n-1)
}