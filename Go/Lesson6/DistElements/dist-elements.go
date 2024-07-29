package main

import "slices"

func Solution(A []int) int {

	slices.Sort(A)

	current := 10000001
	count := 0

	for _, v := range A {

		if current == v || v > 1000000 || v < -1000000 {
			continue
		}

		current = v
		count++
	}

	return count
}
