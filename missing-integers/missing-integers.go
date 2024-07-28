package main

import (
	"fmt"
	"slices"
)

func MissingIntegers(A []int) int {

	slices.Sort(A)
	fmt.Println(A)

	nextElement := 0
	for i := 0; i < len(A); i++ {

		nextElement = A[i] + 1

		nextIndex := i + 1
		if nextIndex >= len(A) {
			break
		}

		if A[i] == A[nextIndex] {
			continue
		}

		if A[nextIndex] == nextElement {
			continue
		}

		if A[i] > 0 && A[nextIndex] != nextElement {
			break
		}
	}
	if nextElement <= 0 {
		return 1
	}

	return nextElement
}
