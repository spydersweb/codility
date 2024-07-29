package main

import (
	"math"
)

func NumJumps(X, Y, D int) int {
	if X > Y {
		return 0
	}

	jumps := 0

	for i := X; i < Y; i += D {
		jumps += 1
	}

	return jumps
}

func NumJumpsMod(X, Y, D int) int {
	if X > Y {
		return 0
	}

	v := float64((Y - X) / D)
	if int(v) < Y {
		v += 1
	}

	return int(math.Round(v))

}
