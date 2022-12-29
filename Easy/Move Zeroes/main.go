package main

import "fmt"

// func moveZeroes(nums []int) {
// 	var counter int

// 	for index := 0; index < len(nums)-counter; index++ {
// 		if nums[index] == 0 {
// 			nums[index] = nums[len(nums)-1-counter]
// 			nums[len(nums)-1-counter] = 0
// 			counter++
// 		}
// 	}
// }

//mora biti sortirano
// func moveZeroes(nums []int) {
// 	var counter int

// 	for index := 0; index < len(nums)-counter; index++ {
// 		if nums[index] == 0 {
// 			nums = append(nums[:index], nums[index+1:]...)
// 			index--
// 			counter++
// 			nums = append(nums, 0)
// 		}
// 	}
// }

func moveZeroes(nums []int) {
	left := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		nums[left] = nums[i]
		left++
	}
	//Obavezno koristiti u buducnosti
	for ; left < len(nums); left++ {
		nums[left] = 0
	}
}

func main() {
	my_Array := []int{0, 0, 1}

	moveZeroes(my_Array)

	fmt.Println(my_Array)
}
