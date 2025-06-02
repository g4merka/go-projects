// https://www.codewars.com/kata/57eae65a4321032ce000002d
package main

import (
	"fmt"
)

func main() {
	fmt.Println(FakeBin("45385593107843568"))
}

func FakeBin(x string) string {
	runes := []rune(x)
	var newRunes []rune
	for _, v := range runes {
		if v < 53 {
			newRunes = append(newRunes, 48)
			continue
		}
		newRunes = append(newRunes, 49)
	}
	return string(newRunes)
}

/* Best:

func FakeBin(x string) (result string) {
    for _, char := range x {
        if char < '5' {
            result += "0"
        } else {
            result += "1"
        }
    }
    return
}
*/
