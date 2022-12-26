package main

import (
	"fmt"
)

//moj najljepsi program
func countOdds(low int, high int) int {
	return (high + high%2 - low + 1) / 2
}

func main() {
	fmt.Println(countOdds(7, 11))
}
