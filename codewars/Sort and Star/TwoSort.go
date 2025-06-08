// https://www.codewars.com/kata/57cfdf34902f6ba3d300001e
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(
		TwoSort(
			[]string{
				"turns",
				"out",
				"random",
				"test",
				"cases",
				"are",
				"easier",
				"than",
				"writing",
				"out",
				"basic",
				"ones",
			},
		),
	)
}

func TwoSort(arr []string) string {

	copied := make([]string, len(arr))
	copy(copied, arr)
	sort.Strings(copied)

	firstString := copied[0]
	var res string

	for idx, char := range firstString {
		res += string(char)
		if idx < len(firstString)-1 {
			res += "***"
		}
	}
	return res
}

/* Best:
package kata

import (
	"sort"
	"strings"
)

func TwoSort(arr []string) string {
	sort.Strings(arr)
	chars := strings.Split(arr[0], "")
	return strings.Join(chars, "***")
}
*/
