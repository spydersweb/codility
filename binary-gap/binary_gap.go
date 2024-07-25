package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func BinaryGap(N int) int {

	// Convert to a binary representation of the number
	NasBinary := strconv.FormatInt(int64(N), 2)

	fmt.Println(NasBinary)

	// Now strip the string of any trailing 0s
	if strings.LastIndex(NasBinary, "1") < len(NasBinary) {
		NasBinary = NasBinary[0:strings.LastIndex(NasBinary, "1")]
	}

	// Split the remaining string on the "1"
	gaps := strings.Split(NasBinary, "1")

	fmt.Println(gaps)
	// Fail fast
	if len(gaps) <= 2 {
		return 0
	}

	// Order the array
	sort.Sort(sort.Reverse(sort.StringSlice(gaps)))

	return len(gaps[0])
}
