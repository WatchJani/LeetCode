package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	richest := 0

	for _, user := range accounts {
		sum := 0
		for _, value := range user {
			sum += value
		}
		if sum > richest {
			richest = sum
		}
	}

	return richest
}

func main() {
	fmt.Println(maximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))
}
