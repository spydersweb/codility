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
			inputArray:     []int{-3, 1, 2, -2, 5, 6},
			expectedResult: 60,
		},
		{
			name:           "Second Test all same numbers",
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
