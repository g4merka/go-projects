// https://www.codewars.com/kata/56747fd5cb988479af000028

package main

import (
	"fmt"
)

func main() {
	fmt.Println(GetMiddle("testing"))
}

func GetMiddle(s string) string {
	if len(s) <= 2 {
		return s
	}
	s = s[1 : len(s)-1]
	return GetMiddle(s)
}
