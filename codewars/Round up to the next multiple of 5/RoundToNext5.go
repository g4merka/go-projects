// https://www.codewars.com/kata/55d1d6d5955ec6365400006d

package main

import (
	"fmt"
)

func main() {
	fmt.Println(RoundToNext5(-2))
}

func RoundToNext5(n int) int {
	if n%5 == 0 {
		return n
	}
	return RoundToNext5(n + 1)
}

/*
func RoundToNext5(n int) int {
  for n % 5 != 0 {
    n++
  }
  return n
}
*/
