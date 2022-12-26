package main

import "fmt"

func removeDuplicates(nums []int) int {

	lastUniq := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[lastUniq] {
			nums[lastUniq+1] = nums[i]
			lastUniq++
		}
	}

	return lastUniq + 1
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}
