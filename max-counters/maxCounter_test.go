package main

import (
	"reflect"
	"testing"
)

type test struct {
	name           string
	inputArray     []int
	numCounters    int
	expectedResult []int
}

func TestSolution(t *testing.T) {

	tests := []test{
		{
			name:           "test1",
			inputArray:     []int{3, 4, 4, 6, 1, 4, 4},
			numCounters:    5,
			expectedResult: []int{3, 2, 2, 4, 2},
		},
	}

	for _, test := range tests {
		results := Solution(test.numCounters, test.inputArray)
		if !reflect.DeepEqual(results, test.expectedResult) {
			t.Errorf("Unexpected result. Exp: %v, Got: %v\n", test.expectedResult, results)
		}
	}
}
