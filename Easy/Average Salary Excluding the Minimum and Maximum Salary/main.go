package main

import (
	"fmt"
)

func average(salary []int) float64 {

	max := salary[0]
	min := salary[0]
	sum := 0

	for _, value := range salary {
		if max < value {
			max = value
		}

		if min > value {
			min = value
		}
		sum += value
	}

	return float64(sum-(max+min)) / float64(len(salary)-2)
}

func main() {
	fmt.Println(average([]int{48000, 59000, 99000, 13000, 78000, 45000, 31000, 17000, 39000, 37000, 93000, 77000, 33000, 28000, 4000, 54000, 67000, 6000, 1000, 11000}))
}
