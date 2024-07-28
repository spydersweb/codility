package main

func Solution(A, B, K int) int {
	c := 0
	if A < B && B < 2000000000 {
		for i := A; i <= B; i++ {
			if i%K == 0 {
				c++
			}
		}
	}

	return c
}
