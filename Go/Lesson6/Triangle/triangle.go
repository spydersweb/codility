package main

func Solution(A []int) int {
	found := 0

	if len(A) < 3 {
		return found
	}

	P := 0
	Q := 1
	R := 2

	maxThirdIdx := len(A) - 1
	maxSecondIdx := len(A) - 2
	maxFirstIdx := len(A) - 3

	for {
		if P == maxFirstIdx &&
			Q == maxSecondIdx &&
			R == maxThirdIdx {
			break
		}

		if (A[P]+A[Q]) > A[R] &&
			(A[R]+A[P]) > A[Q] &&
			(A[Q]+A[R]) > A[P] {
			found = 1
			break
		}

		if R < maxThirdIdx {
			R++
			continue
		} else if R == maxThirdIdx && Q < maxSecondIdx {
			Q++
			R = Q + 1
			continue
		} else if R == maxThirdIdx && Q == maxSecondIdx {
			P++
			Q = P + 1
			R = Q + 1
		}
	}

	return found
}
