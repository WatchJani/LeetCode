package main

import "fmt"

//moj drugi najbolji zadatak uradjen
func subtractProductAndSum(n int) int {

	sum := 0
	product := 1

	for n > 0 {
		sum += n % 10
		product *= n % 10
		n /= 10
	}

	return product - sum
}

func main() {
	fmt.Println(subtractProductAndSum(4421))
}
