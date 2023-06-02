package logic

import "math"

const size = 16

func Function5(row [size]int32) [size]int32 {
	var result [size]int32

	for i := 0; i < size; i++ {
		result[i] = row[i]
	}

	return result
}

func Function10(row [size]int32) [size]int32 {
	var result [size]int32

	for i := 0; i < size; i++ {
		result[i] = int32(math.Abs(float64(row[i] - 1)))
	}

	return result
}

func Function0(row [size]int32) [size]int32 {
	var result [size]int32

	for i := 0; i < size; i++ {
		result[i] = 0
	}

	return result
}

func Function15(row [size]int32) [size]int32 {
	var result [size]int32

	for i := 0; i < size; i++ {
		result[i] = 1
	}

	return result
}
