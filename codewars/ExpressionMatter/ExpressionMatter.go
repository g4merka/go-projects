/*
https://www.codewars.com/kata/5ae62fcf252e66d44d00008e
*/
package main

import "fmt"

func main() {
	fmt.Println(ExpressionMatter(2, 4, 3))
}

func ExpressionMatter(a, b, c int) int {
	if a != 1 && b != 1 && c != 1 {
		return a * b * c
	} else if a == 1 && c == 1 {
		return a + b + c
	} else {
		if a <= c {
			return (a + b) * c
		}
		return a * (b + c)
	}
}