package main

import (
	"testing"
)

type test struct {
	name           string
	inputArray     []int
	expectedResult int
}

func TestMissingIntegers(t *testing.T) {
	tests := []test{
		{
			name:           "Testing for missing within an array",
			inputArray:     []int{1, 3, 6, 4, 1, 2},
			expectedResult: 5,
		},
		{
			name:           "Testing for end of array",
			inputArray:     []int{1, 2, 3},
			expectedResult: 4,
		},
		{
			name:           "Testing for negatives",
			inputArray:     []int{-1, -3},
			expectedResult: 1,
		},
	}

	for _, test := range tests {
		result := MissingIntegers(test.inputArray)
		if result != test.expectedResult {
			t.Errorf("Unexpected result. Exp: %v, Got: %v\n", test.expectedResult, result)
		}
	}
}
