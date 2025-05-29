package main

import (
	"fmt"
)

func ExpressionMatter(a int, b int, c int) int {
	if a != 1 && b != 1 && c != 1 {
		return a * b * c
	}
	// return 
}

func main() {
	fmt.Println(ExpressionMatter(1, 2, 3))
}
