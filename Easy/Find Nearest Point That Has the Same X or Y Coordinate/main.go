package main

import (
	"fmt"
	"math"
)

//old code
// func nearestValidPoint(x int, y int, points [][]int) int {

// 	myIndex := -1
// 	minDistance := math.MaxInt

// 	for index, points := range points {
// 		if x == points[0] && y == points[1] {
// 			return index
// 		}

// 		if distance := math.Abs(float64(x-points[0])) + math.Abs(float64(y-points[1])); (x == points[0] || y == points[1]) && distance < float64(minDistance) {
// 			myIndex = index
// 			minDistance = int(distance)
// 		}
// 	}

// 	return myIndex
// }

//better code
func nearestValidPoint(x int, y int, points [][]int) int {
	myIndex := -1
	minDistance := math.MaxInt

	for index, points := range points {
		if x == points[0] && y == points[1] {
			return index
		}

		if x == points[0] || y == points[1] {
			distance := math.Abs(float64(x-points[0])) + math.Abs(float64(y-points[1]))
			if distance < float64(minDistance) {
				minDistance = int(distance)
				myIndex = index
			}
		}
	}

	return myIndex
}

func main() {
	fmt.Println(nearestValidPoint(28, 51, [][]int{
		{25, 45}, {60, 19}, {11, 38}, {32, 22}, {1, 24}, {26, 25}, {52, 36}, {45, 54}, {45, 30}, {45, 39}, {39, 38}, {25, 38}, {39, 57}, {47, 51}, {47, 49}, {37, 21}, {16, 43}, {53, 33}, {10, 50}, {30, 29}, {3, 31}, {45, 26}, {22, 40}, {2, 31}, {57, 42}, {25, 42}, {37, 13}, {13, 54}, {17, 5}, {50, 32},
	}))
}
