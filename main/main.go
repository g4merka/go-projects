package main

import (
	"fmt"
)

func main() {
	fmt.Println(ExpressionMatter(3, 5, 6))
}

func ExpressionMatter(a, b, c int) [3]int {
		return [3]int{a, b, c}
}