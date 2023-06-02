package sum

import (
	"fmt"
)

const size = 16

func Sum(word *[size]int32) {
	numA, numB := getFields(*word)

	numS := sumBin(numA, numB)
	fmt.Printf("\nnumA %d \nnumB %d \nsum  %d \n", numA, numB, numS)

	for i := range numS {
		(*word)[i] = numS[i]
	}

}

func getFields(word [size]int32) ([5]int32, [5]int32) {
	var numA, numB [5]int32
	for i := 5; i < 9; i++ {
		numA[i-5] = word[i]
	}
	for i := 9; i < 13; i++ {
		numB[i-9] = word[i]
	}

	return numA, numB
}

func sumBin(numA [5]int32, numB [5]int32) [5]int32 {

	for i := range numA {
		fmt.Println(numA)
		if numB[i] == 1 {
			findZero(&numA, i)
		}
	}

	return numA
}

func findZero(num *[5]int32, index int) {

	for i := index; i < len(num); i++ {
		if num[i] == 0 {
			num[i] = 1
			break
		} else {
			num[i] = 0
		}
	}

}
