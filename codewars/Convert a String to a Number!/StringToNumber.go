// https://www.codewars.com/kata/544675c6f971f7399a000e79
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(StringToNumber("1234"))
}

func StringToNumber(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}