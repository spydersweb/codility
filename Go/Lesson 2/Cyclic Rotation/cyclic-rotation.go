package main

func CycleArray(A []int, K int) []int {

	if K%len(A) == 0 {
		return A
	}

	// Duplicate the length of the array
	newArray := make([]int, len(A))

	// Cycle the number of input times
	for i := 0; i < K; i++ {

		last := A[len(A)-1]

		// Take a copy of the last 3 elements
		for j, v := range A[0 : len(A)-1] {
			newArray[j+1] = v
		}

		newArray[0] = last

		// Overwrite A to match newArray
		copy(A, newArray)
	}

	return newArray
}
