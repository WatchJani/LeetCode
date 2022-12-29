package main

import "fmt"

// func maximumWealth(accounts [][]int) int {

// 	var sum int

// 	for _, account := range accounts {
// 		miniSum := 0
// 		for _, value := range account {
// 			miniSum += value
// 		}
// 		if miniSum > sum {
// 			sum = miniSum
// 		}
// 	}

// 	return sum
// }

//nesto novo nauceno (malo bolje rjesenje od moga)
func maximumWealth(accounts [][]int) (max int) {
	for _, account := range accounts {
		sum := 0
		for _, money := range account {
			sum += money
		}
		if sum > max {
			max = sum
		}
	}
	return
}

func main() {
	fmt.Println(maximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}}))
}
