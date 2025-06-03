// https://www.codewars.com/kata/54ff3102c1bad923760001f3
package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(GetCount("abracadabra"))
}

func GetCount(str string) (count int) {
	re := regexp.MustCompile(`(?i)[aeiou]`)
	res := re.FindAllString(str, -1)
	return len(res)
}

/* Best:
func GetCount(str string) (count int) {
  for _, c := range str {
    switch c {
    case 'a', 'e', 'i', 'o', 'u':
      count++
    }
  }
  return count
}
*/