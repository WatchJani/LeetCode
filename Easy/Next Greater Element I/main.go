package main

import "fmt"

//najboljii :D
func nextGreaterElement(nums1 []int, nums2 []int) []int {

	newIndex := -1
	newArray := make([]int, len(nums1))

	for index, nums1 := range nums1 {
		for _, nums2 := range nums2 {

			if nums1 == nums2 {
				newIndex = nums1
			}

			if newIndex != -1 {
				if newIndex < nums2 {
					newArray[index] = nums2
					break
				}
			}
		}
		if newArray[index] == 0 {
			newArray[index] = -1
		}
		newIndex = -1
	}

	return newArray
}

func main() {
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}
