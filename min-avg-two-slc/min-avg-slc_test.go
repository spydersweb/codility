package main

import "testing"

type test struct {
	name          string
	inputArray    []int
	expectedIndex int
}

func TestSolution(t *testing.T) {

	tests := []test{
		{
			name:          "default test",
			inputArray:    []int{4, 2, 2, 5, 1, 5, 8},
			expectedIndex: 1,
		},
		{
			name:          "all same values should return 0",
			inputArray:    []int{1, 1, 1, 1, 1},
			expectedIndex: 0,
		},
		{
			name:          "all same values should return 0",
			inputArray:    []int{2, 2, 2, 1},
			expectedIndex: 2,
		},
		{
			name:          "Middle index with 3 item range",
			inputArray:    []int{2, 2, 2, 1, 1, 1, 2, 2, 2},
			expectedIndex: 3,
		},
	}

	for _, test := range tests {
		result := Solution(test.inputArray)
		if result != test.expectedIndex {
			t.Errorf("unexpected result. Exp: %v, Got: %v\n", test.expectedIndex, result)
		}
	}
}
