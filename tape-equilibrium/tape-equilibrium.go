package main

import "slices"

func TapeEquilibrium(A []int) int {
	results := []int{}

	for i := 0; i < len(A); i++ {
		pSum := SumArray(A[0 : i+1])

		bSum := SumArray(A[i+1 : len(A)])
		var diff int = bSum - pSum
		if pSum >= bSum {
			diff = pSum - bSum
		}
		results = append(results, diff)
	}

	slices.Sort(results)

	return results[0]
}

func SumArray(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
