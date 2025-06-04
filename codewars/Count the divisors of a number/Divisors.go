// https://www.codewars.com/kata/542c0f198e077084c0000c2e

package main

import (
	"fmt"
)

func main() {
	fmt.Println(Divisors(10))
}

func Divisors(n int) int {
	var res int
	for i := 1; i <= n; i++ {
		if n % i == 0 {
			res++
			continue
		}
	}
	return res
}