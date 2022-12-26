package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y := x
	length := int(math.Log10(float64(x))) + 1
	d := (int(math.Pow10(length - 1)))
	for i := 0; i < length/2; i++ {
		rDigit := y / d
		y = y % d
		d = d / 10
		lDigit := x % 10
		x = x / 10
		if rDigit != lDigit {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(121))
}
