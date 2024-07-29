package main

import (
	"testing"
)

func TestOddOccurances(t *testing.T) {

	type test struct {
		name           string
		inputArray     []int
		expectedResult int
	}

	tests := []test{
		{
			name:           "Testing simple array",
			inputArray:     []int{9, 5, 3, 9, 7, 3, 7},
			expectedResult: 5,
		},
		{
			name:           "Testing array with 3 indexes",
			inputArray:     []int{9, 9, 3, 9, 7, 3, 7},
			expectedResult: 9,
		},
	}

	for _, test := range tests {
		result := OddOccurances(test.inputArray)
		if result != test.expectedResult {
			t.Errorf("Unexpected result. Exp: %d, Got: %d", test.expectedResult, result)
		}
	}
}
