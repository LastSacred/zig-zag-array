package main

func zigZagTestCases(testCases [][]int) [][]int {
	var outputCases [][]int
	for _, testCase := range testCases {
		outputCase := zigZagTestCase(testCase)
		outputCases = append(outputCases, outputCase)
	}

	return outputCases
}

func zigZagTestCase(testCase []int) []int {
	for i := 0; i+1 < len(testCase); i++ {
		number := testCase[i]
		nextNumber := testCase[i+1]

		var swap bool
		if i%2 == 0 {
			if number > nextNumber {
				swap = true
			}
		} else if number < nextNumber {
			swap = true
		}

		if swap {
			testCase[i], testCase[i+1] = testCase[i+1], testCase[i]
		}
	}

	return testCase
}
