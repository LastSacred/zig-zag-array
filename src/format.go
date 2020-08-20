package main

import (
	"strconv"
	"strings"
)

func formatInput(rawInput string) [][]int {
	topSlice := strings.Split(rawInput, "\n")

	var formattedInput [][]int
	for i, subString := range topSlice {
		if i < 2 || i%2 != 0 {
			continue
		}

		stringSlice := strings.Split(subString, " ")
		intSlice := stringSliceToIntSlice(stringSlice)
		formattedInput = append(formattedInput, intSlice)
	}

	return formattedInput
}

func stringSliceToIntSlice(stringSlice []string) []int {
	var intSlice []int
	for _, subString := range stringSlice {
		number, _ := strconv.Atoi(subString)
		intSlice = append(intSlice, number)
	}

	return intSlice
}

func formatOutput(rawOutput [][]int) string {
	var unjoinedOutput []string
	for _, rawCase := range rawOutput {
		var unjoinedCase []string
		for _, number := range rawCase {
			unjoinedCase = append(unjoinedCase, strconv.Itoa(number))
		}

		joinedCase := strings.Join(unjoinedCase, " ")
		unjoinedOutput = append(unjoinedOutput, joinedCase)
	}

	formattedOutput := strings.Join(unjoinedOutput, "\n")
	return formattedOutput
}
