package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args[1]
	testCases := formatInput(input)
	rawOutput := zigZagTestCases(testCases)
	output := formatOutput(rawOutput)

	fmt.Println(output)
}
