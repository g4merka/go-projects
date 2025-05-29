package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 6, 2}
	fmt.Println(findSmallest(nums))
}
func findSmallest(arr []int) []int {
	return arr[len(arr)-1]
}