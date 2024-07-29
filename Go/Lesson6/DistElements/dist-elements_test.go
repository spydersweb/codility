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
			name:           "Default test",
			inputArray:     []int{2, 1, 1, 2, 3, 1},
			expectedResult: 3,
		},
		{
			name:           "Different test with 0",
			inputArray:     []int{99, 98, 0, 1},
			expectedResult: 4,
		},
		{
			name:           "All same elements",
			inputArray:     []int{1, 1, 1, 1},
			expectedResult: 1,
		},
	}

	for _, test := range tests {
		result := Solution(test.inputArray)
		if result != test.expectedResult {
			t.Errorf("unexpected error. Exp: %v, Got: %v", test.expectedResult, result)
		}
	}
}
