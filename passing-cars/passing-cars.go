package main

func Solution(A []int) int {
	passes := 0
	for i, v := range A {
		if v > 0 {
			continue
		}
		for _, v := range A[i:] {
			if v == 1 {
				passes++
			}
		}
	}

	if passes > 1000000000 {
		return -1
	}

	return passes
}
