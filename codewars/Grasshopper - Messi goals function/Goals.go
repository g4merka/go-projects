// https://www.codewars.com/kata/55f73be6e12baaa5900000d4

package main

import (
	"fmt"
)

func main() {
	fmt.Println(Goals(43, 10, 5))
}

func Goals(laLigaGoals, copaDelReyGoals, championsLeagueGoals int) int {
	return laLigaGoals + copaDelReyGoals + championsLeagueGoals
}
