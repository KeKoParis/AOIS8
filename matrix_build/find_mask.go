package matrix_build

func FindMask(matrix [size][size]int32, mask [3]int32) ([size]int32, int) {

	for i := 0; i < size; i++ {
		if compare(StraightRow(matrix, i), mask) {
			return StraightRow(matrix, i), i
		}
	}

	return matrix[0], -1
}

func compare(row [size]int32, mask [3]int32) bool {

	for i := 0; i < 3; i++ {
		if row[size-4+i] != mask[i] {
			return false
		}
	}

	return true
}
