package main

import (
	"Lab8/find"
	"Lab8/logic"
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
	fmt.Println("Word: ", index, row)

	matrix[index] = row

	println("\nFunctions, input 3d word:")
	funcRow := mb.StraightRow(matrix[2], 2)
	fmt.Println(funcRow)

	println("f5 & f10")
	fmt.Println(logic.Function10(logic.Function5(funcRow)))

	println("f0 & f15")
	fmt.Println(logic.Function15(logic.Function0(funcRow)))

	println("\nFind:\nAttribute:")
	fmt.Println(mb.StraightRow(matrix[4], 4))

	fmt.Println("\nResult:\n", find.SearchByMatch(matrix, 1))
}
