package main

import (
	"fmt"
)

func main() {
	testCases := formatInput(input)
	rawOutput := zigZagTestCases(testCases)
	output := formatOutput(rawOutput)

	fmt.Println(output)
}
