package find

import (
	mb "Lab8/matrix_build"
)

const size = 16

func SearchByMatch(matrix [size][size]int32, index int) [size]int32 {
	attribute := mb.StraightRow(matrix[index], index)
	var result, currWord [size]int32
	matches, currMatches := 0, 0

	for i := range matrix {
		currWord = mb.StraightRow(matrix[i], i)

		for j := 0; j < size; j++ {
			if currWord[j] == attribute[i] {
				currMatches++
			}
		}
		if currMatches >= matches && i != index {
			matches = currMatches
			result = currWord
		}
	}

	return result
}
