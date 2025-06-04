// https://www.codewars.com/kata/555086d53eac039a2a000083
package main

import (
	"fmt"
)

func main() {
	fmt.Println(LoveFunc(1, 4))
}

func LoveFunc(flower1, flower2 int) bool {
	if (flower1 % 2 != 0 && flower2 % 2 != 0) || (flower1 % 2 == 0 && flower2 % 2 == 0) {
		return false
	}
	return true
}
/* Best:

func LoveFunc(flower1, flower2 int) bool {
  return flower1 % 2 != flower2 % 2
}
*/