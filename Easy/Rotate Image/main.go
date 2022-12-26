package main

import "fmt"

//180
func rotate(matrix [][]int) [][]int {

	width := len(matrix[0])
	height := len(matrix)

	for pixel := 0; pixel < height/2*width; pixel++ {
		i := pixel / height
		j := pixel % width

		myValue := matrix[i][j]
		matrix[i][j] = matrix[height-i-1][width-j-1]
		matrix[height-i-1][width-j-1] = myValue
	}

	return matrix
}

//90 LOSIJA VERZIJA
func Rotate(matrix [][]int) [][]int {

	newMatrix := make([][]int, len(matrix))

	for index := range newMatrix {
		newMatrix[index] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix)*len(matrix[0]); i++ {
		newMatrix[i/len(matrix)][i%len(matrix)] = matrix[len(matrix)-i%len(matrix[0])-1][i/len(matrix)]
	}

	return newMatrix
}

//90
func ROTATE(matrix [][]int) [][]int {
	for i := 0; i < len(matrix)/2*len(matrix[0]); i++ {
		matrix[i/len(matrix)][i%len(matrix)] = matrix[len(matrix)-i%len(matrix[0])-1][i/len(matrix)]
	}

	return matrix
}

func main() {

	//TESTS

	//parno
	parno := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	// neparno := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }
	fmt.Println(ROTATE(parno))

}
