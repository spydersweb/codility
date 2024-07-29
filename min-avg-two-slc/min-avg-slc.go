package main

func Solution(A []int) int {
	startIndex := 0
	minAvg := 0.00
	start := 0
	incrementor := 2

	for {

		if incrementor == len(A) {
			break
		}

		for i := 0; i < len(A); i++ {
			start = i

			end := i + incrementor
			if end <= len(A) {
				testRange := A[start:end]
				sum := 0.00
				for _, v := range testRange {
					sum += float64(v)
				}
				sum = sum / float64(len(testRange))
				if minAvg == 0 || sum < minAvg {
					minAvg = sum
					startIndex = start
				}
			}
		}
		incrementor++
	}

	return startIndex
}
