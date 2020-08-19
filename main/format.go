package main

import (
	"strconv"
	"strings"
)

func formatInput(input string) [][]int {
	sliceA := strings.Split(input, "\n")

	var sliceB [][]int
	for i, subString := range sliceA {
		subSliceA := strings.Split(subString, " ")

		var subSliceB []int
		for i, subString := range subSliceA {
			subSliceB[i], _ = strconv.Atoi(subString)
		}

		sliceB[i] = subSliceB
	}

	return sliceB
}
