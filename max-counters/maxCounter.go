package main

import (
	"fmt"
	"slices"
)

func Solution(N int, A []int) []int {

	counters := make([]int, N)

	for i := 0; i < len(A); i++ {
		if A[i] > len(counters) {
			// maximise all counters
			copy(counters, MaxCounter(counters))
			continue
		}

		counters[A[i]-1] += 1
	}

	return counters
}

func MaxCounter(B []int) []int {

	slices.Sort(B)

	fmt.Println(B)
	m := B[len(B)-1]

	for i, _ := range B {
		B[i] = m
	}

	return B
}
