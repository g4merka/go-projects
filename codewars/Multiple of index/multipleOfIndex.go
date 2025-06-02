// https://www.codewars.com/kata/5a34b80155519e1a00000009
package main

import (
	"fmt"
)

func main() {
	fmt.Println(multipleOfIndex([]int{22, -6, 32, 82, 9, 25}))
}

func multipleOfIndex(ints []int) []int {
	var newInts []int

	for i := 1; i < len(ints); i++ {
		if ints[i] % i == 0 {
			newInts = append(newInts, ints[i])
		}
	}
	return newInts
}

// Best:
/*
func multipleOfIndex (ints []int) []int {
  res := make([]int, 0)
  
  for i, v := range ints {
    if i > 0 && v % i == 0 {
            res = append(res, v)
        }
    }
  return res
}
*/