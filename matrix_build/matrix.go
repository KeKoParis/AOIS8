package matrix_build

import "math/rand"

const size = 16

func GetMatrix() [size][size]int32 {
	var matrix [size][size]int32

	rand.Seed(12)

	for i := range matrix {
		for j := range matrix {
			matrix[i][j] = rand.Int31n(2)
		}
	}

	return matrix
}

func Normalize(matrix *[size][size]int32) {
	var copyMatrix [size][size]int32
	for i := range matrix {
		for j := range matrix[i] {
			copyMatrix[j][i] = matrix[i][j]
		}
	}

	*matrix = copyMatrix
}

func Diagonalize(matrix *[size][size]int32) {
	for i := range *matrix {
		DiagonalizeRow(&matrix[i], i)
	}
}

func DiagonalizeRow(row *[size]int32, index int) {
	var currRow [size]int32

	position := 0
	for i := index; i < size; i++ {
		currRow[i] = row[position]
		position++
	}
	for i := 0; i < index; i++ {
		currRow[i] = row[position]
		position++
	}

	*row = currRow
}

func PutWord(matrix *[size][size]int32, row [size]int32, index int) {
	Normalize(matrix)
	DiagonalizeRow(&row, index)

	matrix[index] = row

	Normalize(matrix)
}

func StraightRow(matrix [size][size]int32, index int) [size]int32 {
	Normalize(&matrix)
	row := matrix[index]
	var currArr [size]int32

	position := 0
	for i := index; i < size; i++ {
		currArr[position] = row[i]
		position++
	}

	for i := 0; i < index; i++ {
		currArr[position] = row[i]
		position++
	}

	return currArr
}
