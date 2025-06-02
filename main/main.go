package main

import "fmt"

func main() {
	str := "abba"
	runes := []rune(str)

	res := runes[1]+runes[2]
	fmt.Println(runes[1], runes[2], res)
}
