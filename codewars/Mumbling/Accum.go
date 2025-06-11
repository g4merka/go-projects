// https://www.codewars.com/kata/5667e8f4e3f572a8f2000039

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Accum("ZpglnRxqenU"))
}

func Accum(s string) string {
	lowStr := strings.ToLower(s)
	var sb strings.Builder
	for i, v := range lowStr {
		if i == 0 {
			sb.WriteString(strings.ToUpper(string(v)))
			sb.WriteString("-")
		} else {
			sb.WriteString(strings.ToUpper(string(v)))
			for i > 0 {
				sb.WriteString(string(v))
				i--
			}
			sb.WriteString("-")
		}

	}
	return strings.TrimRight(sb.String(), "-")
}

/* Best:
import "strings"

func Accum(s string) string {
    parts := make([]string, len(s))

    for i := 0; i < len(s); i++ {
      parts[i] = strings.ToUpper(string(s[i])) + strings.Repeat(strings.ToLower(string(s[i])), i)
    }

    return strings.Join(parts, "-")
}
*/
