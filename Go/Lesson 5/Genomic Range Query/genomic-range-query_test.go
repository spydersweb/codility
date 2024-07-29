package main

import (
	"reflect"
	"testing"
)

type test struct {
	name           string
	sequence       string
	startArray     []int
	endArray       []int
	expectedResult []int
}

func TestSolution(t *testing.T) {
	tests := []test{
		{
			name:           "Default Test",
			sequence:       "CAGCCTA",
			startArray:     []int{2, 5, 0},
			endArray:       []int{4, 5, 6},
			expectedResult: []int{2, 4, 1},
		},
		{
			name:           "Testing bad character",
			sequence:       "CAGPCCTA",
			startArray:     []int{2, 5, 0},
			endArray:       []int{4, 5, 6},
			expectedResult: []int{0, 0, 0},
		},
		{
			name:           "Testing skipping test when start is greater than end",
			sequence:       "CAGCCTA",
			startArray:     []int{2, 6, 0},
			endArray:       []int{4, 5, 6},
			expectedResult: []int{2, 0, 1},
		},
	}

	for _, test := range tests {
		result := Solution(test.sequence, test.startArray, test.endArray)

		if !reflect.DeepEqual(result, test.expectedResult) {
			t.Errorf("Unexpected result. Exp: %v, Got: %v", test.expectedResult, result)
		}
	}
}
