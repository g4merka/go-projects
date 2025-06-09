// https://www.codewars.com/kata/563b662a59afc2b5120000c6

package main

import (
	"fmt"
)

func main() {
	fmt.Println(NbYear(1500000, 0.25, -1000, 2000000))
}

func NbYear(p0 int, percent float64, aug int, p int) int {
	var f0 float64
	i := 0
	for i = 0; p0 < p; i++ {
		f0 = float64(p0) + (float64(p0) * (percent / float64(100))) + float64(aug)
		p0 = int(f0)
	}
	return i
}

/* Best:
func NbYear(p0 int, percent float64, aug int, p int) (n int) {
	if p0 >= p {
		return n
	}
	newP := p0 + aug + int(float64(p0)*percent/100)
	n++
	return n + NbYear(newP, percent, aug, p)
}
*/
