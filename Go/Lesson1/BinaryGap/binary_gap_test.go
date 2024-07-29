package main

import (
	"testing"
)

type test struct {
	name       string
	inputInt   int
	largestGap int
}

func TestBinaryGap(t *testing.T) {
	tests := []test{
		{
			name:       "Testing 1041",
			inputInt:   1041,
			largestGap: 5,
		},
		{
			name:       "Testing 32 trailing zeros",
			inputInt:   32,
			largestGap: 0,
		},
		{
			name:       "Testing 999",
			inputInt:   999,
			largestGap: 2,
		},
	}

	for _, te := range tests {
		t.Log(te.name)

		result := BinaryGap(te.inputInt)
		if te.largestGap != result {
			t.Errorf("Test Error, Exp: %v, Got: %v", te.largestGap, result)
		}

	}
}
