// https://www.codewars.com/kata/52fba66badcd10859f00097e
package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(Dismvowel("This website is for losers LOL!"))
}

func Dismvowel(comment string) string {
	re := regexp.MustCompile(`[aeiouAEIOU]`)
	res := re.ReplaceAllString(comment, "")
	return res
}

// Best:

/*
import "regexp"

func Disemvowel(comment string) string {
  return regexp.MustCompile(`(?i)[aeiou]`).ReplaceAllString(comment, "")
}
*/