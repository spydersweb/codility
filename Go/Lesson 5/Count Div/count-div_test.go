package main

import "testing"

type test struct {
	name           string
	start          int
	end            int
	divisor        int
	expectedResult int
}

func TestSolution(t *testing.T) {
	tests := []test{
		{
			name:           "Simple Test",
			start:          6,
			end:            11,
			divisor:        2,
			expectedResult: 3,
		},
		{
			name:           "B > A",
			start:          11,
			end:            6,
			divisor:        2,
			expectedResult: 0,
		},
		{
			name:           "max int B",
			start:          6,
			end:            2000000001,
			divisor:        2,
			expectedResult: 0,
		},
	}

	for _, test := range tests {
		result := Solution(test.start, test.end, test.divisor)
		if result != test.expectedResult {
			t.Errorf("Unexpected result.  Exp: %v, Got: %v", test.expectedResult, result)
		}
	}
}
