/*
https://www.codewars.com/kata/563a631f7cbbc236cf0000c2
*/

package main

import "fmt"

func main() {
	fmt.Println(Move(3, 6))
}

func Move(position, roll int) int {
	return position + (roll * 2)
}