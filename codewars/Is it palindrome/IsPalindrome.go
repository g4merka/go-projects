// https://www.codewars.com/kata/57a1fd2ce298a731b20006a4
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(IsPalindrome("MAdem"))
}

func IsPalindrome(str string) bool {
	lowercase := strings.ToLower(str)
	runes := []rune(lowercase)
	
	var newRunes []rune
	for i := len(runes)-1; i >= 0; i-- {
		newRunes = append(newRunes, runes[i])
	}

	return lowercase == string(newRunes)
}