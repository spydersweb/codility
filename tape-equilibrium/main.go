package main

import (
	"fmt"
)

func main() {
	A := []int{3, 1, 2, 4, 3}

	results := TapeEquilibrium(A)
	fmt.Printf("%v\n", results)
}
