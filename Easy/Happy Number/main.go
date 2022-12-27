package main

import "fmt"

func myInt(n int) int {
	var sum int

	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}

	return sum
}

func isHappy(n int) bool {
	for n != 1 && n != 4 {
		n = myInt(n)
	}
	return n == 1
}

func main() {
	fmt.Println(isHappy(2))
}
