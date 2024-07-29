package main

import "strings"

const C = 67 // 2
const A = 65 // 1
const G = 71 // 3
const T = 84 // 4

func Solution(S string, P []int, Q []int) []int {
	result := make([]int, len(P))

	// Character check
	for _, s := range S {
		if !strings.Contains("CAGT", string(s)) {
			return result
		}
	}

	if len(P) != len(Q) {
		return result
	}

	for idx, start := range P {

		if start > Q[idx] {
			continue
		}

		testRange := S[start : Q[idx]+1]
		result[idx] = 9

		for _, b := range testRange {
			testValue := 0
			switch b {
			case C:
				testValue = 2
			case A:
				testValue = 1
			case G:
				testValue = 3
			case T:
				testValue = 4
			default:
				testValue = 10
			}

			if testValue < result[idx] {
				result[idx] = testValue
			}
		}

	}

	return result
}
