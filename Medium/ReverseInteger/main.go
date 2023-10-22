package main

import (
	"fmt"
	"math"
)

func ReverseInteger(x int) int {
	myInt := 0

	for x != 0 {
		myInt = myInt*10 + x%10

		x /= 10

		if myInt > math.MaxInt32 || myInt < math.MinInt32 {
			return 0
		}
	}

	return myInt
}

func main() {
	fmt.Println(ReverseInteger(-123))
}
