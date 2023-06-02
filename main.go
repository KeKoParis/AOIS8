package main

import (
	mb "Lab8/matrix_build"
	"Lab8/sum"
	"fmt"
)

const size = 16

func printMatrix(matrix [size][size]int32) {
	for i := range matrix {
		if i < 10 {
			fmt.Println("S", i, "  ", matrix[i])
		} else {
			fmt.Println("S", i, " ", matrix[i])
		}
	}
}

func main() {
	matrix := mb.GetMatrix()
	println("Matrix:")
	printMatrix(matrix)

	mb.Diagonalize(&matrix)

	println("\nDiagonalized:")
	printMatrix(matrix)

	v := [3]int32{1, 0, 1}

	println("\nFound word:")
	row, index := mb.FindMask(matrix, v)
	fmt.Println("Word: ", index, row)

	sum.Sum(&row)
	println("\nWord after:")

}
