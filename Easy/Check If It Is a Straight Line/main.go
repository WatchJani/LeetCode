package main

import (
	"fmt"
)

// func checkStraightLine(coordinates [][]int) bool {
// 	a := coordinates[1][1] - coordinates[0][1]
// 	b := coordinates[0][0] - coordinates[1][0]
// 	c := coordinates[1][0]*coordinates[0][1] - coordinates[0][0]*coordinates[1][1]

// 	for index := 2; index < len(coordinates); index++ {
// 		if a*coordinates[index][0]+b*coordinates[index][1]+c != 0 {
// 			return false
// 		}
// 	}

// 	return true
// }

func checkStraightLine(coordinates [][]int) bool {
	var m, b float64
	if coordinates[0][0] == coordinates[1][0] {
		for i := 2; i < len(coordinates); i++ {
			if coordinates[i][0] != coordinates[0][0] {
				return false
			}
		}
	} else {
		m = (float64(coordinates[1][1]) - float64(coordinates[0][1])) / (float64(coordinates[1][0]) - float64(coordinates[0][0]))
		b = float64(coordinates[0][1]) - (m * float64(coordinates[0][0]))
		for i := 2; i < len(coordinates); i++ {
			if float64(coordinates[i][1]) != (m*float64(coordinates[i][0]))+b {
				return false
			}
		}
	}
	return true
}

func main() {

	test := [][]int{{2, 3}, {6, 7}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}
	//test1 := [][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}
	//test2 := [][]int{{1, 2}, {2, 3}, {3, 5}}
	//test3 := [][]int{{0, 0}, {0, 1}, {0, -1}}
	fmt.Println(checkStraightLine(test))
}
