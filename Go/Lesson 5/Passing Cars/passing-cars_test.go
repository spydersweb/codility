package main

import "testing"

type test struct {
	name           string
	inputArray     []int
	expectedResult int
}

func TestSolution(t *testing.T) {
	tests := []test{
		{
			name:           "Solution example",
			inputArray:     []int{0, 1, 0, 1, 1},
			expectedResult: 5,
		},
		{
			name:           "Solution example expected zero passes",
			inputArray:     []int{1, 1, 1, 1, 0},
			expectedResult: 0,
		},
		{
			name:           "Solution example dodgy value",
			inputArray:     []int{0, 2, 1, 1, 1},
			expectedResult: 3,
		},
	}

	for _, test := range tests {
		result := Solution(test.inputArray)
		if result != test.expectedResult {
			t.Errorf("Unexpected result. Exp: %v, Got: %v", test.expectedResult, result)
		}
	}
}
