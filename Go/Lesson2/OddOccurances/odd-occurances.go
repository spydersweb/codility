package main

import (
	"slices"
)

type P struct {
	value   int
	indexes []int
}

func OddOccurances(A []int) int {

	pairs := []P{}

	for i, v := range A {

		idx := slices.IndexFunc(pairs, func(p P) bool { return p.value == v })
		if idx == -1 {
			pairs = append(pairs, P{
				value:   v,
				indexes: []int{i},
			})
			continue
		}
		pairs[idx].indexes = append(pairs[idx].indexes, i)

	}

	for _, v := range pairs {
		if len(v.indexes)%2 == 0 {
			continue
		}

		return v.value
	}

	return 0
}
