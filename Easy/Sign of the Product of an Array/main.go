package main

import "fmt"

// func arraySign(nums []int) int {

// 	var counter int
// 	for _, value := range nums {
// 		if value < 0 {
// 			counter++
// 		}

// 		if value == 0 {
// 			return 0
// 		}
// 	}

// 	if counter%2 == 0 {
// 		return 1
// 	} else {
// 		return -1
// 	}
// }

func arraySign(nums []int) int {
	product := 1

	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num > 0 {
			product *= 1
		} else {
			product *= -1
		}
	}
	return product
}

func main() {
	fmt.Println(arraySign([]int{-1, 1, -1, 1, -1}))
}
