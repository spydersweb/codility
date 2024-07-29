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
			inputArray:     []int{10, 2, 5, 1, 8, 20},
			expectedResult: 1,
		},
		{
			name:           "Expected failure",
			inputArray:     []int{10, 50, 5, 1},
			expectedResult: 0,
		},
	}

	for _, test := range tests {
		result := Solution(test.inputArray)
		if result != test.expectedResult {
			t.Errorf("unexpected error. Exp: %v, Got: %v", test.expectedResult, result)
		}
	}
}
