package main

func Solution(A []int) int {
	sum := 0
	maxTriplet := 0
	if len(A) < 3 {
		return sum
	}

	first := 0
	second := 1
	third := 2

	maxThirdIdx := len(A) - 1
	maxSecondIdx := len(A) - 2
	maxFirstIdx := len(A) - 3

	// -3, 1, 2, -2, 5, 6
	for {
		if first == maxFirstIdx &&
			second == maxSecondIdx &&
			third == maxThirdIdx {
			break
		}

		sum = A[first] * A[second] * A[third]
		if sum > maxTriplet {
			maxTriplet = sum
		}

		if third < maxThirdIdx {
			second++
			third = second + 1
			continue
		} else if third == maxThirdIdx && second == maxSecondIdx {
			first++
			second = first + 1
			third = second + 1
			continue
		}
	}

	return sum
}
