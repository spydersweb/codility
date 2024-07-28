package main

import "fmt"

func main() {

	A := []int{0, 1, 0, 1, 1}

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

	fmt.Println(passes)
}
