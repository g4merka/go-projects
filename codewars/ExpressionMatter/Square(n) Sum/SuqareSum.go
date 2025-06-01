/*
https://www.codewars.com/kata/515e271a311df0350d00000f
*/

package main

import "fmt"

func main() {
	nums := []int{0}
	fmt.Println(SquareSum(nums))
}

func SquareSum(numbers []int) int {
	var sum int
	for _, value := range numbers {
		value *= value
		sum += value
	}
	return sum
}

/* Best:
func SquareSum(nums []int) (res int) {
    for _, val := range nums {
      res += val * val
    }
    return res
}
*/