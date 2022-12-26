package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {

	if nums[len(nums)-1] < target {
		return len(nums)
	}

	index := 0
	for nums[index] < target && index != len(nums)-1 {
		index++
	}

	return index
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}
