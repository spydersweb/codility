package main

import "testing"

type test struct {
	name          string
	inputArray    []int
	expectedArray []int
	NoCycles      int
}

func TestCyclicArray(t *testing.T) {

	tests := []test{
		{
			name:          "test for modular cycles with multiples of length",
			inputArray:    []int{1, 2, 3, 4},
			expectedArray: []int{1, 2, 3, 4},
			NoCycles:      12,
		},
		{
			name:          "test for cycles less than length",
			inputArray:    []int{1, 2, 3, 4},
			expectedArray: []int{3, 4, 1, 2},
			NoCycles:      2,
		},
		{
			name:          "test for cycles greater  than length",
			inputArray:    []int{1, 2, 3, 4},
			expectedArray: []int{3, 4, 1, 2},
			NoCycles:      6,
		},
		{
			name:          "test for odd cycles",
			inputArray:    []int{1, 2, 3, 4},
			expectedArray: []int{2, 3, 4, 1},
			NoCycles:      3,
		},
	}

	for _, b := range tests {
		gotArray := CycleArray(b.inputArray, b.NoCycles)
		for i := 0; i <= len(b.expectedArray)-1; i++ {
			if gotArray[i] != b.expectedArray[i] {
				t.Errorf("unexpected value in array sequence.  Exp %v, Got %v", b.expectedArray[i], gotArray[i])
			}
		}
	}
}
