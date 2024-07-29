package main

import "testing"

type test struct {
	name           string
	inputArray     []int
	expectedResult int
}

func TestTapeEquilibrium(t *testing.T) {
	tests := []test{
		{
			name:           "Result 1",
			inputArray:     []int{3, 1, 2, 4, 3},
			expectedResult: 1,
		},
		{
			name:           "Result 2",
			inputArray:     []int{3, 1, 2, 4, 3, 5, 90, 1},
			expectedResult: 73,
		},
	}

	for _, test := range tests {
		result := TapeEquilibrium(test.inputArray)
		if result != test.expectedResult {
			t.Errorf("Unexpected Result, Exp: %v; Got: %v", test.expectedResult, result)
		}
	}

}
