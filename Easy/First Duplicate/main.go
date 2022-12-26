package main

import (
	"fmt"
	"math"
)

func firstDuplicate(a []int) int {
	for i := 0; i < len(a); i++ {
		if a[int(math.Abs(float64(a[i])))-1] < 0 {
			return int(math.Abs(float64(a[i])))
		} else {
			a[int(math.Abs(float64(a[i])))-1] = -a[int(math.Abs(float64(a[i])))-1]
		}
	}

	return -1
}

func main() {
	fmt.Println(firstDuplicate([]int{2, 1, 3, 5, 3, 2}))
}
